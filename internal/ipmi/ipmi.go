// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package ipmi implements various IPMI functions.
package ipmi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"

	goipmi "github.com/pensando/goipmi"
	"go.uber.org/zap"
)

const ipmitoolPath = "/usr/sbin/ipmitool"

// Link to the IPMI spec: https://www.intel.com/content/dam/www/public/us/en/documents/product-briefs/ipmi-second-gen-interface-spec-v2-rev1-1.pdf

// Client is a holder for the IPMIClient.
type Client struct {
	IPMIClient *goipmi.Client
}

// NewLocalClient creates a new local ipmi client to use.
func NewLocalClient() (*Client, error) {
	conn := &goipmi.Connection{
		Path:      ipmitoolPath,
		Interface: "open",
	}

	ipmiClient, err := goipmi.NewClient(conn)
	if err != nil {
		return nil, err
	}

	if err = ipmiClient.Open(); err != nil {
		return nil, fmt.Errorf("error opening client: %w", err)
	}

	return &Client{IPMIClient: ipmiClient}, nil
}

// Close the client.
func (c *Client) Close() error {
	return c.IPMIClient.Close()
}

// AttemptUserSetup attempts to set up an IPMI user with the given username.
func (c *Client) AttemptUserSetup(username, password string, logger *zap.Logger) error {
	// Get user summary to see how many user slots
	summResp, err := c.getUserSummary()
	if err != nil {
		return err
	}

	maxUsers := summResp.MaxUsers & 0x1F // Only bits [0:5] provide this number

	// Check if the user already exists by combing through all userIDs
	// nb: we start looking at user id 2, because 1 should always be an unamed admin user and
	//     we don't want to confuse that unnamed admin with an open slot we can take over.
	exists := false
	userID := uint8(0)

	for i := uint8(2); i <= maxUsers; i++ {
		userRes, userErr := c.getUserName(i)
		if userErr != nil {
			// nb: A failure here actually seems to mean that the user slot is unused,
			// even though you can also have a slot with empty user as well. *scratches head*
			// We'll take note of this spot if we haven't already found another empty one.
			if userID == 0 {
				userID = i
			}

			continue
		}

		// Found pre-existing user
		if userRes.Username == username {
			exists = true
			userID = i

			logger.Info("user already present in slot, we'll claim it as our own", zap.Uint8("slot", i))

			break
		} else if (userRes.Username == "" || strings.TrimSpace(userRes.Username) == "(Empty User)") && userID == 0 {
			// If this is the first empty user that's not the UID 1 (which we skip),
			// we'll take this spot for our user
			logger.Info("found empty user slot, noting as a possible place for user", zap.Uint8("slot", i))

			userID = i
		}
	}

	// User didn't pre-exist and there's no room
	// Return without the user :(
	if userID == 0 {
		return errors.New("no slot available for user")
	}

	// Not already present and there's an empty slot so we add the user
	if !exists {
		logger.Info("adding user to slot", zap.Uint8("slot", userID))

		if _, err = c.setUserName(userID, username); err != nil { //nolint:errcheck
			return err
		}
	}

	if _, err = c.setUserPass(userID, password); err != nil { //nolint:errcheck
		return err
	}

	// Make the user an admin
	// Options: 0x91 == Callin true, Link false, IPMI Msg true, Channel 1
	// Limits: 0x03 == Administrator
	// Session: 0x00 No session limit
	if _, err = c.setUserAccess(0x91, userID, 0x04, 0x00); err != nil { //nolint:errcheck
		return err
	}

	// Enable the user
	if _, err = c.enableUser(userID); err != nil { //nolint:errcheck
		return err
	}

	return nil
}

// UserExists checks if a user exists on the BMC.
func (c *Client) UserExists(username string) (bool, error) {
	// Get user summary to see how many user slots
	summResp, err := c.getUserSummary()
	if err != nil {
		return false, err
	}

	maxUsers := summResp.MaxUsers & 0x1F // Only bits [0:5] provide this number

	// Check if the user already exists by combing through all userIDs
	for i := uint8(1); i <= maxUsers; i++ {
		userRes, userErr := c.getUserName(i)
		if userErr != nil {
			continue
		}

		if userRes.Username == username {
			return true, nil
		}
	}

	return false, nil
}

// GetIPPort returns the IPMI IP and port.
func (c *Client) GetIPPort() (ip string, port uint16, err error) {
	// Fetch BMC IP (param 3 in LAN config)
	ipResp, err := c.getLANConfig(0x03)
	if err != nil {
		return "", 0, err
	}

	// Fetch BMC Port (param 8 in LAN config)
	portResp, err := c.getLANConfig(0x08)
	if err != nil {
		return "", 0, err
	}

	ip = net.IP(ipResp.Data).String()
	port = binary.LittleEndian.Uint16(portResp.Data)

	return ip, port, nil
}

// getLANConfig fetches a given param from the LAN Config. (see 23.2).
func (c *Client) getLANConfig(param uint8) (*goipmi.LANConfigResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionTransport,
		Command:         goipmi.CommandGetLANConfig,
		Data: &goipmi.LANConfigRequest{
			ChannelNumber: 0x01,
			Param:         param,
		},
	}

	res := &goipmi.LANConfigResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// getUserSummary returns stats about user table, including max users allowed.
func (c *Client) getUserSummary() (*goipmi.GetUserSummaryResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionApp,
		Command:         goipmi.CommandGetUserSummary,
		Data: &goipmi.GetUserSummaryRequest{
			ChannelNumber: 0x01,
			UserID:        0x01,
		},
	}

	res := &goipmi.GetUserSummaryResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// getUserName fetches a un string given a uid. This is how we check if a user slot is available.
//
// nb: a "failure" here can actually mean that the slot is just open for use
// or you can also have a user with "" as the name which won't
// fail this check and is still open for use.
// (see 22.29).
func (c *Client) getUserName(uid byte) (*goipmi.GetUserNameResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionApp,
		Command:         goipmi.CommandGetUserName,
		Data: &goipmi.GetUserNameRequest{
			UserID: uid,
		},
	}

	res := &goipmi.GetUserNameResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// setUserName” sets a string for the given uid (see 22.28).
func (c *Client) setUserName(uid byte, name string) (*goipmi.SetUserNameResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionApp,
		Command:         goipmi.CommandSetUserName,
		Data: &goipmi.SetUserNameRequest{
			UserID:   uid,
			Username: name,
		},
	}

	res := &goipmi.SetUserNameResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// setUserPass sets the password for a given uid (see 22.30).
// nb: This naively assumes you'll pass a 16 char or less pw string.
//
//	The goipmi function does not support longer right now.
func (c *Client) setUserPass(uid byte, pass string) (*goipmi.SetUserPassResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionApp,
		Command:         goipmi.CommandSetUserPass,
		Data: &goipmi.SetUserPassRequest{
			UserID: uid,
			Pass:   []byte(pass),
		},
	}

	res := &goipmi.SetUserPassResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// setUserAccess tweaks the privileges for a given uid (see 22.26).
func (c *Client) setUserAccess(options, uid, limits, session byte) (*goipmi.SetUserAccessResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionApp,
		Command:         goipmi.CommandSetUserAccess,
		Data: &goipmi.SetUserAccessRequest{
			AccessOptions:    options,
			UserID:           uid,
			UserLimits:       limits,
			UserSessionLimit: session,
		},
	}

	res := &goipmi.SetUserAccessResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// enableUser sets a user as enabled. Actually the same underlying command as setUserPass (see 22.30).
func (c *Client) enableUser(uid byte) (*goipmi.EnableUserResponse, error) {
	req := &goipmi.Request{
		NetworkFunction: goipmi.NetworkFunctionApp,
		Command:         goipmi.CommandEnableUser,
		Data: &goipmi.EnableUserRequest{
			UserID: uid,
		},
	}

	res := &goipmi.EnableUserResponse{}

	err := c.IPMIClient.Send(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

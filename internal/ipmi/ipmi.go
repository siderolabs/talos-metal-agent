// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package ipmi implements various IPMI functions.
package ipmi

import (
	"encoding/binary"
	"errors"
	"net"
	"strings"

	"github.com/bougou/go-ipmi"
	"go.uber.org/zap"
)

// Link to the IPMI spec: https://www.intel.com/content/dam/www/public/us/en/documents/product-briefs/ipmi-second-gen-interface-spec-v2-rev1-1.pdf

// Client is a holder for the ipmiClient.
type Client struct {
	ipmiClient *ipmi.Client
}

// NewLocalClient creates a new local ipmi client to use.
func NewLocalClient() (*Client, error) {
	ipmiClient, err := ipmi.NewOpenClient()
	if err != nil {
		return nil, err
	}

	if err = ipmiClient.Connect(); err != nil {
		return nil, err
	}

	return &Client{ipmiClient: ipmiClient}, nil
}

// Close the client.
func (c *Client) Close() error {
	return c.ipmiClient.Close()
}

// AttemptUserSetup attempts to set up an IPMI user with the given username.
func (c *Client) AttemptUserSetup(username, password string, logger *zap.Logger) error {
	id, exists, err := c.findUserIDByName(username)
	if err != nil {
		return err
	}

	c.ipmiClient.sl

	users, err := c.ipmiClient.ListUser(0x01)
	if err != nil {
		return err
	}

	exists := false
	userID := uint8(0)

	for _, user := range users {
		if user.ID == 1 {
			continue // skip the default admin user
		}

		if user.Name == username {
			userID = user.ID
			exists = true

			break
		}

	}

	// Get user summary to see how many user slots
	userAccessResp, err := c.ipmiClient.GetUserAccess(0x01, 0x01)
	if err != nil {
		return err
	}

	// Check if the user already exists.
	// Start from user ID 2, as user ID 1 is reserved for the default admin user.
	exists := false
	userID := uint8(0)

	for i := uint8(2); i <= userAccessResp.MaxUsersIDCount; i++ {
		userRes, userErr := c.ipmiClient.GetUsername(i)
		if userErr != nil {
			// nb: A failure here actually seems to mean that the user slot is unused,
			// even though you can also have a slot with empty user as well. *scratches head*
			// We'll take note of this spot if we haven't already found another empty one.
			if userID == 0 {
				userID = i
			}

			continue
		}

		// Found the existing user
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

	if userID == 0 {
		return errors.New("no slot available for user")
	}

	// There's an empty slot, so we'll add the user
	if !exists {
		logger.Info("adding user to slot", zap.Uint8("slot", userID))

		if _, err = c.ipmiClient.SetUsername(userID, username); err != nil {
			return err
		}
	}

	if _, err = c.ipmiClient.SetUserPassword(userID, password, false); err != nil {
		return err
	}

	if _, err = c.ipmiClient.SetUserAccess(&ipmi.SetUserAccessRequest{
		EnableChanging:      true,
		EnableIPMIMessaging: true,
		ChannelNumber:       0x01,
		UserID:              userID,
		MaxPrivLevel:        0x04, // admin
		SessionLimit:        0,
	}); err != nil {
		return err
	}

	return c.ipmiClient.EnableUser(userID)
}

func (c *Client) findUserIDByName(username string) (uint8, bool, error) {
	users, err := c.ipmiClient.ListUser(0x01)
	if err != nil {
		return 0, false, err
	}

	for _, user := range users {
		if user.Name == username {
			return user.ID, true, nil
		}
	}

	return 0, false, nil
}

// UserExists checks if a user exists on the BMC.
func (c *Client) UserExists(username string) (bool, error) {
	_, exists, err := c.findUserIDByName(username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// GetIPPort returns the IPMI IP and port.
func (c *Client) GetIPPort() (ip string, port uint16, err error) {
	ipResp, err := c.ipmiClient.GetLanConfigParams(0x01, 0x03)
	if err != nil {
		return "", 0, err
	}

	portResp, err := c.ipmiClient.GetLanConfigParams(0x01, 0x08)
	if err != nil {
		return "", 0, err
	}

	ip = net.IP(ipResp.ConfigData).String()
	port = binary.LittleEndian.Uint16(portResp.ConfigData)

	return ip, port, nil
}

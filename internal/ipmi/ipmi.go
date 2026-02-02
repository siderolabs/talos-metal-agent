// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package ipmi implements various IPMI functions.
package ipmi

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bougou/go-ipmi"
	"go.uber.org/zap"
)

// Link to the IPMI spec: https://www.intel.com/content/dam/www/public/us/en/documents/product-briefs/ipmi-second-gen-interface-spec-v2-rev1-1.pdf

const channelNumber = uint8(0x01)

// Provider is an interface for the IPMI client operations used by this package.
type Provider interface {
	Close(ctx context.Context) error
	GetUserAccess(ctx context.Context, channelNumber uint8, userID uint8) (*ipmi.GetUserAccessResponse, error)
	GetUsername(ctx context.Context, userID uint8) (*ipmi.GetUsernameResponse, error)
	SetUsername(ctx context.Context, userID uint8, username string) (*ipmi.SetUsernameResponse, error)
	SetUserPassword(ctx context.Context, userID uint8, password string, test bool) (*ipmi.SetUserPasswordResponse, error)
	SetUserAccess(ctx context.Context, req *ipmi.SetUserAccessRequest) (*ipmi.SetUserAccessResponse, error)
	EnableUser(ctx context.Context, userID uint8) error
	GetUsers(ctx context.Context, channelNumber uint8) ([]*ipmi.User, error)
	GetLanConfigParamFor(ctx context.Context, channelNumber uint8, param ipmi.LanConfigParameter) error
}

// Client is a holder for the ipmiClient.
type Client struct {
	ipmiClient Provider
	logger     *zap.Logger
}

// NewClient creates a new Client with the given Provider and logger.
func NewClient(ipmiClient Provider, logger *zap.Logger) *Client {
	return &Client{ipmiClient: ipmiClient, logger: logger}
}

// NewLocalClient creates a new local ipmi client to use.
func NewLocalClient(ctx context.Context, logger *zap.Logger) (*Client, error) {
	ipmiClient, err := ipmi.NewOpenClient()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = ipmiClient.Connect(ctx); err != nil {
		return nil, err
	}

	return NewClient(ipmiClient, logger), nil
}

// Close the client.
func (c *Client) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.ipmiClient.Close(ctx)
}

// AttemptUserSetup attempts to set up an IPMI user with the given username.
func (c *Client) AttemptUserSetup(ctx context.Context, username, password string, logger *zap.Logger) error {
	userAccessResponse, err := c.ipmiClient.GetUserAccess(ctx, channelNumber, 0x01)
	if err != nil {
		return err
	}

	userExists := false
	userID := uint8(0)

	// Start from user ID 2, because 1 is the default admin user
	for i := uint8(2); i <= userAccessResponse.MaxUsersIDCount; i++ {
		userRes, userErr := c.ipmiClient.GetUsername(ctx, i)
		if userErr != nil { // This is an empty slot
			if userID == 0 { // We haven't found an empty slot yet
				userID = i // Note this empty slot
			}

			continue
		}

		if userRes.Username == username { // User already exists
			userExists = true
			userID = i

			logger.Info("user already present in slot, we'll claim it as our own", zap.Uint8("slot", i))

			break
		}

		trimmedUsername := strings.TrimSpace(userRes.Username)
		isEmptySlot := trimmedUsername == "" || trimmedUsername == "(Empty User)"

		if isEmptySlot && userID == 0 { // This slot is empty, and we haven't found an empty slot yet
			userID = i // Note this empty slot
		}
	}

	if userID == 0 { // No existing user found, and no empty slot available
		return errors.New("no slot available for user")
	}

	if !userExists { // Add our user to the empty slot
		logger.Info("adding user to slot", zap.Uint8("slot", userID))

		if _, err = c.ipmiClient.SetUsername(ctx, userID, username); err != nil {
			return err
		}
	}

	if _, err = c.ipmiClient.SetUserPassword(ctx, userID, password, false); err != nil {
		return err
	}

	if _, err = c.ipmiClient.SetUserAccess(ctx, &ipmi.SetUserAccessRequest{
		EnableChanging:      true,
		EnableIPMIMessaging: true,
		ChannelNumber:       uint8(ipmi.ChannelMediumIPMB),
		UserID:              userID,
		MaxPrivLevel:        uint8(ipmi.PrivilegeLevelAdministrator),
	}); err != nil {
		return fmt.Errorf("failed to set user access: %w", err)
	}

	if err = c.ipmiClient.EnableUser(ctx, userID); err != nil {
		return fmt.Errorf("failed to enable user: %w", err)
	}

	return nil
}

// UserExists checks if a user exists on the BMC.
func (c *Client) UserExists(ctx context.Context, username string) (bool, error) {
	users, err := c.ipmiClient.GetUsers(ctx, channelNumber)
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.Name == username {
			return true, nil
		}
	}

	return false, nil
}

// GetIPPort returns the IPMI IP and port.
func (c *Client) GetIPPort(ctx context.Context) (string, uint16, error) {
	var (
		ipParam   ipmi.LanConfigParam_IP
		portParam ipmi.LanConfigParam_PrimaryRMCPPort
	)

	if err := c.ipmiClient.GetLanConfigParamFor(ctx, channelNumber, &ipParam); err != nil {
		return "", 0, err
	}

	if err := c.ipmiClient.GetLanConfigParamFor(ctx, channelNumber, &portParam); err != nil {
		c.logger.Warn("failed to get IPMI port", zap.Error(err))
	}

	if portParam.Port == 0 {
		c.logger.Info("defaulting IPMI port to 623")

		portParam.Port = 623
	}

	return ipParam.IP.String(), portParam.Port, nil
}

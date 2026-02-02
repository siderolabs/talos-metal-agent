// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ipmi_test

import (
	"context"
	"errors"
	"net"
	"testing"

	goipmi "github.com/bougou/go-ipmi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"github.com/siderolabs/talos-metal-agent/internal/ipmi"
)

type mockProvider struct {
	getLanConfigParamForFunc func(ctx context.Context, channelNumber uint8, param goipmi.LanConfigParameter) error
}

func (m *mockProvider) Close(context.Context) error {
	return nil
}

func (m *mockProvider) GetUserAccess(context.Context, uint8, uint8) (*goipmi.GetUserAccessResponse, error) {
	return nil, nil //nolint:nilnil
}

func (m *mockProvider) GetUsername(context.Context, uint8) (*goipmi.GetUsernameResponse, error) {
	return nil, nil //nolint:nilnil
}

func (m *mockProvider) SetUsername(context.Context, uint8, string) (*goipmi.SetUsernameResponse, error) {
	return nil, nil //nolint:nilnil
}

func (m *mockProvider) SetUserPassword(context.Context, uint8, string, bool) (*goipmi.SetUserPasswordResponse, error) {
	return nil, nil //nolint:nilnil
}

func (m *mockProvider) SetUserAccess(context.Context, *goipmi.SetUserAccessRequest) (*goipmi.SetUserAccessResponse, error) {
	return nil, nil //nolint:nilnil
}

func (m *mockProvider) EnableUser(context.Context, uint8) error {
	return nil
}

func (m *mockProvider) GetUsers(context.Context, uint8) ([]*goipmi.User, error) {
	return nil, nil //nolint:nilnil
}

func (m *mockProvider) GetLanConfigParamFor(ctx context.Context, channelNumber uint8, param goipmi.LanConfigParameter) error {
	if m.getLanConfigParamForFunc != nil {
		return m.getLanConfigParamForFunc(ctx, channelNumber, param)
	}

	return nil
}

func TestGetIPPort(t *testing.T) {
	t.Parallel()

	testIP := net.IPv4(192, 168, 1, 100)

	tests := []struct {
		name         string
		mockFunc     func(ctx context.Context, channelNumber uint8, param goipmi.LanConfigParameter) error
		expectedIP   string
		expectedPort uint16
		expectedErr  bool
	}{
		{
			name: "both IP and port supported",
			mockFunc: func(_ context.Context, _ uint8, param goipmi.LanConfigParameter) error {
				switch p := param.(type) {
				case *goipmi.LanConfigParam_IP:
					p.IP = testIP
				case *goipmi.LanConfigParam_PrimaryRMCPPort:
					p.Port = 664
				}

				return nil
			},
			expectedIP:   "192.168.1.100",
			expectedPort: 664,
			expectedErr:  false,
		},
		{
			name: "port parameter not supported returns default 623",
			mockFunc: func(_ context.Context, _ uint8, param goipmi.LanConfigParameter) error {
				switch p := param.(type) {
				case *goipmi.LanConfigParam_IP:
					p.IP = testIP
				case *goipmi.LanConfigParam_PrimaryRMCPPort:
					return errors.New("parameter not supported")
				}

				return nil
			},
			expectedIP:   "192.168.1.100",
			expectedPort: 623,
			expectedErr:  false,
		},
		{
			name: "port returns zero defaults to 623",
			mockFunc: func(_ context.Context, _ uint8, param goipmi.LanConfigParameter) error {
				switch p := param.(type) {
				case *goipmi.LanConfigParam_IP:
					p.IP = testIP
				case *goipmi.LanConfigParam_PrimaryRMCPPort:
					p.Port = 0
				}

				return nil
			},
			expectedIP:   "192.168.1.100",
			expectedPort: 623,
			expectedErr:  false,
		},
		{
			name: "IP parameter fails returns error",
			mockFunc: func(_ context.Context, _ uint8, param goipmi.LanConfigParameter) error {
				if _, ok := param.(*goipmi.LanConfigParam_IP); ok {
					return errors.New("failed to get IP")
				}

				return nil
			},
			expectedIP:   "",
			expectedPort: 0,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := &mockProvider{
				getLanConfigParamForFunc: tt.mockFunc,
			}

			logger := zaptest.NewLogger(t)
			client := ipmi.NewClient(mock, logger)

			ip, port, err := client.GetIPPort(t.Context())

			if tt.expectedErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expectedIP, ip)
			assert.Equal(t, tt.expectedPort, port)
		})
	}
}

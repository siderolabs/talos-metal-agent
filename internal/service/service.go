// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package service provides the agent GRPC service server.
package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/cosi-project/runtime/pkg/safe"
	"github.com/cosi-project/runtime/pkg/state"
	"github.com/siderolabs/talos/pkg/machinery/api/machine"
	"github.com/siderolabs/talos/pkg/machinery/api/storage"
	talosclient "github.com/siderolabs/talos/pkg/machinery/client"
	"github.com/siderolabs/talos/pkg/machinery/resources/block"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	agentpb "github.com/siderolabs/talos-metal-agent/api/agent"
)

// IPMIClientFactory is the factory to create IPMI clients.
type IPMIClientFactory func(ctx context.Context) (IPMIClient, error)

// IPMIClient represents an IPMI client.
type IPMIClient interface {
	io.Closer

	// UserExists checks if the user exists.
	UserExists(ctx context.Context, username string) (bool, error)

	// AttemptUserSetup attempts to set up the BMC user.
	AttemptUserSetup(ctx context.Context, username, password string, logger *zap.Logger) error

	// GetIPPort returns the BMC IP and port.
	GetIPPort(ctx context.Context) (string, uint16, error)
}

// TalosClient represents a Talos API client.
type TalosClient interface {
	Reboot(ctx context.Context, opts ...talosclient.RebootMode) error
	State() state.State
	BlockDeviceWipe(ctx context.Context, req *storage.BlockDeviceWipeRequest, callOptions ...grpc.CallOption) error
	MDDestroy(ctx context.Context, req *machine.MDDestroyRequest, callOptions ...grpc.CallOption) error
}

// Server is the agent service server.
type Server struct {
	agentpb.UnimplementedAgentServiceServer

	talosClient       TalosClient
	ipmiClientFactory IPMIClientFactory

	logger *zap.Logger

	sf singleflight.Group

	testMode bool
}

// NewServer creates a new service server.
func NewServer(talosClient TalosClient, ipmiClientFactory IPMIClientFactory, testMode bool, logger *zap.Logger) *Server {
	return &Server{
		talosClient:       talosClient,
		ipmiClientFactory: ipmiClientFactory,
		logger:            logger,
		testMode:          testMode,
	}
}

// Hello is an endpoint to check if the service is available.
func (s *Server) Hello(_ context.Context, _ *agentpb.HelloRequest) (*agentpb.HelloResponse, error) {
	s.logger.Debug("hello", zap.Bool("test_mode", s.testMode))

	return &agentpb.HelloResponse{}, nil
}

// GetPowerManagement returns the power management info.
func (s *Server) GetPowerManagement(ctx context.Context, req *agentpb.GetPowerManagementRequest) (*agentpb.GetPowerManagementResponse, error) {
	s.logger.Debug("get power management", zap.Bool("test_mode", s.testMode))

	return runSingleflight[*agentpb.GetPowerManagementResponse](ctx, agentpb.AgentService_GetPowerManagement_FullMethodName, &s.sf, req, func() (*agentpb.GetPowerManagementResponse, error) {
		if s.testMode {
			return &agentpb.GetPowerManagementResponse{
				Api: &agentpb.GetPowerManagementResponse_API{},
			}, nil
		}

		ipmiClient, err := s.ipmiClientFactory(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error creating ipmi client: %v", err)
		}

		defer ipmiClient.Close() //nolint:errcheck

		ip, port, err := ipmiClient.GetIPPort(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error getting bmc ip port: %v", err)
		}

		checkUsername := req.GetIpmi().GetCheckUsername()

		exists, err := ipmiClient.UserExists(ctx, checkUsername)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error checking if user %q exists: %v", checkUsername, err)
		}

		return &agentpb.GetPowerManagementResponse{
			Ipmi: &agentpb.GetPowerManagementResponse_IPMI{
				Address:    ip,
				Port:       uint32(port),
				UserExists: exists,
			},
		}, nil
	})
}

// SetPowerManagement sets the power management info.
func (s *Server) SetPowerManagement(ctx context.Context, req *agentpb.SetPowerManagementRequest) (*agentpb.SetPowerManagementResponse, error) {
	s.logger.Debug("set power management", zap.Bool("test_mode", s.testMode), zap.String("ipmi_username", req.GetIpmi().GetUsername()))

	return runSingleflight[*agentpb.SetPowerManagementResponse](ctx, agentpb.AgentService_SetPowerManagement_FullMethodName, &s.sf, req, func() (*agentpb.SetPowerManagementResponse, error) {
		if s.testMode {
			return &agentpb.SetPowerManagementResponse{}, nil
		}

		ipmiClient, err := s.ipmiClientFactory(ctx)
		if err != nil {
			return nil, fmt.Errorf("error creating ipmi client: %w", err)
		}

		defer ipmiClient.Close() //nolint:errcheck

		if err = ipmiClient.AttemptUserSetup(ctx, req.GetIpmi().GetUsername(), req.GetIpmi().GetPassword(), s.logger); err != nil {
			return nil, fmt.Errorf("failed to set up IPMI user: %w", err)
		}

		return &agentpb.SetPowerManagementResponse{}, nil
	})
}

// Reboot reboots the machine.
func (s *Server) Reboot(ctx context.Context, req *agentpb.RebootRequest) (*agentpb.RebootResponse, error) {
	s.logger.Info("reboot")

	return runSingleflight[*agentpb.RebootResponse](ctx, agentpb.AgentService_Reboot_FullMethodName, &s.sf, req, func() (*agentpb.RebootResponse, error) {
		if err := s.talosClient.Reboot(ctx, talosclient.WithPowerCycle); err != nil {
			return nil, err
		}

		return &agentpb.RebootResponse{}, nil
	})
}

// WipeDisks wipes the disks.
func (s *Server) WipeDisks(ctx context.Context, req *agentpb.WipeDisksRequest) (*agentpb.WipeDisksResponse, error) {
	s.logger.Info("wipe disks", zap.Bool("zeroes", req.Zeroes), zap.Bool("test_mode", s.testMode))

	return runSingleflight[*agentpb.WipeDisksResponse](ctx, agentpb.AgentService_WipeDisks_FullMethodName, &s.sf, req, func() (*agentpb.WipeDisksResponse, error) {
		method := storage.BlockDeviceWipeDescriptor_FAST
		if req.Zeroes {
			method = storage.BlockDeviceWipeDescriptor_ZEROES
		}

		s.destroyMDArrays(ctx)

		diskList, err := safe.StateListAll[*block.Disk](ctx, s.talosClient.State())
		if err != nil {
			return nil, fmt.Errorf("failed to list disks: %w", err)
		}

		deviceNames := make([]string, 0, diskList.Len())
		devices := make([]*storage.BlockDeviceWipeDescriptor, 0, diskList.Len())

		for disk := range diskList.All() {
			if disk.TypedSpec().Readonly || disk.TypedSpec().CDROM {
				continue
			}

			deviceNames = append(deviceNames, disk.Metadata().ID())
			devices = append(devices, &storage.BlockDeviceWipeDescriptor{
				Device:          disk.Metadata().ID(),
				Method:          method,
				SkipVolumeCheck: true,
			})
		}

		s.logger.Debug("going to wipe disks", zap.Strings("devices", deviceNames))

		if err = s.talosClient.BlockDeviceWipe(ctx, &storage.BlockDeviceWipeRequest{
			Devices: devices,
		}); err != nil {
			return nil, fmt.Errorf("failed to wipe disks: %w", err)
		}

		return &agentpb.WipeDisksResponse{}, nil
	})
}

// mdMajor is the block device major number of MD (software RAID) arrays.
const mdMajor = 9

// destroyMDArrays stops every MD array and clears the superblocks of its members.
//
// Talos assembles the arrays it finds on the disks at boot, and it refuses to wipe a disk that is a member of an array.
// Arrays are looked up among the block devices rather than the disks, because an incomplete array is not a disk, but still holds its members.
//
// Failures are only logged: if an array that could not be destroyed still holds a disk, wiping that disk fails.
func (s *Server) destroyMDArrays(ctx context.Context) {
	deviceList, err := safe.StateListAll[*block.Device](ctx, s.talosClient.State())
	if err != nil {
		s.logger.Warn("failed to list block devices", zap.Error(err))

		return
	}

	for device := range deviceList.All() {
		if device.TypedSpec().Major != mdMajor {
			continue
		}

		s.logger.Info("destroy md array", zap.String("device", device.Metadata().ID()))

		if err = s.talosClient.MDDestroy(ctx, &machine.MDDestroyRequest{
			Device: "/dev/" + device.Metadata().ID(),
		}); err != nil {
			s.logger.Warn("failed to destroy md array", zap.String("device", device.Metadata().ID()), zap.Error(err))
		}
	}
}

type marshaler interface {
	MarshalVT() ([]byte, error)
}

func runSingleflight[T any](ctx context.Context, keyPrefix string, sf *singleflight.Group, req marshaler, fn func() (T, error)) (T, error) {
	var zero T

	b, err := req.MarshalVT()
	if err != nil {
		return zero, err
	}

	hash := md5.Sum(b)
	key := keyPrefix + "-" + hex.EncodeToString(hash[:])

	ch := sf.DoChan(key, func() (any, error) {
		return fn()
	})

	select {
	case <-ctx.Done():
		return zero, ctx.Err()
	case v := <-ch:
		if v.Err != nil {
			return zero, v.Err
		}

		result, ok := v.Val.(T)
		if !ok {
			return zero, fmt.Errorf("unexpected type: %T", v.Val)
		}

		return result, nil
	}
}

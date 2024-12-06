// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package agent contains Talos metal agent mode functionality.
package agent

import (
	"context"
	"fmt"

	"github.com/cosi-project/runtime/pkg/safe"
	"github.com/cosi-project/runtime/pkg/state"
	"github.com/jhump/grpctunnel"
	"github.com/jhump/grpctunnel/tunnelpb"
	"github.com/siderolabs/talos/pkg/grpc/middleware/authz"
	talosclient "github.com/siderolabs/talos/pkg/machinery/client"
	talosconstants "github.com/siderolabs/talos/pkg/machinery/constants"
	"github.com/siderolabs/talos/pkg/machinery/resources/hardware"
	talosrole "github.com/siderolabs/talos/pkg/machinery/role"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	agentpb "github.com/siderolabs/talos-metal-agent/api/agent"
	"github.com/siderolabs/talos-metal-agent/internal/ipmi"
	"github.com/siderolabs/talos-metal-agent/internal/service"
	"github.com/siderolabs/talos-metal-agent/pkg/constants"
)

// Agent is the Talos agent.
type Agent struct {
	logger          *zap.Logger
	providerAddress string
	testMode        bool
}

// New creates a new agent.
func New(providerAddress string, testMode bool, logger *zap.Logger) (*Agent, error) {
	return &Agent{
		providerAddress: providerAddress,
		testMode:        testMode,
		logger:          logger,
	}, nil
}

// Run starts the agent.
func (a *Agent) Run(ctx context.Context) error {
	a.logger.Info("running metal agent", zap.String("provider_address", a.providerAddress), zap.Bool("test_mode", a.testMode))

	talosClient, err := buildTalosClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to build Talos client: %w", err)
	}

	versionResponse, err := talosClient.Version(ctx)
	if err != nil {
		return fmt.Errorf("failed to read Talos version: %w", err)
	}

	systemInformation, err := safe.StateGetByID[*hardware.SystemInformation](ctx, talosClient.COSI, hardware.SystemInformationID)
	if err != nil {
		return fmt.Errorf("failed to read system information: %w", err)
	}

	machineID := systemInformation.TypedSpec().UUID

	a.logger.Info("connected to Talos", zap.String("version", versionResponse.Messages[0].String()), zap.String("machine_uuid", machineID))

	providerDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(idHeaderUnaryInterceptor(machineID)),
		grpc.WithStreamInterceptor(idHeaderStreamInterceptor(machineID)),
	}

	providerConn, err := grpc.NewClient(a.providerAddress, providerDialOptions...)
	if err != nil {
		return fmt.Errorf("failed to create grpc client: %w", err)
	}

	tunnelStub := tunnelpb.NewTunnelServiceClient(providerConn)
	channelServer := grpctunnel.NewReverseTunnelServer(tunnelStub)

	ipmiClientFactory := func() (service.IPMIClient, error) {
		return ipmi.NewLocalClient()
	}

	serviceServer := service.NewServer(talosClient, ipmiClientFactory, a.testMode, a.logger)

	agentpb.RegisterAgentServiceServer(channelServer, serviceServer)

	// Open the reverse tunnel and serve requests.
	if _, err = channelServer.Serve(ctx); err != nil {
		return fmt.Errorf("failed to serve over grpc tunnel: %w", err)
	}

	return nil
}

type talosClientWrapper struct {
	*talosclient.Client
}

func (c talosClientWrapper) State() state.State {
	return c.COSI
}

func buildTalosClient(ctx context.Context) (talosClientWrapper, error) {
	opts := []talosclient.OptionFunc{
		talosclient.WithUnixSocket(talosconstants.MachineSocketPath),
		talosclient.WithGRPCDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(talosAdminAuthzUnaryInterceptor),
			grpc.WithStreamInterceptor(talosAdminAuthzStreamInterceptor)),
	}

	client, err := talosclient.New(ctx, opts...)
	if err != nil {
		return talosClientWrapper{}, fmt.Errorf("failed to construct client: %w", err)
	}

	return talosClientWrapper{client}, nil
}

func idHeaderUnaryInterceptor(id string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = metadata.AppendToOutgoingContext(ctx, constants.MachineIDMetadataKey, id)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func idHeaderStreamInterceptor(id string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, constants.MachineIDMetadataKey, id)

		return streamer(ctx, desc, cc, method, opts...)
	}
}

var (
	talosAdminAuthzUnaryInterceptor grpc.UnaryClientInterceptor = func(ctx context.Context, method string, req, reply any,
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		ctx = injectAdminRole(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}

	talosAdminAuthzStreamInterceptor grpc.StreamClientInterceptor = func(ctx context.Context, desc *grpc.StreamDesc,
		cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx = injectAdminRole(ctx)

		return streamer(ctx, desc, cc, method, opts...)
	}
)

func injectAdminRole(ctx context.Context) context.Context {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	authz.SetMetadata(md, talosrole.MakeSet(talosrole.Admin))

	return metadata.NewOutgoingContext(ctx, md)
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package service_test

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/cosi-project/runtime/pkg/state"
	"github.com/cosi-project/runtime/pkg/state/impl/inmem"
	"github.com/cosi-project/runtime/pkg/state/impl/namespaced"
	"github.com/siderolabs/talos/pkg/machinery/api/machine"
	"github.com/siderolabs/talos/pkg/machinery/api/storage"
	talosclient "github.com/siderolabs/talos/pkg/machinery/client"
	"github.com/siderolabs/talos/pkg/machinery/resources/block"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc"

	agentpb "github.com/siderolabs/talos-metal-agent/api/agent"
	"github.com/siderolabs/talos-metal-agent/internal/service"
)

// talosClient mimics Talos: destroying an array removes its block device and disk.
type talosClient struct {
	st            state.State
	failToDestroy string
	calls         []string
}

func (c *talosClient) Reboot(context.Context, ...talosclient.RebootMode) error {
	return nil
}

func (c *talosClient) State() state.State {
	return c.st
}

func (c *talosClient) BlockDeviceWipe(_ context.Context, req *storage.BlockDeviceWipeRequest, _ ...grpc.CallOption) error {
	for _, device := range req.GetDevices() {
		c.calls = append(c.calls, "wipe "+device.GetDevice())
	}

	return nil
}

func (c *talosClient) MDDestroy(ctx context.Context, req *machine.MDDestroyRequest, _ ...grpc.CallOption) error {
	c.calls = append(c.calls, "destroy "+req.GetDevice())

	if req.GetDevice() == c.failToDestroy {
		return errors.New("destroy failed")
	}

	id := filepath.Base(req.GetDevice())

	if err := c.st.Destroy(ctx, block.NewDevice(block.NamespaceName, id).Metadata()); err != nil {
		return err
	}

	if err := c.st.Destroy(ctx, block.NewDisk(block.NamespaceName, id).Metadata()); err != nil && !state.IsNotFoundError(err) {
		return err
	}

	return nil
}

func TestWipeDisksDestroysMDArrays(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	st := state.WrapCore(namespaced.NewState(inmem.Build))

	for id, major := range map[string]int{
		"vda":   253,
		"vdb":   253,
		"vdc":   253,
		"md127": 9, // assembled array with vda and vdb
		"md126": 9, // incomplete array with vdc, which Talos does not report as a disk
	} {
		device := block.NewDevice(block.NamespaceName, id)
		device.TypedSpec().Type = block.DeviceTypeDisk
		device.TypedSpec().Major = major

		require.NoError(t, st.Create(ctx, device))
	}

	for _, id := range []string{"vda", "vdb", "vdc", "md127"} {
		require.NoError(t, st.Create(ctx, block.NewDisk(block.NamespaceName, id)))
	}

	// a failed destroy must not stop the other arrays from being destroyed or the disks from being wiped
	client := &talosClient{st: st, failToDestroy: "/dev/md126"}
	server := service.NewServer(client, nil, false, zaptest.NewLogger(t))

	_, err := server.WipeDisks(ctx, &agentpb.WipeDisksRequest{})
	require.NoError(t, err)
	require.Len(t, client.calls, 5)

	assert.ElementsMatch(t, []string{"destroy /dev/md126", "destroy /dev/md127"}, client.calls[:2])
	assert.ElementsMatch(t, []string{"wipe vda", "wipe vdb", "wipe vdc"}, client.calls[2:])
}

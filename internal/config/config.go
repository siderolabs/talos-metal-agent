// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package config contains the configuration for the agent.
package config

import (
	"strconv"

	"github.com/siderolabs/go-procfs/procfs"
	"go.uber.org/zap"

	"github.com/siderolabs/talos-metal-agent/pkg/config"
)

// Config contains the configuration for the agent.
type Config struct {
	ProviderAddress string
	TestMode        bool
	TLSSkipVerify   bool
}

// LoadFromKernelCmdline loads the Config from the kernel arguments.
func LoadFromKernelCmdline(logger *zap.Logger) Config {
	var providerAddress string

	cmdline := procfs.ProcCmdline()

	providerAddressParam := cmdline.Get(config.MetalProviderAddressKernelArg)
	if providerAddressParam != nil {
		providerAddressVal := providerAddressParam.First()
		if providerAddressVal != nil {
			providerAddress = *providerAddressVal
		}
	}

	testMode := parseBooleanFlag(cmdline, config.TestModeKernelArg, logger)
	tlsSkipVerify := parseBooleanFlag(cmdline, config.TLSSkipVerifyKernelArg, logger)

	return Config{
		ProviderAddress: providerAddress,
		TestMode:        testMode,
		TLSSkipVerify:   tlsSkipVerify,
	}
}

func parseBooleanFlag(cmdline *procfs.Cmdline, key string, logger *zap.Logger) bool {
	var val bool

	param := cmdline.Get(key)
	if param != nil {
		testModeVal := param.First()
		if testModeVal != nil {
			var err error

			val, err = strconv.ParseBool(*testModeVal)
			if err != nil {
				logger.Error("failed to parse bool arg", zap.String("key", key), zap.String("value", *testModeVal), zap.Error(err))
			}
		}
	}

	return val
}

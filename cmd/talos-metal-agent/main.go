// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package main contains the entrypoint for the Talos metal agent.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/siderolabs/talos-metal-agent/internal/agent"
	internalconfig "github.com/siderolabs/talos-metal-agent/internal/config"
	"github.com/siderolabs/talos-metal-agent/internal/version"
	"github.com/siderolabs/talos-metal-agent/pkg/config"
)

const testModeFlag = "test-mode"

var rootCmdArgs struct {
	providerAddress string
	testMode        bool
	debug           bool
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     version.Name,
	Short:   "Run the Talos metal agent",
	Version: version.Tag,
	Args:    cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		cmd.SilenceUsage = true // if the args are parsed fine, no need to show usage
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		logger, err := initLogger()
		if err != nil {
			return fmt.Errorf("failed to create logger: %w", err)
		}

		conf := internalconfig.LoadFromKernelCmdline(logger)

		if rootCmdArgs.providerAddress != "" {
			conf.ProviderAddress = rootCmdArgs.providerAddress
		}

		if cmd.Flags().Changed(testModeFlag) {
			conf.TestMode = rootCmdArgs.testMode
		}

		defer logger.Sync() //nolint:errcheck

		return run(cmd.Context(), conf.ProviderAddress, conf.TestMode, conf.TLSSkipVerify, logger)
	},
}

func run(ctx context.Context, providerAddress string, testMode, tlsSkipVerify bool, logger *zap.Logger) error {
	ag, err := agent.New(providerAddress, testMode, tlsSkipVerify, logger)
	if err != nil {
		return fmt.Errorf("failed to create agent: %w", err)
	}

	if err = ag.Run(ctx); err != nil {
		return fmt.Errorf("failed to run agent: %w", err)
	}

	return nil
}

func initLogger() (*zap.Logger, error) {
	var loggerConfig zap.Config

	if rootCmdArgs.debug {
		loggerConfig = zap.NewDevelopmentConfig()
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		loggerConfig.Level.SetLevel(zap.DebugLevel)
	} else {
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.Level.SetLevel(zap.InfoLevel)
	}

	return loggerConfig.Build()
}

func main() {
	if err := runCmd(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func runCmd() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.Flags().StringVar(&rootCmdArgs.providerAddress, "provider-address", "", fmt.Sprintf(
		"The infra provider address to connect to. If not specified explicitly, the value of the kernel arg %q will be used.", config.MetalProviderAddressKernelArg))
	rootCmd.Flags().BoolVar(&rootCmdArgs.testMode, testModeFlag, false, "Enable test mode. In this mode, "+
		"the agent will assume that the power management is done via an external API (e.g., the power API served by 'talosctl cluster create').")
	rootCmd.Flags().BoolVar(&rootCmdArgs.debug, "debug", false, "Enable debug mode & logs.")
}

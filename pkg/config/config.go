// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package config contains public configuration options for the agent.
package config

const (
	// MetalProviderAddressKernelArg is the kernel argument that contains the provider address.
	MetalProviderAddressKernelArg = "metal.provider.address"

	// TestModeKernelArg is the kernel argument that contains the test mode flag.
	TestModeKernelArg = "metal.provider.test.mode"

	// TLSSkipVerifyKernelArg is the kernel argument to disable TLS verification in TLS connection to the metal provider.
	TLSSkipVerifyKernelArg = "metal.provider.tls.skip.verify"
)

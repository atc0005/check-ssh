// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"strings"

	"github.com/atc0005/check-ssh/internal/textutils"
)

// validate verifies all Config struct fields have been provided acceptable
// values.
func (c Config) validate(_ AppType) error {
	if strings.TrimSpace(c.Server) == "" {
		return fmt.Errorf(
			"%w: missing server FQDN or IP Address",
			ErrUnsupportedOption,
		)
	}

	if strings.TrimSpace(c.Username) == "" {
		return fmt.Errorf(
			"%w: missing username",
			ErrUnsupportedOption,
		)
	}

	if strings.TrimSpace(c.Password) == "" {
		return fmt.Errorf(
			"%w: missing password",
			ErrUnsupportedOption,
		)
	}

	// FIXME: Should this be optional?
	//
	// if strings.TrimSpace(c.SSHCommand) == "" {
	// 	return fmt.Errorf(
	// 		"%w: missing command to execute via SSH",
	// 		ErrUnsupportedOption,
	// 	)
	// }

	// TCP Port 0 is used by server applications to indicate that they should
	// bind to an available port. Specifying port 0 for a client application
	// is not useful.
	if c.TCPPort <= 0 {
		return fmt.Errorf(
			"%w: invalid TCP port number %d",
			ErrUnsupportedOption,
			c.TCPPort,
		)
	}

	// Validate the specified network type
	supportedNetworkTypes := supportedNetworkTypes()
	if !textutils.InList(c.NetworkType, supportedNetworkTypes, true) {
		return fmt.Errorf(
			"%w: invalid network type;"+
				" got %v, expected one of %v",
			ErrUnsupportedOption,
			c.NetworkType,
			supportedNetworkTypes,
		)
	}

	// Validate the specified logging level
	supportedLogLevels := supportedLogLevels()
	if !textutils.InList(c.LoggingLevel, supportedLogLevels, true) {
		return fmt.Errorf(
			"%w: invalid logging level;"+
				" got %v, expected one of %v",
			ErrUnsupportedOption,
			c.LoggingLevel,
			supportedLogLevels,
		)
	}

	supportedPluginStateLabels := supportedPluginStateLabels()
	if !textutils.InList(c.LoginFailureState, supportedPluginStateLabels, true) {
		return fmt.Errorf(
			"%w: invalid plugin state for %q flag;"+
				" got %q, expected one of %q",
			ErrUnsupportedOption,
			LoginFailureStateFlagLong,
			c.LoginFailureState,
			supportedPluginStateLabels,
		)
	}

	if !textutils.InList(c.LoginSuccessState, supportedPluginStateLabels, true) {
		return fmt.Errorf(
			"%w: invalid plugin state for %q flag;"+
				" got %q, expected one of %q",
			ErrUnsupportedOption,
			LoginSuccessStateFlagLong,
			c.LoginSuccessState,
			supportedPluginStateLabels,
		)
	}

	// Specifying the same plugin state for a successful SSH login and a
	// failed SSH login is unsupported.
	//
	// FIXME: Do we really need to assert that these are not equal?
	//
	// if strings.EqualFold(c.LoginSuccessState, c.LoginFailureState) {
	// 	return fmt.Errorf(
	// 		"%w: value %q given for both of flag %q and flag %q",
	// 		ErrUnsupportedOption,
	// 		c.LoginSuccessState,
	// 		LoginSuccessStateFlagLong,
	// 		LoginFailureStateFlagLong,
	// 	)
	// }

	// 	switch {
	// 	case appType.Inspector:
	//
	// 	case appType.Plugin:
	//
	// 	}

	// Optimist
	return nil
}

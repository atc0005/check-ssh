// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"strings"

	"github.com/atc0005/go-nagios"
)

// supportedLogLevels returns a list of valid log levels supported by tools in
// this project.
func supportedLogLevels() []string {
	return []string{
		LogLevelDisabled,
		LogLevelPanic,
		LogLevelFatal,
		LogLevelError,
		LogLevelWarn,
		LogLevelInfo,
		LogLevelDebug,
		LogLevelTrace,
	}
}

// supportedNetworkTypes returns a list of valid network types.
func supportedNetworkTypes() []string {
	return []string{
		netTypeTCPAuto,
		netTypeTCP4,
		netTypeTCP6,
	}
}

// supportedPluginStateLabels returns a list of valid plugin states for
// plugins in this project.
//
// While state labels are traditionally listed in all caps (e.g., WARNING), we
// provide the values in lowercase as a convenience for display in help text
// output. Comparison logic which use elements from the provided function is
// responsible for normalizing all compared values.
func supportedPluginStateLabels() []string {
	stateLabels := nagios.SupportedStateLabels()
	for i := range stateLabels {
		stateLabels[i] = strings.ToLower(stateLabels[i])
	}

	return stateLabels
}

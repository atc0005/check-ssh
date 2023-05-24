// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"strings"

	"github.com/atc0005/check-ssh/internal/config"
	"github.com/atc0005/check-ssh/internal/sshutils"
	"github.com/atc0005/go-nagios"
	"golang.org/x/crypto/ssh"
)

// processAuthError is responsible for processing a SSH authentication error
// given the specified configuration settings and setting the plugin exit
// state accordingly.
func processAuthError(err error, client *ssh.Client, cfg *config.Config, plugin *nagios.Plugin) {
	logger := cfg.Log.With().Logger()

	switch {
	case errors.Is(err, sshutils.ErrMissingValue):
		// FIXME: Should this be a panic instead? This scenario is not
		// expected to occur as config validation should prevent us from
		// arriving at this point without all required values.
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("server", cfg.Username).
			Int("port", cfg.TCPPort).
			Str("net_type", cfg.NetworkType).
			Msg("Failed to provide all required options")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"SSH client configuration failure",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return

	case errors.Is(err, sshutils.ErrIPAddressParsingFailed):
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("net_type", cfg.NetworkType).
			Msg("Failed to parse given server value as IP Address")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"Failed to parse given server value as IP Address",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return

	// TODO: Confirm behavior when giving IP Address as SSH server host
	// value; how does net.LookupHost handle bare IP Addresses? Do we need
	// to match format and skip name resolution attempt?
	case errors.Is(err, sshutils.ErrDNSLookupFailed):
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("net_type", cfg.NetworkType).
			Msg("Failed to resolve server hostname")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"SSH server hostname resolution failure",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return

	case errors.Is(err, sshutils.ErrNoIPAddressesForChosenNetworkType):
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("net_type", cfg.NetworkType).
			Msg("Failed to resolve server hostname to an IP Address of the chosen network type")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"SSH server hostname resolution failure",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return

	case errors.Is(err, sshutils.ErrNetworkConnectionFailed):
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("server", cfg.Username).
			Int("port", cfg.TCPPort).
			Str("net_type", cfg.NetworkType).
			Msg("Failed to establish network connection to SSH server")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"SSH connection failure",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return

	case errors.Is(err, sshutils.ErrAuthenticationFailed):
		// Conditional exit state depending on what the user requested
		// when SSH authentication fails for a given user.
		switch {
		case strings.EqualFold(cfg.LoginFailureState, nagios.StateOKLabel):
			logger.Debug().
				Err(err).
				Str("server", cfg.Server).
				Str("server", cfg.Username).
				Int("port", cfg.TCPPort).
				Str("net_type", cfg.NetworkType).
				Msg("Failed to establish network connection to SSH server (as expected)")

			setPluginOutput(
				cfg.LoginFailureState,
				"SSH authentication failed (as expected)",
				"",
				err,
				client,
				cfg,
				plugin,
			)

			return

		default:
			logger.Error().
				Err(err).
				Str("server", cfg.Server).
				Str("server", cfg.Username).
				Int("port", cfg.TCPPort).
				Str("net_type", cfg.NetworkType).
				Msg("Failed to establish network connection to SSH server")

			setPluginOutput(
				cfg.LoginFailureState,
				"SSH authentication failed",
				"",
				err,
				client,
				cfg,
				plugin,
			)

			return
		}

	case errors.Is(err, sshutils.ErrClientSessionFailed):
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("server", cfg.Username).
			Int("port", cfg.TCPPort).
			Str("net_type", cfg.NetworkType).
			Msg("Failed to establish SSH client session")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"SSH client session failure",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return

	default:
		// FIXME: Replace "General" with something more specific that
		// indicates that an error occurred, but that we do not have a
		// more specific cause.
		logger.Error().
			Err(err).
			Str("server", cfg.Server).
			Str("server", cfg.Username).
			Int("port", cfg.TCPPort).
			Str("net_type", cfg.NetworkType).
			Msg("General SSH connection failure")

		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"General SSH connection failure",
			"",
			err,
			client,
			cfg,
			plugin,
		)

		return
	}
}

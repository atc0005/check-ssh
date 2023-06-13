// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/atc0005/check-ssh/internal/config"
	"github.com/atc0005/check-ssh/internal/sshutils"
	"github.com/atc0005/go-nagios"
	"github.com/rs/zerolog"
)

func main() {
	plugin := nagios.NewPlugin()

	// defer this from the start so it is the last deferred function to run
	defer plugin.ReturnCheckResults()

	// Setup configuration by parsing user-provided flags.
	cfg, cfgErr := config.New(config.AppType{Plugin: true})

	switch {
	case errors.Is(cfgErr, config.ErrVersionRequested):
		fmt.Println(config.Version())

		return

	case errors.Is(cfgErr, config.ErrHelpRequested):
		fmt.Println(cfg.Help())

		return

	case cfgErr != nil:
		// We make some assumptions when setting up our logger as we do not
		// have a working configuration based on sysadmin-specified choices.
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true}
		logger := zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()

		logger.Err(cfgErr).Msg("Error initializing application")

		// plugin.ServiceOutput = fmt.Sprintf(
		// 	"%s: Error initializing application",
		// 	nagios.StateUNKNOWNLabel,
		// )
		// plugin.AddError(cfgErr)
		// plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
		setPluginOutput(
			nagios.StateUNKNOWNLabel,
			"Error initializing application",
			"",
			cfgErr,
			nil,
			cfg,
			plugin,
		)

		return
	}

	if cfg.EmitBranding {
		// If enabled, show application details at end of notification
		plugin.BrandingCallback = config.Branding("Notification generated by ")
	}

	logger := cfg.Log.With().Logger()

	sshPasswordAuthConfig := sshutils.SSHPasswordAuthConfig{
		Server:        cfg.Server,
		Port:          cfg.TCPPort,
		Username:      cfg.Username,
		Password:      cfg.Password,
		NetworkType:   cfg.NetworkType,
		ClientVersion: config.SSHClientVersion(),
	}

	client, session, loginErr := sshutils.LoginWithPasswordAuth(
		sshPasswordAuthConfig,
		logger,
	)

	if loginErr != nil {
		processAuthError(loginErr, nil, cfg, plugin)

		// The plugin exit state has been set based on the specific error
		// encountered and the sysadmin specified plugin configuration.
		return
	}

	defer closeClient(client, cfg, plugin)

	defer closeSession(session, cfg, plugin)

	// At this point we have successfully authenticated. Per given
	// configuration settings, this may not be desired.
	//
	// Apply conditional exit state depending on what the user requested.
	if !strings.EqualFold(cfg.LoginSuccessState, nagios.StateOKLabel) {
		logger.Error().
			Err(sshutils.ErrAuthenticationSucceeded).
			Str("server", cfg.Server).
			Str("server", cfg.Username).
			Int("port", cfg.TCPPort).
			Str("net_type", cfg.NetworkType).
			Msg("Successfully authenticated to SSH server (UNEXPECTED)")

		setPluginOutput(
			cfg.LoginSuccessState,
			"SSH authentication successful (UNEXPECTED)",
			"",
			fmt.Errorf(
				"expected login failure: %w",
				sshutils.ErrAuthenticationSucceeded,
			),
			client,
			cfg,
			plugin,
		)

		return
	}

	logger.Debug().
		Str("server", cfg.Server).
		Str("server", cfg.Username).
		Int("port", cfg.TCPPort).
		Str("net_type", cfg.NetworkType).
		Str("client_version", string(client.Conn.ClientVersion())).
		Str("server_version", string(client.Conn.ServerVersion())).
		Msg("Successfully authenticated to SSH server (as expected)")

	// NOTE: SSH command execution is skipped if:
	//
	// - command not specified
	// - login failed
	// - login successful AND this is considered a non-OK result
	//
	// TODO: Document this behavior.
	switch {
	case cfg.SSHCommand != "":
		result, sshCmdErr := processSSHCommand(session, cfg)
		if sshCmdErr != nil {
			// TODO: Determine how we will handle SSH command failure.
			//
			// In what scenario is an error expected/OK? If we don't fail at
			// successful login, is the SSH command execution expected?
			logger.Error().Err(sshCmdErr).Msg("SSH command execution failure")

			sshCommandUsed := fmt.Sprintf(
				"SSH command: %q%s%s",
				cfg.SSHCommand,
				nagios.CheckOutputEOL,
				nagios.CheckOutputEOL,
			)

			setPluginOutput(
				nagios.StateUNKNOWNLabel,
				"SSH command execution failure",
				sshCommandUsed,
				sshCmdErr,
				client,
				cfg,
				plugin,
			)

			return
		}

		logger.Debug().Msg("SSH command execution successful (as expected)")
		setPluginOutput(
			cfg.LoginSuccessState,
			"SSH authentication and command execution successful (as expected)",
			result,
			nil,
			client,
			cfg,
			plugin,
		)

	default:
		setPluginOutput(
			cfg.LoginSuccessState,
			"SSH authentication successful (as expected)",
			"NOTE: SSH command execution not requested.",
			nil,
			client,
			cfg,
			plugin,
		)
	}
}

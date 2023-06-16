// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"strings"

	"github.com/atc0005/check-ssh/internal/config"
	"github.com/atc0005/go-nagios"
	"golang.org/x/crypto/ssh"
)

// setPluginOutput is a helper function used to set plugin output and state
// values.
func setPluginOutput(
	stateLabel string,
	message string,
	extendedMessage string,
	err error,
	client *ssh.Client,
	cfg *config.Config,
	plugin *nagios.Plugin,
) {
	plugin.ServiceOutput = fmt.Sprintf(
		"%s: %s",
		strings.ToUpper(stateLabel),
		message,
	)

	if err != nil {
		plugin.AddError(err)
	}

	if cfg != nil {
		setLongServiceOutput(extendedMessage, client, cfg, plugin)
	}

	plugin.ExitStatusCode = nagios.StateLabelToExitCode(stateLabel)
}

func setLongServiceOutput(extendedMessage string, client *ssh.Client, cfg *config.Config, plugin *nagios.Plugin) {
	var report strings.Builder

	// fmt.Fprintf(
	// 	&report,
	// 	"%s---%s%s",
	// 	nagios.CheckOutputEOL,
	// 	nagios.CheckOutputEOL,
	// 	nagios.CheckOutputEOL,
	// )

	fmt.Fprintf(
		&report,
		"%sConfiguration settings: %s%s",
		nagios.CheckOutputEOL,
		nagios.CheckOutputEOL,
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* Server: %v%s",
		cfg.Server,
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* Port: %v%s",
		cfg.TCPPort,
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* Username: %v%s",
		cfg.Username,
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* NetworkType: %v%s",
		cfg.NetworkType,
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* Timeout: %v%s",
		cfg.Timeout(),
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* LoginFailureState: %v%s",
		strings.ToUpper(cfg.LoginFailureState),
		nagios.CheckOutputEOL,
	)

	fmt.Fprintf(
		&report,
		"* LoginSuccessState: %v%s",
		strings.ToUpper(cfg.LoginSuccessState),
		nagios.CheckOutputEOL,
	)

	if extendedMessage != "" {
		fmt.Fprintf(
			&report,
			"%s------%s%s",
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
		)

		fmt.Fprintf(
			&report,
			"%s%s",
			nagios.CheckOutputEOL,
			extendedMessage,
		)
	}

	fmt.Fprintf(
		&report,
		"%s",
		nagios.CheckOutputEOL,
	)

	if cfg.ShowVerbose {
		if client != nil {
			fmt.Fprintf(
				&report,
				"SSH Client version: %s%s",
				client.ClientVersion(),
				nagios.CheckOutputEOL,
			)
			fmt.Fprintf(
				&report,
				"SSH Server version: %s%s",
				client.ServerVersion(),
				nagios.CheckOutputEOL,
			)
		}
	}

	plugin.LongServiceOutput = report.String()
}

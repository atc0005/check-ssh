// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/atc0005/check-ssh/internal/config"
	"github.com/atc0005/go-nagios"
	"golang.org/x/crypto/ssh"
)

// closeClient is responsible for closing the given client connection or
// logging the error and adding it to the accumulated error collection for
// later processing.
func closeClient(client *ssh.Client, cfg *config.Config, plugin *nagios.Plugin) {
	logger := cfg.Log.With().Logger()

	if closeErr := client.Close(); closeErr != nil {
		if !errors.Is(closeErr, io.EOF) {
			logger.Error().Err(closeErr).Msgf(
				"failed to close client connection to %s",
				cfg.Server,
			)

			plugin.AddError(fmt.Errorf(
				"failed to close client connection to %s: %w",
				cfg.Server,
				closeErr,
			))
		}
	}

	logger.Debug().Msgf(
		"Successfully closed client connection to %s",
		cfg.Server,
	)
}

// closeSession is responsible for closing the given session or logging the
// error and adding it to the accumulated error collection for later
// processing.
func closeSession(session *ssh.Session, cfg *config.Config, plugin *nagios.Plugin) {
	logger := cfg.Log.With().Logger()

	if closeErr := session.Close(); closeErr != nil {
		if !errors.Is(closeErr, io.EOF) {
			logger.Error().Err(closeErr).Msgf(
				"failed to close connection to remote shell on %s",
				cfg.Server,
			)

			plugin.AddError(fmt.Errorf(
				"failed to close connection to remote shell on %s: %w",
				cfg.Server,
				closeErr,
			))
		}
	}

	logger.Debug().Msgf(
		"Successfully closed remote shell on %s",
		cfg.Server,
	)
}

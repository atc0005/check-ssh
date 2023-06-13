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
	"strings"

	"github.com/atc0005/check-ssh/internal/config"
	"github.com/atc0005/go-nagios"
	"golang.org/x/crypto/ssh"
)

var (
	// ErrSessionNotSet indicates that the required session value was not
	// provided.
	ErrSessionNotSet = errors.New("session value not set")

	// ErrConfigNotSet indicates that the required application configuration
	// value was not provided.
	ErrConfigNotSet = errors.New("config value not set")

	// ErrSSHCommandNotSet indicates that the required SSH command value was
	// not provided.
	ErrSSHCommandNotSet = errors.New("ssh command value not set")
)

func sshCommandSummary(cmdOutput []byte, cfg *config.Config) string {
	var sshCommandSummary string

	// Replace backslash characters with slashes in an effort to prevent
	// unintentional escaping.
	escapedOutput := strings.ReplaceAll(string(cmdOutput), `\\`, `\`)
	escapedOutput = strings.ReplaceAll(escapedOutput, `\`, `/`)

	switch {
	case strings.TrimSpace(string(cmdOutput)) != "":
		sshCommandSummary = fmt.Sprintf(
			"SSH command: %q%s%sSSH command output: %s%s%s%s",
			cfg.SSHCommand,
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
			escapedOutput,
			nagios.CheckOutputEOL,
		)
	default:
		sshCommandSummary = fmt.Sprintf(
			"SSH command: %q%s%sSSH command output: None%s",
			cfg.SSHCommand,
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
			nagios.CheckOutputEOL,
		)
	}

	return sshCommandSummary
}

func processSSHCommand(session *ssh.Session, cfg *config.Config) (string, error) {
	switch {
	case session == nil:
		return "", fmt.Errorf(
			"failed to process SSH command: session not set %w",
			ErrSessionNotSet,
		)
	case cfg == nil:
		return "", fmt.Errorf(
			"failed to process SSH command: %w",
			ErrConfigNotSet,
		)
	case cfg.SSHCommand == "":
		return "", fmt.Errorf(
			"failed to process SSH command: %w",
			ErrSSHCommandNotSet,
		)
	}

	cmdOutput, err := session.CombinedOutput(cfg.SSHCommand)
	if err != nil {
		return "", fmt.Errorf(
			"failed to process SSH command: %w",
			err,
		)
	}

	summary := sshCommandSummary(cmdOutput, cfg)

	return summary, err
}

// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package sshutils

import (
	"fmt"
	"net"
	"strings"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/ssh"
)

// newClient opens a connection to the specified SSH server using the
// specified authentication details, network type. A connected client is
// returned or an error if one occurs.
func newClient(sshPasswordAuthConfig SSHPasswordAuthConfig, sshConfig *ssh.ClientConfig, logger zerolog.Logger) (*ssh.Client, error) {
	logger = logger.With().
		Str("hostname", sshPasswordAuthConfig.Server).
		Str("net_type", sshPasswordAuthConfig.NetworkType).
		Logger()

	logger.Debug().Msg("resolving hostname")

	addrs, err := resolveIPAddresses(sshPasswordAuthConfig, logger)
	if err != nil {
		return nil, fmt.Errorf(
			"resolve hostname %s to %s IPs: %w",
			sshPasswordAuthConfig.Server,
			networkTypeToIPTypeStr(sshPasswordAuthConfig.NetworkType),
			err,
		)
	}

	conn, connectErr := openConnection(
		addrs,
		sshPasswordAuthConfig.Port,
		sshPasswordAuthConfig.NetworkType,
		logger,
	)

	if connectErr != nil {
		return nil, fmt.Errorf(
			"failed to create client connection to %s (port %d): %w",
			sshPasswordAuthConfig.Server,
			sshPasswordAuthConfig.Port,
			connectErr,
		)
	}

	sshConn, sshNewChan, sshRequestChan, clientConnErr := ssh.NewClientConn(
		conn,
		conn.RemoteAddr().String(),
		sshConfig,
	)

	if clientConnErr != nil {
		ip, port, splitErr := net.SplitHostPort(conn.RemoteAddr().String())
		if splitErr == nil {
			return nil, fmt.Errorf(
				"failed to establish authenticated SSH session to %s (ip: %s, port: %s): %s: %w",
				sshPasswordAuthConfig.Server,
				ip,
				port,
				clientConnErr,
				ErrAuthenticationFailed,
			)
		}

		return nil, fmt.Errorf(
			"failed to establish authenticated SSH session to %s using IP:Port %s: %s: %w",
			sshPasswordAuthConfig.Server,
			conn.RemoteAddr().String(),
			clientConnErr,
			ErrAuthenticationFailed,
		)
	}

	client := ssh.NewClient(sshConn, sshNewChan, sshRequestChan)

	return client, nil
}

// LoginWithPasswordAuth attempts to use the given credentials and SSH
// server host value to perform password authentication. If successful,
// pointers to a SSH client and session are returned, otherwise an error is
// returned.
func LoginWithPasswordAuth(
	sshPasswordAuthConfig SSHPasswordAuthConfig,
	logger zerolog.Logger,
) (*ssh.Client, *ssh.Session, error) {
	switch {
	case strings.TrimSpace(sshPasswordAuthConfig.Username) == "":
		return nil, nil, fmt.Errorf(
			"%w: required username not provided",
			ErrMissingValue,
		)
	case strings.TrimSpace(sshPasswordAuthConfig.Server) == "":
		return nil, nil, fmt.Errorf(
			"%w: required host not provided",
			ErrMissingValue,
		)
	case strings.TrimSpace(sshPasswordAuthConfig.Password) == "":
		return nil, nil, fmt.Errorf(
			"%w: required password not provided",
			ErrMissingValue,
		)
	}

	sshConfig := &ssh.ClientConfig{
		User:          sshPasswordAuthConfig.Username,
		Auth:          []ssh.AuthMethod{ssh.Password(sshPasswordAuthConfig.Password)},
		ClientVersion: sshPasswordAuthConfig.ClientVersion,
		Timeout:       sshPasswordAuthConfig.Timeout,
	}
	// TODO: Implement support for using a given host key to validate the SSH
	// server.
	//
	// sshConfig.HostKeyCallback = ssh.FixedHostKey()
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey() // nolint:gosec

	client, err := newClient(sshPasswordAuthConfig, sshConfig, logger)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"failed to login to %s: %w",
			sshPasswordAuthConfig.Server,
			err,
		)
	}

	session, err := client.NewSession()
	if err != nil {
		if closeErr := client.Close(); closeErr != nil {
			return nil, nil, fmt.Errorf(
				"error closing connection to server: %v: %w",
				err,
				ErrClientSessionFailed,
			)
		}

		return nil, nil, fmt.Errorf(
			"%v: %w",
			err,
			ErrClientSessionFailed,
		)
	}

	return client, session, nil
}

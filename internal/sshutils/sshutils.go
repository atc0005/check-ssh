// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package sshutils

import (
	"errors"
	"time"
)

// Known, named networks used for SSH connections. These names match the
// network names used by the `net` standard library package.
const (

	// NetTypeTCPAuto indicates that either of IPv4 or IPv6 will be used to
	// establish a connection depending on the specified IP Address.
	NetTypeTCPAuto string = "tcp"

	// NetTypeTCP4 indicates an IPv4-only network.
	NetTypeTCP4 string = "tcp4"

	// NetTypeTCP6 indicates an IPv6-only network.
	NetTypeTCP6 string = "tcp6"
)

var (
	// ErrMissingValue indicates that an expected value was missing.
	ErrMissingValue = errors.New("missing expected value")

	// ErrClientConnectionFailed indicates a general failure to establish a
	// client connection. This can occur due to invalid destination service
	// port, wrong username or password or other authentication related issue.
	ErrClientConnectionFailed = errors.New("client connection failed")

	// ErrClientSessionFailed indicates a failure to open a new client session.
	ErrClientSessionFailed = errors.New("client session failed")

	// ErrDNSLookupFailed indicates a failure to resolve a hostname to an IP
	// Address.
	ErrDNSLookupFailed = errors.New("failed to resolve hostname")

	// ErrIPAddressParsingFailed indicates a failure to parse a given value as
	// an IP Address.
	ErrIPAddressParsingFailed = errors.New("failed to parse IP Address")

	// ErrNoIPAddressesForChosenNetworkType indicates a failure to obtain any
	// IP Addresses of the specified network type (e.g., IPv4 vs IPv6) for a
	// given hostname.
	ErrNoIPAddressesForChosenNetworkType = errors.New("no resolved IP Addresses for chosen network type")

	// ErrNetworkConnectionFailed indicates a failure to establish a network
	// connection to the specified host.
	ErrNetworkConnectionFailed = errors.New("failed to establish network connection")

	// ErrAuthenticationFailed indicates a failure to authenticate to the
	// specified host.
	ErrAuthenticationFailed = errors.New("failed to authenticate")

	// ErrAuthenticationSucceeded indicates successful authentication to the
	// specified host. Depending on specified settings this may be unexpected.
	ErrAuthenticationSucceeded = errors.New("successfully authenticated")
)

// SSHPasswordAuthConfig is a SSH configuration used for password
// authentication.
type SSHPasswordAuthConfig struct {
	// Server is the SSH server FQDN or IP Address.
	Server string

	// Port is the port used by the SSH service.
	Port int

	// Username is the valid user for the given SSH server.
	Username string

	// Password is the valid password for the specified user.
	Password string

	// NetworkType indicates whether an attempt should be made to connect to
	// only IPv4, only IPv6 or SSH servers listening on either of IPv4 or IPv6
	// addresses ("auto").
	NetworkType string

	// ClientVersion contains the version identification string that will be
	// used for the connection. If empty, a reasonable default is used.
	ClientVersion string

	// Timeout is the maximum amount of time for the TCP connection to
	// establish. A Timeout of zero means no timeout.
	Timeout time.Duration
}

// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package sshutils

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
)

//
// FIXME: Move general network related functionality to a separate package.
//

// openConnection receives a list of IP Addresses and returns a net.Conn value
// for the first successful connection attempt. An error is returned instead
// if one occurs.
func openConnection(ctx context.Context, addrs []string, port int, netType string, logger zerolog.Logger) (net.Conn, error) {
	if len(addrs) < 1 {
		logger.Error().Msg("empty list of IP Addresses received")

		return nil, fmt.Errorf(
			"empty list of IP Addresses received: %w",
			ErrMissingValue,
		)
	}

	var (
		c          net.Conn
		connectErr error
	)

	for _, addr := range addrs {
		logger.Debug().
			Str("ip_address", addr).
			Msg("Connecting to server")

		s := net.JoinHostPort(addr, strconv.Itoa(port))

		// Unless sysadmin explicitly requested one of IPv4 or IPv6 network
		// types we fall back to default behavior.
		switch strings.ToLower(netType) {
		case NetTypeTCP4:
		case NetTypeTCP6:
		default:
			netType = NetTypeTCPAuto
		}

		// Attempt to connect to the given IP Address.
		dialer := &net.Dialer{}
		c, connectErr = dialer.DialContext(ctx, netType, s)

		// pass in explicitly set SSH config using provided server name, but
		// attempt to connect to specific IP Address returned from earlier
		// lookup. We'll attempt to loop over each available IP Address until
		// we are able to successfully connect to one of them.
		// c, connectErr = client.DialWithDialerTLS(&dialer, s, tlsConfig)

		if connectErr != nil {
			logger.Debug().
				Err(connectErr).
				Str("ip_address", addr).
				Msg("error connecting to server")

			continue
		}

		// If no connection errors were received, we can consider the
		// connection attempt a success and skip further attempts to connect
		// to any remaining IP Addresses for the specified server name.
		logger.Debug().
			Str("ip_address", addr).
			Msg("Connected to server")

		return c, nil
	}

	// If all connection attempts failed, report the last connection error.
	// Log all failed IP Addresses for review.
	if connectErr != nil {
		errMsg := fmt.Sprintf(
			"failed to connect to server using any of %d IP Addresses (%s)",
			len(addrs),
			strings.Join(addrs, ", "),
		)
		logger.Debug().
			Err(connectErr).
			Str("failed_ip_addresses", strings.Join(addrs, ", ")).
			Msg(errMsg)

		return nil, fmt.Errorf(
			"%s; last error: %v: %w",
			errMsg,
			connectErr,
			ErrNetworkConnectionFailed,
		)
	}

	return c, nil
}

// networkTypeToIPTypeStr resolves a network type or name (e.g., tcp4, tcp6)
// to a human readable IP Address type.
func networkTypeToIPTypeStr(netType string) string {
	switch strings.ToLower(netType) {
	case NetTypeTCP4:
		return "IPv4"
	case NetTypeTCP6:
		return "IPv6"
	default:
		return "IPv4 or IPv6"
	}
}

func lookupIPs(ctx context.Context, server string, logger zerolog.Logger) ([]string, error) {
	resolver := &net.Resolver{}
	lookupResults, lookupErr := resolver.LookupHost(ctx, server)

	if lookupErr != nil {
		logger.Error().
			Err(lookupErr).
			Str("server", server).
			Msg("error resolving hostname")

		return nil, fmt.Errorf(
			"error resolving hostname %s: %v: %w",
			server,
			lookupErr,
			ErrDNSLookupFailed,
		)
	}

	// FIXME: Is this length check really needed? Presumably if there were
	// zero results returned an error would have also been returned?
	switch {
	case len(lookupResults) < 1:
		errMsg := fmt.Sprintf(
			"failed to resolve hostname %s to IP Addresses",
			server,
		)

		logger.Error().
			Str("server", server).
			Msg(errMsg)

		return nil, fmt.Errorf(
			"%s: %w",
			errMsg,
			ErrDNSLookupFailed,
		)

	default:
		logger.Debug().
			Int("count", len(lookupResults)).
			Str("ips", strings.Join(lookupResults, ", ")).
			Str("server", server).
			Msg("successfully resolved IP Addresses for hostname")
	}

	return lookupResults, nil
}

func ipStringsToNetIPs(ipStrings []string, logger zerolog.Logger) ([]net.IP, error) {
	ips := make([]net.IP, 0, len(ipStrings))

	logger.Debug().Msg("converting DNS lookup results to net.IP values for net type validation")

	for i := range ipStrings {
		ip := net.ParseIP(ipStrings[i])
		if ip == nil {
			return nil, fmt.Errorf(
				"error parsing %s: %w",
				ipStrings[i],
				ErrIPAddressParsingFailed,
			)
		}

		ips = append(ips, ip)
	}

	// FIXME: Is this length check really needed? Presumably if there were
	// zero results from the parsing attempt an error would have ready been
	// returned?
	switch {
	case len(ips) < 1:
		errMsg := fmt.Sprintf(
			"failed to to convert DNS lookup results to net.IP values after receiving %d DNS lookup results ([%s])",
			len(ipStrings),
			strings.Join(ipStrings, ", "),
		)

		logger.Error().Msg(errMsg)

		return nil, fmt.Errorf(
			"%s: %w",
			errMsg,
			ErrIPAddressParsingFailed,
		)

	default:
		logger.Debug().Msg("successfully converted DNS lookup results to net.IP values")
	}

	return ips, nil
}

func netIPsToIPStrings(netIPs []net.IP) []string {
	ipStrs := make([]string, len(netIPs))
	for i := range netIPs {
		ipStrs[i] = netIPs[i].String()
	}

	return ipStrs
}

func filterNetIPsToIPv4(netIPs []net.IP, logger zerolog.Logger) []net.IP {
	filteredIPs := make([]net.IP, 0, len(netIPs))

	for i := range netIPs {
		if netIPs[i].To4() != nil {
			logger.Debug().
				Str("ipv4_address", netIPs[i].String()).
				Msg("matched IPv4 address")

			filteredIPs = append(filteredIPs, netIPs[i])
		}
	}

	return filteredIPs
}

func filterNetIPsToIPv6(netIPs []net.IP, logger zerolog.Logger) []net.IP {
	filteredIPs := make([]net.IP, 0, len(netIPs))

	for i := range netIPs {
		if netIPs[i].To4() == nil {
			// If earlier attempts to parse the IP Address succeeded (by way
			// of it being a net.IP value), but this is not considered an IPv4
			// address, we will consider it a valid IPv6 address.
			logger.Debug().
				Str("ipv6_address", netIPs[i].String()).
				Msg("matched IPv6 address")

			filteredIPs = append(filteredIPs, netIPs[i])
		}
	}

	return filteredIPs
}

func filterNetIPsToNetworkType(netIPs []net.IP, netType string, logger zerolog.Logger) ([]net.IP, error) {
	var filteredIPs []net.IP

	// Flag validation ensures that we see valid named networks as supported
	// by the `net` stdlib package, along with the "auto" keyword. Here we pay
	// attention to only the valid named networks. Since we're working with
	// user specified keywords, we compare case-insensitively.
	switch strings.ToLower(netType) {
	case NetTypeTCP4:
		logger.Debug().Msg("user opted for IPv4-only connectivity, gathering only IPv4 addresses")

		filteredIPs = filterNetIPsToIPv4(netIPs, logger)

	case NetTypeTCP6:
		logger.Debug().Msg("user opted for IPv6-only connectivity, gathering only IPv6 addresses")

		filteredIPs = filterNetIPsToIPv6(netIPs, logger)

	// either of IPv4 or IPv6 is acceptable
	default:
		logger.Debug().Msg("auto behavior enabled, gathering all addresses")

		filteredIPs = netIPs
	}

	// No IPs remain after filtering against IPv4-only or IPv6-only
	// requirement.
	switch {
	case len(filteredIPs) < 1:
		errMsg := fmt.Sprintf(
			"failed to gather IP Addresses when filtering %d IPs by specified network type %s ([%s])",
			len(netIPs),
			netType,
			strings.Join(netIPsToIPStrings(netIPs), ", "),
		)

		logger.Error().Msg(errMsg)

		return nil, fmt.Errorf(
			"%s: %w",
			errMsg,
			ErrNoIPAddressesForChosenNetworkType,
		)

	default:
		logger.Debug().
			Int("num_input_ips", len(netIPs)).
			Int("num_remaining_ips", len(filteredIPs)).
			Str("network_type", netType).
			Str("ips", strings.Join(netIPsToIPStrings(filteredIPs), ", ")).
			Msg("successfully gathered IP Addresses for specified network type")
	}

	return filteredIPs, nil
}

func resolveIPAddresses(ctx context.Context, sshPasswordAuthConfig SSHPasswordAuthConfig, logger zerolog.Logger) ([]string, error) {
	lookupResults, lookupErr := lookupIPs(ctx, sshPasswordAuthConfig.Server, logger)
	if lookupErr != nil {
		return nil, lookupErr
	}

	netIPs, ipConvertErr := ipStringsToNetIPs(lookupResults, logger)
	if ipConvertErr != nil {
		return nil, ipConvertErr
	}

	filteredNetIPs, filterIPsErr := filterNetIPsToNetworkType(netIPs, sshPasswordAuthConfig.NetworkType, logger)
	if filterIPsErr != nil {
		return nil, filterIPsErr
	}

	ipStrings := netIPsToIPStrings(filteredNetIPs)

	return ipStrings, nil
}

// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Updated via Makefile builds. Setting placeholder value here so that
// something resembling a version string will be provided for non-Makefile
// builds.
var version = "devbuild"

var (
	// ErrVersionRequested indicates that the user requested application version
	// information.
	ErrVersionRequested = errors.New("version information requested")

	// ErrHelpRequested indicates that the user requested application
	// help/usage information.
	ErrHelpRequested = errors.New("help/usage information requested")

	// ErrUnsupportedOption indicates that an unsupported option was specified.
	ErrUnsupportedOption = errors.New("unsupported option")

	// ErrConfigNotInitialized indicates that the configuration is not in a
	// usable state and application execution can not successfully proceed.
	ErrConfigNotInitialized = errors.New("configuration not initialized")
)

// AppType represents the type of application that is being
// configured/initialized. Not all application types will use the same
// features and as a result will not accept the same flags. Unless noted
// otherwise, each of the application types are incompatible with each other,
// though some flags are common to all types.
type AppType struct {

	// Plugin represents an application used as a Nagios plugin.
	Plugin bool

	// Inspector represents an application used for one-off or isolated
	// checks. Unlike a Nagios plugin which is focused on specific attributes
	// resulting in a severity-based outcome, an Inspector application is
	// intended for examining a small set of targets for
	// informational/troubleshooting purposes.
	Inspector bool
}

// InspectorSettings is the collection of settings specific to the Inspector
// application type.
// type InspectorSettings struct {
// 	// ShowAll indicates whether the user opted to display information for ALL
// 	// processes. This can produce a lot of output
// 	ShowAll bool
// }

// Config represents the application configuration as specified via
// command-line flags.
type Config struct {

	// flagSet provides a useful hook to allow evaluating defined flags
	// against a list of expected flags. This field is exported so that the
	// flagset is accessible to tests from within this package and from
	// outside of the config package.
	flagSet *flag.FlagSet

	// LoggingLevel is the supported logging level for this application.
	LoggingLevel string

	// NetworkType indicates whether an attempt should be made to connect to
	// only IPv4, only IPv6 or SSH servers listening on either of IPv4 or IPv6
	// addresses ("auto").
	NetworkType string

	// Server is the SSH server FQDN or IP Address.
	Server string

	// Username is the valid user for the given SSH server.
	Username string

	// Password is the valid password for the specified user.
	Password string

	// SSHCommand is the valid shell command to execute via SSH after a
	// successful connection and session is established.
	SSHCommand string

	// LoginSuccessState is the plugin state to be used when a SSH login
	// attempt is successful.
	LoginSuccessState string

	// LoginFailureState is the plugin state to be ued when a SSH login
	// attempt is unsuccessful.
	LoginFailureState string

	// TCPPort is the port used by the SSH service.
	TCPPort int

	// timeout is the number of seconds allowed before the connection attempt
	// to the SSH service is abandoned and an error returned.
	timeout int

	// Log is an embedded zerolog Logger initialized via config.New().
	Log zerolog.Logger

	// InspectorSettings is the collection of settings specific to the
	// Inspector application type.
	// InspectorSettings InspectorSettings

	// EmitBranding controls whether "generated by" text is included at the
	// bottom of application output. This output is included in the Nagios
	// dashboard and notifications. This output may not mix well with branding
	// output from other tools such as atc0005/send2teams which also insert
	// their own branding output.
	EmitBranding bool

	// EmitSSHCommandOutput controls whether SSH command output is emitted by
	// the plugin *if* a SSH command was specified and executed.
	EmitSSHCommandOutput bool

	// ShowVersion is a flag indicating whether the user opted to display only
	// the version string and then immediately exit the application.
	ShowVersion bool

	// ShowVerbose is a flag indicating whether the user opted to display
	// verbose details in the final plugin output.
	ShowVerbose bool

	// ShowHelp indicates whether the user opted to display usage information
	// and exit the application.
	ShowHelp bool
}

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}

// SSHClientVersion emits application name & version in a format expected by
// remote SSH servers.
func SSHClientVersion() string {
	// Replace forbidden characters with underscore to comply with RFC4253
	// requirements for 'softwareversion' string.
	//
	// Per https://www.rfc-editor.org/rfc/rfc4253#section-4.2:
	//
	//	SSH-protoversion-softwareversion SP comments CR LF
	//
	// Both the 'protoversion' and 'softwareversion' strings MUST consist of
	// printable US-ASCII characters, with the exception of whitespace
	// characters and the minus sign (-).  The 'softwareversion' string is
	// primarily used to trigger compatibility extensions and to indicate
	// the capabilities of an implementation.  The 'comments' string SHOULD
	// contain additional information that might be useful in solving user
	// problems.  As such, an example of a valid identification string is
	//
	// SSH-2.0-billsSSH_3.6.3q3<CR><LF>
	//
	// This identification string does not contain the optional 'comments'
	// string and is thus terminated by a CR and LF immediately after the
	// 'softwareversion' string
	//
	r := strings.NewReplacer(
		" ", "_",
		"-", "_",
	)

	softwareversion := fmt.Sprintf(
		"%s_%s",
		r.Replace(myAppName),
		r.Replace(version),
	)

	comment := fmt.Sprintf(
		"(%s)",
		myAppURL,
	)

	return fmt.Sprintf(
		"SSH-2.0-%s %s",
		softwareversion,
		comment,
	)
}

// Branding accepts a message and returns a function that concatenates that
// message with version information. This function is intended to be called as
// a final step before application exit after any other output has already
// been emitted.
func Branding(msg string) func() string {
	return func() string {
		return strings.Join([]string{msg, Version()}, "")
	}
}

// Usage is a custom override for the default Help text provided by the flag
// package. Here we prepend some additional metadata to the existing output.
func Usage(flagSet *flag.FlagSet, w io.Writer) func() {
	// Make one attempt to override output so that calling Config.Help() later
	// will have a chance to also override the output destination.
	flag.CommandLine.SetOutput(w)

	switch {
	// Uninitialized flagset, provide stub usage information.
	case flagSet == nil:
		return func() {
			_, _ = fmt.Fprintln(w, "Failed to initialize configuration; nil FlagSet")
		}

	// Non-nil flagSet, proceed
	default:
		// Make one attempt to override output so that calling Config.Help()
		// later will have a chance to also override the output destination.
		flagSet.SetOutput(w)

		return func() {
			_, _ = fmt.Fprintln(flag.CommandLine.Output(), "\n"+Version()+"\n")
			_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
			flagSet.PrintDefaults()
		}
	}
}

// Help emits application usage information to the previously configured
// destination for usage and error messages.
func (c *Config) Help() string {
	var helpTxt strings.Builder

	// Override previously specified output destination, redirect to Builder.
	flag.CommandLine.SetOutput(&helpTxt)

	switch {
	// Handle nil configuration initialization.
	case c == nil || c.flagSet == nil:
		// Fallback message noting the issue.
		_, _ = fmt.Fprintln(&helpTxt, ErrConfigNotInitialized)

	default:
		// Emit expected help output to builder.
		c.flagSet.SetOutput(&helpTxt)
		c.flagSet.Usage()
	}

	return helpTxt.String()
}

// New is a factory function that produces a new Config object based on user
// provided flag and config file values. It is responsible for validating
// user-provided values and initializing the logging settings used by this
// application.
func New(appType AppType) (*Config, error) {
	var config Config

	// NOTE: Need to make sure we allow execution to continue on encountered
	// errors. This is so that we can check for those errors as return values
	// both within the main apps and tests for this package.
	config.flagSet = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	if err := config.handleFlagsConfig(appType); err != nil {
		return nil, fmt.Errorf(
			"failed to set flags configuration: %w",
			err,
		)
	}

	switch {
	// The configuration was successfully initialized, so we're good with
	// returning it for use by the caller.
	case config.ShowVersion:
		return &config, ErrVersionRequested

	// The configuration was successfully initialized, so we're good with
	// returning it for use by the caller.
	case config.ShowHelp:
		return &config, ErrHelpRequested
	}

	if err := config.validate(appType); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	// initialize logging just as soon as validation is complete
	if err := config.setupLogging(appType); err != nil {
		return nil, fmt.Errorf(
			"failed to set logging configuration: %w",
			err,
		)
	}

	return &config, nil
}

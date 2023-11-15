// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

const myAppName string = "check-ssh"
const myAppURL string = "https://github.com/atc0005/check-ssh"

// ExitCodeCatchall indicates a general or miscellaneous error has occurred.
// This exit code is not directly used by monitoring plugins in this project.
// See https://tldp.org/LDP/abs/html/exitcodes.html for additional details.
const ExitCodeCatchall int = 1

// shorthandFlagSuffix is appended to short flag help text to emphasize that
// the flag is a shorthand version of a longer flag.
const shorthandFlagSuffix = " (shorthand)"

// Shared flags help text.
const (
	versionFlagHelp          string = "Whether to display application version and then immediately exit application."
	logLevelFlagHelp         string = "Sets log level."
	brandingFlagHelp         string = "Toggles emission of branding details with plugin status details. This output is disabled by default."
	helpFlagHelp             string = "Emit this help text"
	serverFlagHelp           string = "The SSH server FQDN or IP Address."
	usernameFlagHelp         string = "The valid user for the given SSH server."
	passwordFlagHelp         string = "The valid password for the specified user." //nolint:gosec
	tcpPortFlagHelp          string = "The port used by the SSH service."
	timeoutFlagHelp          string = "Timeout value in seconds allowed before a connection attempt to a SSH service is abandoned and an error returned."
	networkTypeFlagHelp      string = "Limits network connections to one of tcp4 (IPv4-only), tcp6 (IPv6-only) or auto (either)."
	sshCommandFlagHelp       string = "The valid shell command to execute via SSH."
	sshCommandOutputFlagHelp string = "Toggles emission of SSH command output (if a SSH command is executed). This output is disabled by default."
)

// Plugin flags help text.
const (
	// loginSuccessIsWarningFlagHelp  string = "Indicates that a successful login via SSH is considered a WARNING state."
	// loginSuccessIsCriticalFlagHelp string = "Indicates that a successful login via SSH is considered a CRITICAL state."
	// loginSuccessIsOKFlagHelp       string = "Indicates that a successful login via SSH is considered an OK state."
	loginSuccessStateFlagHelp string = "The plugin state to use when a SSH login is successful."
	loginFailureStateFlagHelp string = "The plugin state to use when a SSH login fails."
	verboseFlagHelp           string = "Whether to display verbose details in the final plugin output."
)

// Flag names for consistent references. Exported so that they're available
// from tests.
const (
	HelpFlagLong              string = "help"
	HelpFlagShort             string = "h"
	VersionFlagLong           string = "version"
	VerboseFlagLong           string = "verbose"
	BrandingFlag              string = "branding"
	TimeoutFlagLong           string = "timeout"
	TimeoutFlagShort          string = "t"
	LogLevelFlagLong          string = "log-level"
	LogLevelFlagShort         string = "ll"
	LoginSuccessStateFlagLong string = "login-success-state"
	LoginFailureStateFlagLong string = "login-failure-state"
	ServerFlagLong            string = "server"
	UsernameFlagLong          string = "username"
	PasswordFlagLong          string = "password"
	PortFlagLong              string = "port"
	NetTypeFlagLong           string = "net-type"
	SSHCommandFlagLong        string = "command"
	SSHCommandOutputFlagLong  string = "command-output"
)

// Default flag settings if not overridden by user input
const (
	defaultHelp                  bool   = false
	defaultLogLevel              string = "info"
	defaultVerbose               bool   = false
	defaultEmitBranding          bool   = false
	defaultEmitSSHCommandOutput  bool   = false
	defaultDisplayVersionAndExit bool   = false
	defaultServer                string = ""
	defaultUsername              string = ""
	defaultPassword              string = ""
	defaultTCPPort               int    = 22
	defaultNetworkType           string = netTypeTCPAuto
	defaultSSHCommand            string = ""
	defaultTimeout               int    = 10
	defaultLoginSuccessState     string = "ok"
	defaultLoginFailureState     string = "critical"

	// defaultLoginSuccessState     string = "ok"       // FIXME: Should this be a nagios package constant instead?
	// defaultLoginFailureState     string = "critical" // FIXME: Should this be a nagios package constant instead?

	// FIXME: Yank this once I'm sure I won't use them.
	//
	// defaultLoginSuccessIsWarning  bool   = false
	// defaultLoginSuccessIsCritical bool   = false
	// defaultLoginSuccessIsOK       bool   = true
	// defaultIgnoreHostKey         bool   = true
)

const (
	// netTypeTCPAuto is a custom keyword indicating that either of IPv4 or
	// IPv6 is an acceptable network type.
	netTypeTCPAuto string = "auto"

	// netTypeTCP4 indicates that IPv4 network connections are required.
	netTypeTCP4 string = "tcp4"

	// netTypeTCP6 indicates that IPv6 network connections are required
	netTypeTCP6 string = "tcp6"
)

const (
	appTypePlugin    string = "plugin"
	appTypeInspector string = "Inspector" // TODO: This could be used to list remote host keys, test connectivity
)

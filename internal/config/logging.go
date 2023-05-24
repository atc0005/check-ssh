// Copyright 2023 Adam Chalkley
//
// https://github.com/atc0005/check-ssh
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

const (
	// LogLevelDisabled maps to zerolog.Disabled logging level
	LogLevelDisabled string = "disabled"

	// LogLevelPanic maps to zerolog.PanicLevel logging level
	LogLevelPanic string = "panic"

	// LogLevelFatal maps to zerolog.FatalLevel logging level
	LogLevelFatal string = "fatal"

	// LogLevelError maps to zerolog.ErrorLevel logging level
	LogLevelError string = "error"

	// LogLevelWarn maps to zerolog.WarnLevel logging level
	LogLevelWarn string = "warn"

	// LogLevelInfo maps to zerolog.InfoLevel logging level
	LogLevelInfo string = "info"

	// LogLevelDebug maps to zerolog.DebugLevel logging level
	LogLevelDebug string = "debug"

	// LogLevelTrace maps to zerolog.TraceLevel logging level
	LogLevelTrace string = "trace"
)

// setLoggingLevel applies the requested logging level to filter out messages
// with a lower level than the one configured.
func setLoggingLevel(logLevel string) error {
	switch logLevel {
	case LogLevelDisabled:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	case LogLevelPanic:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case LogLevelFatal:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case LogLevelError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case LogLevelWarn:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case LogLevelTrace:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		return fmt.Errorf(
			"%w: invalid logging level provided: %v",
			ErrUnsupportedOption,
			logLevel,
		)
	}

	// signal that a case was triggered as expected
	return nil
}

// setupLogging is responsible for configuring logging settings for this
// application
func (c *Config) setupLogging(appType AppType) error {
	// We set some common fields here so that we don't have to repeat them
	// explicitly later. This approach is intended to help standardize the log
	// messages to make them easier to search through later when
	// troubleshooting. We can extend the logged fields as needed by each CLI
	// application or Nagios plugin to cover unique details.
	switch {
	case appType.Inspector:
		// CLI app logging uses ConsoleWriter to generate human-friendly,
		// colorized output to stdout.
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
		c.Log = zerolog.New(consoleWriter).With().Timestamp().Caller().
			Str("version", Version()).
			Str("logging_level", c.LoggingLevel).
			Str("app_type", appTypeInspector).
			Logger()

	case appType.Plugin:
		// Plugin logging uses ConsoleWriter to generate human-friendly, (but
		// for this app type) uncolorized output to stderr. Log output is sent
		// to stderr to prevent mixing in with stdout output intended for the
		// Nagios console.
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true}
		c.Log = zerolog.New(consoleWriter).With().Timestamp().Caller().
			Str("version", Version()).
			Str("logging_level", c.LoggingLevel).
			Str("app_type", appTypePlugin).
			Logger()
	}

	return setLoggingLevel(c.LoggingLevel)
}

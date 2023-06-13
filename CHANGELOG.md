# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/check-ssh/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.1.0] - 2023-06-13

### Overview

- Initial release
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)
  - Windows (x86, x64)

### Added

Initial release!

This release provides early release versions of a plugin used to monitor SSH
access:

| Tool Name         | Overall Status | Description                                                     |
| ----------------- | -------------- | --------------------------------------------------------------- |
| `check_ssh_login` | Alpha          | Nagios plugin used to monitor for unexpected SSH login results. |

See the project README for additional details.

[Unreleased]: https://github.com/atc0005/check-ssh/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.0

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

## [v0.3.2] - 2023-07-13

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.10` to `1.19.11`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.1` to `go-ci-oldstable-build-v0.11.3`
  - `golang.org/x/crypto`
    - `v0.10.0` to `v0.11.0`
  - `golang.org/x/sys`
    - `v0.9.0` to `v0.10.0`
- (GH-44) List supported network types in flag help text
- (GH-50) Update RPM postinstall scripts to use restorecon

### Fixed

- (GH-47) README missing performance data metrics table
- (GH-48) Correct logging format listed in README

## [v0.3.1] - 2023-06-30

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)
  - Windows (x86, x64)

### Changed

- Dependencies
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.0` to `go-ci-oldstable-build-v0.11.1`
  - `atc0005/go-nagios`
    - `v0.15.0` to `v0.16.0`

### Fixed

- (GH-40) Removed unused config constant
- (GH-41) Fix logger field label for username value

## [v0.3.0] - 2023-06-21

### Overview

- Add new flag
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)
  - Windows (x86, x64)

### Added

- (GH-34) Allow toggling SSH command output on/off

## [v0.2.0] - 2023-06-20

### Overview

- Add new flag
- Dependency updates
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)
  - Windows (x86, x64)

### Added

- (GH-30) Add support for setting custom timeout

### Changed

- Dependencies
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.6` to `go-ci-oldstable-build-v0.11.0`

## [v0.1.1] - 2023-06-14

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)
  - Windows (x86, x64)

### Changed

- Dependencies
  - `golang.org/x/crypto`
    - `v0.9.0` to `v0.10.0`
- (GH-25) Add CodeQL workflow configuration

### Fixed

- (GH-26) Enable building of Windows assets

## [v0.1.0] - 2023-06-13

### Overview

- Initial release
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)

### Added

Initial release!

This release provides early release versions of a plugin used to monitor SSH
access:

| Tool Name         | Overall Status | Description                                                     |
| ----------------- | -------------- | --------------------------------------------------------------- |
| `check_ssh_login` | Alpha          | Nagios plugin used to monitor for unexpected SSH login results. |

See the project README for additional details.

[Unreleased]: https://github.com/atc0005/check-ssh/compare/v0.3.2...HEAD
[v0.3.2]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.0
[v0.2.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.2.0
[v0.1.1]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.0

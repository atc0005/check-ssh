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

## [v0.3.4] - 2023-10-06

### Changed

#### Dependency Updates

- (GH-99) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-88) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.4 to go-ci-oldstable-build-v0.13.5 in /dependabot/docker/builds
- (GH-90) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-92) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-100) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-107) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-111) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.10 in /dependabot/docker/builds
- (GH-96) ghaw: bump actions/checkout from 3 to 4
- (GH-98) go.mod: bump golang.org/x/crypto from 0.12.0 to 0.13.0
- (GH-93) go.mod: bump golang.org/x/sys from 0.11.0 to 0.12.0

## [v0.3.3] - 2023-08-17

### Added

- (GH-56) Add initial automated release notes config
- (GH-58) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.3` to `go-ci-oldstable-build-v0.13.4`
  - `rs/zerolog`
    - `v1.29.1` to `v1.30.0`
  - `golang.org/x/crypto`
    - `v0.11.0` to `v0.12.0`
  - `golang.org/x/sys`
    - `v0.10.0` to `v0.11.0`
- (GH-60) Update Dependabot config to monitor both branches
- (GH-82) Update project to Go 1.20 series

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

[Unreleased]: https://github.com/atc0005/check-ssh/compare/v0.3.4...HEAD
[v0.3.4]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.4
[v0.3.3]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.0
[v0.2.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.2.0
[v0.1.1]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.0

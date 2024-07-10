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

## [v0.3.14] - 2024-07-10

### Changed

#### Dependency Updates

- (GH-289) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.7 to go-ci-oldstable-build-v0.20.8 in /dependabot/docker/builds
- (GH-293) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.8 to go-ci-oldstable-build-v0.21.2 in /dependabot/docker/builds
- (GH-295) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.2 to go-ci-oldstable-build-v0.21.3 in /dependabot/docker/builds
- (GH-298) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.3 to go-ci-oldstable-build-v0.21.4 in /dependabot/docker/builds
- (GH-303) Go Dependency: Bump golang.org/x/crypto from 0.24.0 to 0.25.0
- (GH-301) Go Dependency: Bump golang.org/x/sys from 0.21.0 to 0.22.0
- (GH-296) Go Runtime: Bump golang from 1.21.11 to 1.21.12 in /dependabot/docker/go

## [v0.3.13] - 2024-06-07

### Changed

#### Dependency Updates

- (GH-269) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.4 to go-ci-oldstable-build-v0.20.5 in /dependabot/docker/builds
- (GH-272) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.5 to go-ci-oldstable-build-v0.20.6 in /dependabot/docker/builds
- (GH-283) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.6 to go-ci-oldstable-build-v0.20.7 in /dependabot/docker/builds
- (GH-271) Go Dependency: Bump github.com/rs/zerolog from 1.32.0 to 1.33.0
- (GH-281) Go Dependency: Bump golang.org/x/crypto from 0.23.0 to 0.24.0
- (GH-280) Go Dependency: Bump golang.org/x/sys from 0.20.0 to 0.21.0
- (GH-282) Go Runtime: Bump golang from 1.21.10 to 1.21.11 in /dependabot/docker/go

### Fixed

- (GH-274) Remove inactive maligned linter
- (GH-275) Fix errcheck linting errors

## [v0.3.12] - 2024-05-11

### Changed

#### Dependency Updates

- (GH-252) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.1 to go-ci-oldstable-build-v0.20.2 in /dependabot/docker/builds
- (GH-262) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.2 to go-ci-oldstable-build-v0.20.3 in /dependabot/docker/builds
- (GH-263) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.3 to go-ci-oldstable-build-v0.20.4 in /dependabot/docker/builds
- (GH-257) Go Dependency: Bump golang.org/x/crypto from 0.22.0 to 0.23.0
- (GH-253) Go Dependency: Bump golang.org/x/sys from 0.19.0 to 0.20.0
- (GH-259) Go Runtime: Bump golang from 1.21.9 to 1.21.10 in /dependabot/docker/go

## [v0.3.11] - 2024-04-11

### Changed

#### Dependency Updates

- (GH-233) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.4 to go-ci-oldstable-build-v0.16.0 in /dependabot/docker/builds
- (GH-235) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.0 to go-ci-oldstable-build-v0.16.1 in /dependabot/docker/builds
- (GH-237) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.1 to go-ci-oldstable-build-v0.19.0 in /dependabot/docker/builds
- (GH-238) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.19.0 to go-ci-oldstable-build-v0.20.0 in /dependabot/docker/builds
- (GH-244) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.0 to go-ci-oldstable-build-v0.20.1 in /dependabot/docker/builds
- (GH-242) Go Dependency: Bump golang.org/x/crypto from 0.21.0 to 0.22.0
- (GH-243) Go Dependency: Bump golang.org/x/sys from 0.18.0 to 0.19.0
- (GH-240) Go Runtime: Bump golang from 1.21.8 to 1.21.9 in /dependabot/docker/go

## [v0.3.10] - 2024-03-08

### Changed

#### Dependency Updates

- (GH-228) Add todo/release label to "Go Runtime" PRs
- (GH-217) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.2 to go-ci-oldstable-build-v0.15.3 in /dependabot/docker/builds
- (GH-226) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.3 to go-ci-oldstable-build-v0.15.4 in /dependabot/docker/builds
- (GH-213) canary: bump golang from 1.21.6 to 1.21.7 in /dependabot/docker/go
- (GH-209) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.15.0 to go-ci-oldstable-build-v0.15.2 in /dependabot/docker/builds
- (GH-219) Go Dependency: Bump golang.org/x/crypto from 0.19.0 to 0.20.0
- (GH-221) Go Dependency: Bump golang.org/x/crypto from 0.20.0 to 0.21.0
- (GH-220) Go Dependency: Bump golang.org/x/sys from 0.17.0 to 0.18.0
- (GH-225) Go Runtime: Bump golang from 1.21.7 to 1.21.8 in /dependabot/docker/go
- (GH-215) Update Dependabot PR prefixes (redux)
- (GH-214) Update Dependabot PR prefixes
- (GH-210) Update project to Go 1.21 series

## [v0.3.9] - 2024-02-20

### Changed

#### Dependency Updates

- (GH-201) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-203) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.6 to go-ci-oldstable-build-v0.14.9 in /dependabot/docker/builds
- (GH-207) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.9 to go-ci-oldstable-build-v0.15.0 in /dependabot/docker/builds
- (GH-192) go.mod: bump github.com/rs/zerolog from 1.31.0 to 1.32.0
- (GH-200) go.mod: bump golang.org/x/crypto from 0.18.0 to 0.19.0
- (GH-199) go.mod: bump golang.org/x/sys from 0.16.0 to 0.17.0

## [v0.3.8] - 2024-02-02

### Changed

#### Dependency Updates

- (GH-176) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.3 to go-ci-oldstable-build-v0.14.4 in /dependabot/docker/builds
- (GH-182) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.4 to go-ci-oldstable-build-v0.14.5 in /dependabot/docker/builds
- (GH-186) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.5 to go-ci-oldstable-build-v0.14.6 in /dependabot/docker/builds
- (GH-181) go.mod: bump github.com/atc0005/go-nagios from 0.16.0 to 0.16.1

### Fixed

- (GH-188) Add missing context usage

## [v0.3.7] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-169) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-172) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.2 to go-ci-oldstable-build-v0.14.3 in /dependabot/docker/builds
- (GH-161) ghaw: bump github/codeql-action from 2 to 3
- (GH-167) go.mod: bump golang.org/x/crypto from 0.16.0 to 0.18.0
- (GH-165) go.mod: bump golang.org/x/sys from 0.15.0 to 0.16.0

## [v0.3.6] - 2023-12-09

### Changed

#### Dependency Updates

- (GH-154) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-155) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.1 to go-ci-oldstable-build-v0.14.2 in /dependabot/docker/builds
- (GH-149) go.mod: bump golang.org/x/crypto from 0.15.0 to 0.16.0
- (GH-150) go.mod: bump golang.org/x/sys from 0.14.0 to 0.15.0

## [v0.3.5] - 2023-11-15

### Changed

#### Dependency Updates

- (GH-138) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-128) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-129) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.10 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-143) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds
- (GH-132) go.mod: bump github.com/mattn/go-isatty from 0.0.19 to 0.0.20
- (GH-115) go.mod: bump github.com/rs/zerolog from 1.30.0 to 1.31.0
- (GH-120) go.mod: bump golang.org/x/crypto from 0.13.0 to 0.14.0
- (GH-140) go.mod: bump golang.org/x/crypto from 0.14.0 to 0.15.0
- (GH-119) go.mod: bump golang.org/x/sys from 0.12.0 to 0.13.0
- (GH-136) go.mod: bump golang.org/x/sys from 0.13.0 to 0.14.0

### Fixed

- (GH-145) Fix goconst linting errors

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

[Unreleased]: https://github.com/atc0005/check-ssh/compare/v0.3.14...HEAD
[v0.3.14]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.14
[v0.3.13]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.13
[v0.3.12]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.12
[v0.3.11]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.11
[v0.3.10]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.10
[v0.3.9]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.9
[v0.3.8]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.8
[v0.3.7]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.7
[v0.3.6]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.6
[v0.3.5]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.5
[v0.3.4]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.4
[v0.3.3]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.3.0
[v0.2.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.2.0
[v0.1.1]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-ssh/releases/tag/v0.1.0

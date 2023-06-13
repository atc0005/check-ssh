<!-- omit in toc -->
# check-ssh

Go-based tooling used to monitor SSH access.

[![Latest Release](https://img.shields.io/github/release/atc0005/check-ssh.svg?style=flat-square)](https://github.com/atc0005/check-ssh/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/check-ssh.svg)](https://pkg.go.dev/github.com/atc0005/check-ssh)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/check-ssh)](https://github.com/atc0005/check-ssh)
[![Lint and Build](https://github.com/atc0005/check-ssh/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/check-ssh/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/check-ssh/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/check-ssh/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
  - [`check_ssh_login` plugin](#check_ssh_login-plugin)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
  - [Deployment](#deployment)
- [Configuration](#configuration)
  - [Command-line arguments](#command-line-arguments)
    - [`check_ssh_login`](#check_ssh_login)
- [Examples](#examples)
  - [Successful authentication](#successful-authentication)
    - [`OK` result](#ok-result)
    - [non-OK result](#non-ok-result)
  - [Failed authentication](#failed-authentication)
    - [`OK` result](#ok-result-1)
    - [non-OK result](#non-ok-result-1)
  - [Execute shell command](#execute-shell-command)
    - [No output](#no-output)
    - [SSH command output generated](#ssh-command-output-generated)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo is intended to provide various tools used to monitor SSH access.

| Tool Name         | Overall Status | Description                                                     |
| ----------------- | -------------- | --------------------------------------------------------------- |
| `check_ssh_login` | Alpha          | Nagios plugin used to monitor for unexpected SSH login results. |

## Features

### `check_ssh_login` plugin

Nagios plugin (`check_ssh_login`) used to monitor for for unexpected SSH login
results.

- Optional override of authentication behavior
  - successful authentication defaults to `OK` state
  - failed authentication defaults to `CRITICAL` state
  - both states may be overridden to reflect specific scenarios (see
    [examples](#examples))
    - e.g., an assertion that an unauthorized account should not be permitted
    - e.g., an assertion that an authorized account should be permitted

- Optional SSH command execution
  - output from command execution is collected/emitted
  - **NOTE**: some output may be problematic as [Nagios disallows some
    characters in generated output][nagios-illegal-macro-chars]

- Optional limiting of network type
  - to either of IPv4 and IPv6
  - default behavior is `auto`

- Optional, leveled logging using `rs/zerolog` package
  - JSON-format output (to `stderr`)
  - choice of `disabled`, `panic`, `fatal`, `error`, `warn`, `info` (the
    default), `debug` or `trace`

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 11
- Ubuntu Linux 22.04
- Red Hat Enterprise Linux 8

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/check-ssh`
   1. `cd check-ssh`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - manually, explicitly specifying target OS and architecture
     - `GOOS=linux GOARCH=amd64 go build -mod=vendor ./cmd/check_ssh_login/`
       - most likely this is what you want (if building manually)
       - substitute `amd64` with the appropriate architecture if using
         different hardware (e.g., `arm64` or `386`)
   - using Makefile `linux` recipe
     - `make linux`
       - generates x86 and x64 binaries
   - using Makefile `release-build` recipe
     - `make release-build`
       - generates the same release assets as provided by this project's
         releases
1. Locate generated binaries
   - if using `Makefile`
     - look in `/tmp/check-ssh/release_assets/check_ssh_login/`
   - if using `go build`
     - look in `/tmp/check-ssh/`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**: Depending on which `Makefile` recipe you use the generated binary
may be compressed and have an `xz` extension. If so, you should decompress the
binary first before deploying it (e.g., `xz -d check_ssh_login-linux-amd64.xz`).

### Using release binaries

1. Download the [latest release][repo-url] binaries
1. Decompress binaries
   - e.g., `xz -d check_ssh_login-linux-amd64.xz`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

### Deployment

1. Place `check_ssh_login` in a location where it can be executed by the
   monitoring agent
   - Usually the same place as other Nagios plugins
   - For example, on a default Red Hat Enterprise Linux system using
   `check_nrpe` the `check_ssh_login` plugin would be deployed to
   `/usr/lib64/nagios/plugins/check_ssh_login` or
   `/usr/local/nagios/libexec/check_ssh_login`

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

## Configuration

### Command-line arguments

- Use the `-h` or `--help` flag to display current usage information.
- Flags marked as **`required`** must be set via CLI flag.
- Flags *not* marked as required are for settings where a useful default is
  already defined, but may be overridden if desired.

#### `check_ssh_login`

| Flag                  | Required | Default    | Repeat | Possible                                                                | Description                                                                               |
| --------------------- | -------- | ---------- | ------ | ----------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `h`, `help`           | No       | `false`    | No     | `h`, `help`                                                             | Show Help text along with the list of supported flags.                                    |
| `ll`, `log-level`     | No       | `info`     | No     | `disabled`, `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace` | Log message priority filter. Log messages with a lower level are ignored.                 |
| `command`             | No       | *empty*    | No     | *valid shell command*                                                   | The valid shell command to execute via SSH.                                               |
| `login-failure-state` | No       | `critical` | No     | `ok`, `warning`, `critical`, `unknown`, `dependent`                     | The plugin state to use when a SSH login fails.                                           |
| `login-success-state` | No       | `critical` | No     | `ok`, `warning`, `critical`, `unknown`, `dependent`                     | The plugin state to use when a SSH login is successful.                                   |
| `net-type`            | No       | `auto`     | No     | `tcp4`, `tcp6`, `auto`                                                  | Limits network connections to one of tcp4 (IPv4-only), tcp6 (IPv6-only) or auto (either). |
| `password`            | Yes      | *empty*    | No     | *valid password*                                                        | The valid password for the specified user.                                                |
| `port`                | No       | `22`       | No     | *positive whole number between 1-65535, inclusive*                      | The port used by the SSH service.                                                         |
| `server`              | Yes      | *empty*    | No     | *fully-qualified domain name or IP Address*                             | The SSH server FQDN or IP Address.                                                        |
| `username`            | Yes      | *empty*    | No     | *valid user account*                                                    | The valid user for the given SSH server.                                                  |
| `verbose`             | No       | `false`    | No     | `verbose`                                                               | Whether to display verbose details in the final plugin output.                            |
| `version`             | No       | `false`    | No     | `version`                                                               | Whether to display application version and then immediately exit application.             |

## Examples

### Successful authentication

#### `OK` result

This output is emitted by the plugin when a successful authentication occurs
*and* one of:

- the default login success state is not overridden
- the sysadmin has explicitly specified that a successful authentication is to
  be considered an `OK` plugin state

Both of the following examples are equivalent (due to the default values for the
`login-failure-state` and `login-success-state` flags).

Explicit flag values:

```console
$ ./check_ssh_login --server 192.168.236.140 --port 22 --username $username --password $password --login-failure-state "critical" --login-success-state "ok"
OK: SSH authentication successful (as expected)


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: root
* NetworkType: auto
* LoginFailureState: CRITICAL
* LoginSuccessState: OK

------


NOTE: SSH command execution not requested.

 | 'time'=190ms;;;;
```

Default flag values:

```console
$ ./check_ssh_login --server 192.168.236.140 --port 22 --username $username --password $password
OK: SSH authentication successful (as expected)


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: root
* NetworkType: auto
* LoginFailureState: CRITICAL
* LoginSuccessState: OK

------


NOTE: SSH command execution not requested.

 | 'time'=190ms;;;;
```

Regarding the output:

- The `NetworkType` is listed as `auto`
  - this indicates that the plugin automatically negotiated whether an IPv4 or
    IPv6 network connection would be used
  - a sysadmin can explicitly specify whether `auto` (based on DNS query
    result), IPv4 or IPv6 network connection type is used
- The default (or explicitly specified) plugin exit states chosen for
  successful and failed authentication attempts is listed for easy reference
  - this is intended to help explain why a particular exit state value was
    chosen
- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these
  metrics may be collected and exposed as graphs/charts

#### non-OK result

This output is emitted by the plugin when a successful authentication occurs
*and* the sysadmin has explicitly specified that a successful authentication
attempt is *not* to be considered an `OK` plugin state.

This is useful for asserting that an unauthorized account is not allowed to
login to a target system. If an authentication attempt is successful the
sysadmin wishes to be notified as this may represent a configuration
regression on the monitored system.

```console
$ ./check_ssh_login --server 192.168.236.140 --port 22 --username $username --password $password --login-failure-state "ok" --login-success-state "critical"
5:28AM ERR cmd\check_ssh_login\main.go:115 > Successfully authenticated to SSH server (UNEXPECTED) error="successfully authenticated" app_type=plugin logging_level=info net_type=auto port=22 server=root version="check-ssh devbuild (https://github.com/atc0005/check-ssh)"
CRITICAL: SSH authentication successful (UNEXPECTED)

**ERRORS**

* expected login failure: successfully authenticated

**DETAILED INFO**


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: root
* NetworkType: auto
* LoginFailureState: OK
* LoginSuccessState: CRITICAL


 | 'time'=216ms;;;;
```

Regarding the output:

- The first line of output is error output emitted to `stderr`
  - this can be muted by disabling log output via CLI flag
- The `NetworkType` is listed as `auto`
  - this indicates that the plugin automatically negotiated whether an IPv4 or
    IPv6 network connection would be used
  - a sysadmin can explicitly specify whether `auto` (based on DNS query
    result), IPv4 or IPv6 network connection type is used
- The default (or explicitly specified) plugin exit states chosen for
  successful and failed authentication attempts is listed for easy reference
  - this is intended to help explain why a particular exit state value was
    chosen
- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these
  metrics may be collected and exposed as graphs/charts

### Failed authentication

#### `OK` result

This output is emitted by the plugin when a failed authentication occurs *and*
the sysadmin has explicitly specified that a failed authentication attempt is
to be considered an `OK` plugin state.

This is useful for asserting that an unauthorized account is not allowed to
login to a target system.

Here we emulate that scenario by specifying `fakeuser` as our user account.

```console
./check_ssh_login --server 192.168.236.140 --port 22 --username fakeuser --password $password --login-failure-state "ok" --login-success-state "critical"
OK: SSH authentication failed (as expected)

**ERRORS**

* failed to login to 192.168.236.140: failed to establish authenticated SSH session to 192.168.236.140 (ip: 192.168.236.140, port: 22): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none password], no supported methods remain: failed to authenticate

**DETAILED INFO**


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: fakeuser
* NetworkType: auto
* LoginFailureState: OK
* LoginSuccessState: CRITICAL


 | 'time'=1892ms;;;;
```

Regarding the output:

- The `NetworkType` is listed as `auto`
  - this indicates that the plugin automatically negotiated whether an IPv4 or
    IPv6 network connection would be used
  - a sysadmin can explicitly specify whether `auto` (based on DNS query
    result), IPv4 or IPv6 network connection type is used
- The default (or explicitly specified) plugin exit states chosen for
  successful and failed authentication attempts is listed for easy reference
  - this is intended to help explain why a particular exit state value was
    chosen
- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these
  metrics may be collected and exposed as graphs/charts

#### non-OK result

This output is emitted by the plugin when a failed authentication occurs *and*
this is unexpected.

This is useful for asserting that an authorized account is allowed to login to
a target system. If the authorized user account is unable to login it
represents a service interruption/outage.

Here we emulate that scenario by specifying `fakeuser` as our user account.

```console
./check_ssh_login --server 192.168.236.140 --port 22 --username fakeuser --password $password
5:42AM ERR cmd\check_ssh_login\auth.go:164 > Failed to establish network connection to SSH server error="failed to login to 192.168.236.140: failed to establish authenticated SSH session to 192.168.236.140 (ip: 192.168.236.140, port: 22): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none password], no supported methods remain: failed to authenticate" app_type=plugin logging_level=info net_type=auto port=22 server=fakeuser version="check-ssh devbuild (https://github.com/atc0005/check-ssh)"
CRITICAL: SSH authentication failed

**ERRORS**

* failed to login to 192.168.236.140: failed to establish authenticated SSH session to 192.168.236.140 (ip: 192.168.236.140, port: 22): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none password], no supported methods remain: failed to authenticate

**DETAILED INFO**


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: fakeuser
* NetworkType: auto
* LoginFailureState: CRITICAL
* LoginSuccessState: OK


 | 'time'=3752ms;;;;
```

Regarding the output:

- The first line of output is error output emitted to `stderr`
  - this can be muted by disabling log output via CLI flag
- The `NetworkType` is listed as `auto`
  - this indicates that the plugin automatically negotiated whether an IPv4 or
    IPv6 network connection would be used
  - a sysadmin can explicitly specify whether `auto` (based on DNS query
    result), IPv4 or IPv6 network connection type is used
- The default (or explicitly specified) plugin exit states chosen for
  successful and failed authentication attempts is listed for easy reference
  - this is intended to help explain why a particular exit state value was
    chosen
- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these
  metrics may be collected and exposed as graphs/charts

### Execute shell command

#### No output

This output is emitted by the plugin when a successful authentication occurs
*and* a valid shell command is given for remote execution.

```console
$ ./check_ssh_login --server 192.168.236.140 --port 22 --username $username --password $password --command 'touch /tmp/ssh_check_login'
OK: SSH authentication and command execution successful (as expected)


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: root
* NetworkType: auto
* LoginFailureState: CRITICAL
* LoginSuccessState: OK

------


SSH command: "touch /tmp/ssh_check_login"

SSH command output: None


 | 'time'=276ms;;;;
```

Regarding the output:

- The SSH command is listed along with the output
  - in this case no output was generated by running the command
- The `NetworkType` is listed as `auto`
  - this indicates that the plugin automatically negotiated whether an IPv4 or
    IPv6 network connection would be used
  - a sysadmin can explicitly specify whether `auto` (based on DNS query
    result), IPv4 or IPv6 network connection type is used
- The default (or explicitly specified) plugin exit states chosen for
  successful and failed authentication attempts is listed for easy reference
  - this is intended to help explain why a particular exit state value was
    chosen
- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these
  metrics may be collected and exposed as graphs/charts

#### SSH command output generated

This output is emitted by the plugin when a successful authentication occurs
*and* a valid shell command is given for remote execution.

```console
$ ./check_ssh_login --server 192.168.236.140 --port 22 --username $username --password $password --command 'touch /tmp/ssh_check_login && ls -la /tmp/'
OK: SSH authentication and command execution successful (as expected)


Configuration settings:

* Server: 192.168.236.140
* Port: 22
* Username: root
* NetworkType: auto
* LoginFailureState: CRITICAL
* LoginSuccessState: OK

------


SSH command: "touch /tmp/ssh_check_login && ls -la /tmp/"

SSH command output:

total 0
drwxrwxrwt.  6 root root 165 Jun 12 05:57 .
dr-xr-xr-x. 17 root root 224 Mar  1 05:22 ..
-rw-r--r--.  1 root root   0 Jun 12 06:00 ssh_check_login
drwx------.  2 root root   6 Jun  9 10:54 vmware-root_836-2722107930
drwx------.  2 root root   6 Jun  5 03:14 vmware-root_837-3988228548
drwx------.  2 root root   6 Jun  5 03:39 vmware-root_838-2730562456
drwx------.  2 root root   6 Jun  6 09:12 vmware-root_846-2697139606



 | 'time'=212ms;;;;
```

Regarding the output:

- The SSH command is listed along with the output
  - some output may be problematic as [Nagios disallows some characters in
    generated output][nagios-illegal-macro-chars]
- The `NetworkType` is listed as `auto`
  - this indicates that the plugin automatically negotiated whether an IPv4 or
    IPv6 network connection would be used
  - a sysadmin can explicitly specify whether `auto` (based on DNS query
    result), IPv4 or IPv6 network connection type is used
- The default (or explicitly specified) plugin exit states chosen for
  successful and failed authentication attempts is listed for easy reference
  - this is intended to help explain why a particular exit state value was
    chosen
- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these

  metrics may be collected and exposed as graphs/charts

## License

See the [LICENSE](LICENSE) file for details.

## References

- Related projects
  - <https://github.com/atc0005/check-cert>
  - <https://github.com/atc0005/check-illiad>
  - <https://github.com/atc0005/check-mail>
  - <https://github.com/atc0005/check-process>
  - <https://github.com/atc0005/check-restart>
  - <https://github.com/atc0005/check-statuspage>
  - <https://github.com/atc0005/check-vmware>
  - <https://github.com/atc0005/check-whois>
  - <https://github.com/atc0005/send2teams>
  - <https://github.com/atc0005/nagios-debug>
  - <https://github.com/atc0005/go-nagios>

- Documentation
  - <https://pkg.go.dev/golang.org/x/crypto/ssh>

- Examples
  - <https://linuxhint.com/golang-ssh-examples/>
  - <https://gist.github.com/ilmanzo/9cf5ed25ea3bb5ba7e588ffb95ab1940>
  - <https://github.com/Scalingo/go-ssh-examples/blob/master/client.go>
  - <https://blog.tarkalabs.com/ssh-recipes-in-go-part-one-5f5a44417282?gi=573433453c9d>
  - <https://stackoverflow.com/questions/24437809/connect-to-a-server-using-ssh-and-a-pem-key-with-golang>

- Logging
  - <https://github.com/rs/zerolog>

- Nagios
  - <https://github.com/atc0005/go-nagios>
  - <https://nagios-plugins.org/doc/guidelines.html>
  - <https://www.monitoring-plugins.org/doc/guidelines.html>
  - <https://icinga.com/docs/icinga-2/latest/doc/05-service-monitoring/>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/check-ssh>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

[nagios-illegal-macro-chars]: <https://serverfault.com/questions/242357/need-to-have-illegal-characters-in-nagios-serviceoutput-and-longserviceoutput> "Nagios Illegal Macro Output Characters"

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Traceroute.md) Traceroute | [HOME](../README.md) | GNU InetUtils [>>](./GNUInetUtils.md) |

## Hostname

Name: Hostname<br />
Summary: A set of commands to query the host's name, DNS domain, and NIS domain<br />
License: GPL v2 or later<br />
Version: 3.25<br />
URL: [https://salsa.debian.org/meskes/hostname/-/archive/debian/3.25/hostname-debian-3.25.tar.gz](https://salsa.debian.org/meskes/hostname/-/archive/debian/3.25/hostname-debian-3.25.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 264 KiB<br />

## Configuration

The Hostname package has no configuration script.

## Compilation and Installation

To compile Hostname, run the following command:

```bash
make
```

Finally, to install Hostname into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | dnsdomainname domainname hostname nisdomainname ypdomainname |

| Navigation |||
| --- | --- | ---: |
| [<<](./Traceroute.md) Traceroute | [HOME](../README.md) | GNU InetUtils [>>](./GNUInetUtils.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./NetTools.md) Net Tools | [HOME](../README.md) | Hostname [>>](./Hostname.md) |

## Traceroute

Name: Traceroute<br />
Summary: TCP/IP Packet route reporting utility<br />
License: GPL v2 or later<br />
Version: 2.1.6<br />
URL: [https://sourceforge.net/projects/traceroute/files/traceroute/traceroute-2.1.6/traceroute-2.1.6.tar.gz](https://sourceforge.net/projects/traceroute/files/traceroute/traceroute-2.1.6/traceroute-2.1.6.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 264 KiB<br />

## Pre-configuration

To fix a couple bugs and allow the autoconf/automake tools to work correctly, apply the following patches:

```bash
patch -p1 <../patches/traceroute/traceroute-autotools.patch
patch -p1 <../patches/traceroute/traceroute-secure_getenv.patch
```

## Configuration

To configure Traceroute for install into the build root, run the following commands:

```bash
autoreconf -fiv
./configure --prefix=/usr                            \
            --docdir=/usr/share/doc/traceroute-2.1.6 \
            --enable-year2038
```

## Compilation and Installation

To compile Traceroute, run the following command:

```bash
make
```

Finally, to install Traceroute into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | traceroute |

| Navigation |||
| --- | --- | ---: |
| [<<](./NetTools.md) Net Tools | [HOME](../README.md) | Hostname [>>](./Hostname.md) |

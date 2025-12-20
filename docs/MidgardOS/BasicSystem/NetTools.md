# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Expat_32bit.md) Expat 32-bit | [HOME](../README.md) | Traceroute [>>](./Traceroute.md) |

## Net Tools

Name: Net Tools<br />
Summary: Tools for network configuration<br />
License: GPL v2<br />
Version: 2.10<br />
URL: [https://sourceforge.net/projects/net-tools/files/net-tools-2.10.tar.xz](https://sourceforge.net/projects/net-tools/files/net-tools-2.10.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 2.1 MiB<br />

## Pre-Configuration

The Net Tools sources have a number of security bugs that need patched. To do so, apply the following patches:

```bash
patch -p1 <../patches/net-tools/0007-Introduce-T-notrim-option-in-netstat.patch
patch -p1 <../patches/net-tools/net-tools-CVE-2025-46836.patch
patch -p1 <../patches/net-tools/net-tools-CVE-2025-46836-regression.patch
patch -p1 <../patches/net-tools/net-tools-CVE-2025-46836-error-reporting.patch
patch -p1 <../patches/net-tools/net-tools-parse_hex-stack-overflow.patch
patch -p1 <../patches/net-tools/net-tools-proc_gen_fmt-buffer-overflow.patch
patch -p1 <../patches/net-tools/net-tools-ifconfig-avoid-unsafe-memcpy.patch
patch -p1 <../patches/net-tools/net-tools-ax25+netrom-overflow-1.patch
patch -p1 <../patches/net-tools/net-tools-ax25+netrom-overflow-2.patch
patch -p1 <../patches/net-tools/net-tools-ifconfig-long-name-warning.patch
```

## Configuration

To configure Net Tools for install into the build root, run the following command:

```bash
make config
```

## Compilation and Installation

To compile Net Tools, run the following command:

```bash
make
```

Finally, to install Net Tools into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/bin/dnsdomainname
rm -fv DESTDIR/bin/domainname
rm -fv DESTDIR/bin/hostname
rm -fv DESTDIR/bin/nisdomainname
rm -fv DESTDIR/bin/ypdomainname
rm -fv DESTDIR/sbin/mii-tool
rm -fv DESTDIR/sbin/rarp
rm -rfv DESTDIR/usr/share/man/*/man1
rm -rfv DESTDIR/usr/share/man/man1
rm -fv DESTDIR/usr/share/man/*/man8/mii-tool.8
rm -fv DESTDIR/usr/share/man/*/man8/rarp.8
rm -fv DESTDIR/usr/share/man/man8/mii-tool.8
rm -fv DESTDIR/usr/share/man/man8/rarp.8
cp -Rv DESTDIR/bin/* /usr/bin/
cp -Rv DESTDIR/sbin/* /usr/sbin/
cp -Rv DESTDIR/usr/share/* /usr/share/
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | arp ifconfig ipmaddr iptunnel nameif netstat plipconfig route slattach |

| Navigation |||
| --- | --- | ---: |
| [<<](./Expat_32bit.md) Expat 32-bit | [HOME](../README.md) | Traceroute [>>](./Traceroute.md) |

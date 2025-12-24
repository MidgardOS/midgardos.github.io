# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GZip.md) GZip | [HOME](../README.md) | Kbd [>>](./Kbd.md) |

## IPRoute2

Name: IPRoute2<br />
Summary: Tools for basic and advanced networking<br />
License: $PKG_LICENSE<br />
Version: 6.18.0<br />
URL: [https://www.kernel.org/pub/linux/utils/net/iproute2/iproute2-6.18.0.tar.xz](https://www.kernel.org/pub/linux/utils/net/iproute2/iproute2-6.18.0.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 3.8 MiB<br />

## Notes

The IPRoute2 package has the option to add extra functionality by being rebuilt after installing the following packages:

- Berkeley DB
- LibBPF
- NetFilter LibMNL

## Configuration

To configure the IPRoute2 package, run the following command:

```bash
./configure --prefix=/usr       \
            --libdir=/usr/lib64 \
            --color=auto
```

## Compilation and Installation

To compile IPRoute2, run the following command:

```bash
sed -i /ARPD/d Makefile
rm -fv man/man8/arpd.8
make NETNS_RUN_DIR=/run/netns
```

This package does not at this time have a working test suite.

Finally, to install IPRoute2 into the build tree, run the following command:

```bash
make SBINDIR=/usr/sbin install
install -v -D -m644 -o root -g root COPYING README* -t /usr/share/doc/iproute2-6.18.0
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | bridge, ctstat, genl, ifstat, ip, lnstat, nstat, routel, rtacct, rtmon, rtstat, ss, and tc |

| Navigation |||
| --- | --- | ---: |
| [<<](./GZip.md) GZip | [HOME](../README.md) | Kbd [>>](./Kbd.md) |

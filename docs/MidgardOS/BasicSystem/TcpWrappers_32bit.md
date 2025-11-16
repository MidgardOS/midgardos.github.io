# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./TcpWrappers_64bit.md) TCP Wrappers 64-bit | [HOME](../README.md) | GNU Autoconf [>>](./GNUAutoconf.md) |

## TCP Wrappers 32-bit

Name: TCP Wrappers 32-bit<br />
Summary: Access and authorization wrapper for TCP daemons<br />
License: BSD<br />
Version: 7.6<br />
URL: [https://github.com/tcp-wrappers/code/archive/refs/tags/tcp_wrappers_7.6-ipv6.4.tar.gz](https://github.com/tcp-wrappers/code/archive/refs/tags/tcp_wrappers_7.6-ipv6.4.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 312 KiB<br />

## Configuration

The TCP Wrappers sources do not have a traditional configuration script.

## Compilation and Installation

To compile TCP Wrappers 32-bit, run the following commands:

```bash
make clean
rm -fv libwrap.so libwrap.so.0
CC="gcc -m32 -march=i686" make REAL_DAEMON_DIR=/usr/sbin MAJOR=0 MINOR=7 REL=6 linux
```

Finally, to install TCP Wrappers 32-bit into the build tree, run the following commands:

```bash
BUILD_ROOT=DESTDIR
rm -rfv ${BUILD_ROOT}
mkdir -pv ${BUILD_ROOT}/usr/lib
cp -Rv libwrap.so* ${BUILD_ROOT}/usr/lib/
cp -Rv ${BUILD_ROOT}/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of TCP Wrappers for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./TcpWrappers_64bit.md) TCP Wrappers 64-bit | [HOME](../README.md) | GNU Autoconf [>>](./GNUAutoconf.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./tcp_wrappers64bit.md) tcp_wrappers 64-bit | [HOME](../README.md) | audit 64-bit [>>](./audit64bit.md) |

## tcp_wrappers 32-bit

Name: tcp_wrappers 32-bit<br />
Summary: Access and authorization wrapper for TCP daemons<br />
License: BSD<br />
Version: 7.6<br />
URL: [https://github.com/tcp-wrappers/code/archive/refs/tags/tcp_wrappers_7.6-ipv6.4.tar.gz](https://github.com/tcp-wrappers/code/archive/refs/tags/tcp_wrappers_7.6-ipv6.4.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 312 KiB<br />

## Configuration

The TCP Wrappers sources do not have a traditional configuration script.

## Compilation and Installation

To compile tcp_wrappers 32-bit, run the following commands:

```bash
make clean
CC="gcc -m32 -march=i686" make REAL_DAEMON_DIR=/usr/sbin MAJOR=0 MINOR=7 REL=6 linux
```

Finally, to install tcp_wrappers 32-bit into the build tree, run the following command:

```bash
BUILD_ROOT=DESTDIR
rm -rfv ${BUILD_ROOT}
mkdir -pv ${BUILD_ROOT}/usr/lib
cp -Rv libwrap.so* ${BUILD_ROOT}/usr/lib/
cp -Rv ${BUILD_ROOT}/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of tcp_wrappers for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./tcp_wrappers64bit.md) tcp_wrappers 64-bit | [HOME](../README.md) | audit 64-bit [>>](./audit64bit.md) |

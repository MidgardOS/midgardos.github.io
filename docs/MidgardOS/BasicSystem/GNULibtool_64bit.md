# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUAutoconf.md) GNU Autoconf | [HOME](../README.md) | GNU Libtool 32-bit [>>](./GNULibtool_32bit.md) |

# GNU Libtool 64-bit

Name: GNU Libtool 64-bit<br />
Summary: A tool to build shared libraries<br />
License: GPL v3+<br />
Version: 2.5.4<br />
URL: [https://ftp.gnu.org/pub/gnu/libtool/libtool-2.5.4.tar.xz](https://ftp.gnu.org/pub/gnu/libtool/libtool-2.5.4.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.8 MiB<br />

## Configuration

To configure GNU Libtool 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --sysconfdir=/etc       \
            --libexecdir=/usr/lib64 \
            --disable-static
```

## Compilation and Installation

To compile GNU Libtool 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Note that there are three known test failures for tests 72, 119, and 172.

Finally, to install GNU Libtool 64-bit into the tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libltdl.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | libtool and libtoolize |
| Installed Libraries | libltdl.so |
| Installed Files | /usr/share/libtool |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUAutoconf.md) GNU Autoconf | [HOME](../README.md) | GNU Libtool 32-bit [>>](./GNULibtool_32bit.md) |

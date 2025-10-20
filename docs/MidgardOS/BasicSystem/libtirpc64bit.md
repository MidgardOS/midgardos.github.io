# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libcap-ng32bit.md) libcap-ng 32-bit | [HOME](../README.md) | libtirpc 32-bit [>>](./libtirpc32bit.md) |

## libtirpc 64-bit

Name: libtirpc 64-bit<br />
Summary: The library for the RPC API that used to be part of GLibC<br />
License: BSD<br />
Version: 1.3.7<br />
URL: [https://sourceforge.net/projects/libtirpc/files/libtirpc/1.3.7/libtirpc-1.3.7.tar.bz2/download](https://sourceforge.net/projects/libtirpc/files/libtirpc/1.3.7/libtirpc-1.3.7.tar.bz2/download)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.2 MiB<br />

## Configuration

To configure the libtirpc 64-bit package for compilation, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --sysconfdir=/etc       \
            --enable-rpcdb          \
            --enable-authdes        \
            --disable-static        \
            --disable-gssapi
```

## Compilation and Installation

To compile libtirpc 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install libtirpc 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libtirpc.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libtirpc.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./libcap-ng32bit.md) libcap-ng 32-bit | [HOME](../README.md) | libtirpc 32-bit [>>](./libtirpc32bit.md) |

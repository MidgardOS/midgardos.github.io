# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libtirpc64bit.md) libtirpc 64-bit | [HOME](../README.md) | libnsl2 64-bit [>>](./libnsl264bit.md) |

## libtirpc 32-bit

Name: libtirpc 32-bit<br />
Summary: The library for the RPC API that used to be part of GLibC<br />
License: BSD<br />
Version: 1.3.7<br />
URL: [https://sourceforge.net/projects/libtirpc/files/libtirpc/1.3.7/libtirpc-1.3.7.tar.bz2/download](https://sourceforge.net/projects/libtirpc/files/libtirpc/1.3.7/libtirpc-1.3.7.tar.bz2/download)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.2 MiB<br />

## Configuration

To configure the libtirpc 32-bit package for compilation, run the following command:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --sysconfdir=/etc        \
            --enable-rpcdb           \
            --enable-authdes         \
            --disable-static         \
            --disable-gssapi
```

## Compilation and Installation

To compile libtirpc 32-bit, run the following command:

```bash
make
```

Finally, to install libtirpc 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib
```

## Contents

See the contents section of the 64-bit build of libtirpc for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libtirpc64bit.md) libtirpc 64-bit | [HOME](../README.md) | libnsl2 64-bit [>>](./libnsl264bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNettle_64bit.md) GNU Nettle 64-bit | [HOME](../README.md) | GNU LibTasn1 64-bit [>>](./GNULibTasn1_64bit.md) |

## GNU Nettle 32-bit

Name: GNU Nettle 32-bit<br />
Summary: A low-level cryptography library<br />
License: LGPL v3 or later/GPL v2 or later<br />
Version: 3.10.2<br />
URL: [https://ftp.gnu.org/gnu/nettle/nettle-3.10.2.tar.gz](https://ftp.gnu.org/gnu/nettle/nettle-3.10.2.tar.gz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 3.8 MiB<br />

## Configuration

To configure GNU Nettle 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" CXX="g++ -m32 -march=i686"    \
./configure --host=i686-pc-linux-gnu                    \
            --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
            --docdir=/usr/share/doc/nettle-3.10.2       \
            --disable-static                            \
            --enable-shared                             \
            --enable-fat
```

## Compilation and Installation

To compile GNU Nettle 32-bit, run the following command:

```bash
make
```

Finally, to install GNU Nettle 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
for LIB in "libhogweed.so.6.11" "libnettle.so.8.11"; do
    chmod -v 755 DESTDIR/usr/lib/$LIB
done
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU Nettle for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNettle_64bit.md) GNU Nettle 64-bit | [HOME](../README.md) | GNU LibTasn1 64-bit [>>](./GNULibTasn1_64bit.md) |

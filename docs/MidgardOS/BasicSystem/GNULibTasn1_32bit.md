# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibTasn1_64bit.md) GNU LibTasn1 64-bit | [HOME](../README.md) | GNU TLS 64-bit [>>](./GNUTLS_64bit.md) |

## GNU LibTasn1 32-bit

Name: GNU LibTasn1 32-bit<br />
Summary: An ASN.1 implementation library<br />
License: GPL v3 or later/LGPL v2.1 or later<br />
Version: 4.20.0<br />
URL: [https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.20.0.tar.gz](https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.20.0.tar.gz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 708 KiB<br />

## Configuration

To configure GNU LibTasn1 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" \
./configure --host=i686-pc-linux-gnu    \
            --prefix=/usr               \
            --libdir=/usr/lib           \
            --libexecdir=/usr/lib       \
            --disable-static            \
            --enable-shared
```

## Compilation and Installation

To compile GNU LibTasn1 32-bit, run the following command:

```bash
make
```

Finally, to install GNU LibTasn1 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libtasn1.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU LibTasn1 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibTasn1_64bit.md) GNU LibTasn1 64-bit | [HOME](../README.md) | GNU TLS 64-bit [>>](./GNUTLS_64bit.md) |

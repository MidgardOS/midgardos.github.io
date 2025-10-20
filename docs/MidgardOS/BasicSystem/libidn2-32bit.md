# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL64bit.md) ISL 64-bit | [HOME](../README.md) | Attr 64-bit [>>](./Attr64bit.md) |

## GNU LibIDN2 64-bit

Name: GNU LibIDN2 64-bit<br />
Summary: Version 2 of the Internationalized Domain Name library<br />
License: LGPL v3 or later<br />
Version: 2.3.8<br />
URL: [https://ftp.gnu.org/gnu/libidn/libidn2-2.3.8.tar.gz](https://ftp.gnu.org/gnu/libidn/libidn2-2.3.8.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 21 MiB<br />

## Configuration

To configure GNU LibIDN2 64-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --disable-static
```

## Compilation and Installation

To compile GNU LibIDN2 64-bit, run the following command:

```bash
make
```

Finally, to install GNU LibIDN2 64-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libidn2.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU LibIDN2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL64bit.md) ISL 64-bit | [HOME](../README.md) | Attr 64-bit [>>](./Attr64bit.md) |

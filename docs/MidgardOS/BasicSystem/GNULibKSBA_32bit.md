# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibKSBA_32bit.md) GNU LibKSBA 64-bit | [HOME](../README.md) | GNU LibAssuan 64-bit [>>](./GNULibAssuan_64bit.md) |

## GNU LibKSBA 32-bit

Name: GNU LibKSBA 32-bit<br />
Summary: A library for manipulating X.509 certificates<br />
License: GPL v2 and later/LGPL v3 and later<br />
Version: 1.6.7<br />
URL: [https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2](https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2)<br />

Average Build Time: less than 0,1 SBU<br />
Used Install Space: 1.3 MiB<br />

## Configuration

To configure GNU LibKSBA 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32 -march=i686"                           \
./configure --host=i686-pc-linux-gnu                \
            --prefix=/usr                           \
            --libdir=/usr/lib                       \
            --libexecdir=/usr/lib                   \
            --docdir=/usr/share/doc/libksba-1.6.7   \
            --disable-static
```

## Compilation and Installation

To compile GNU LibKSBA 32-bit, run the following command:

```bash
make
```

Finally, to install GNU LibKSBA 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU LibKSBA for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibKSBA_32bit.md) GNU LibKSBA 64-bit | [HOME](../README.md) | GNU LibAssuan 64-bit [>>](./GNULibAssuan_64bit.md) |

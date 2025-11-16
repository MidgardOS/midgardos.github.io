# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibUniString_64bit.md) libunistring 64-bit | [HOME](../README.md) | File 64-bit [>>](./File_64bit.md) |

## libunistring 32-bit

Name: libunistring 32-bit<br />
Summary: Unicode string handling in C<br />
License: GPL v3 or later<br />
Version: 1.3<br />
URL: [https://ftp.gnu.org/gnu/libunistring/libunistring-1.3.tar.xz](https://ftp.gnu.org/gnu/libunistring/libunistring-1.3.tar.xz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 58 MiB<br />

## Configuration

To configure libunistring 32-bit for install into the build root, run the following command:

```bash
make clean
CC="gcc -m32" \
./configure --prefix=/usr         \
            --libdir=/usr/lib     \
            --libexecdir=/usr/lib \
            --disable-static      \
            --docdir=/usr/share/doc/libunistring-1.3
```

## Compilation and Installation

To compile libunistring 32-bit, run the following command:

```bash
make
```

Finally, to install libunistring 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
rm -fv /usr/lib/libunistring.la
```

## Contents

See the contents section of the 64-bit build of libunistring for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibUniString_64bit.md) libunistring 64-bit | [HOME](../README.md) | File 64-bit [>>](./File_64bit.md) |

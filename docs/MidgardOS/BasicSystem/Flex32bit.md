# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Flex64bit.md) Flex 64-bit | [HOME](../README.md) | TCL [>>](./TCL.md) |

## Flex 32-bit

Name: Flex 32-bit<br />
Summary: The fast lexer<br />
License: BSD-3-Clause<br />
Version: 2.6.4<br />
URL: [https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz](https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 33 MiB<br />

## Configuration

To configure Flex 32-bit for install into the build root, run the following command:

```bash
CC="gcc -m32" ./configure                       \
            --host=i686-pc-linux-gnu            \
            --prefix=/usr                       \
            --libdir=/usr/lib                   \
            --libexecdir=/usr/lib               \
            --docdir=/usr/share/doc/flex-2.6.4  \
            --disable-static
```

## Compilation and Installation

To compile Flex 32-bit, run the following command:

```bash
make
```

Finally, to install Flex 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp DESTDIR/usr/lib/* /usr/lib
rm -fv /usr/lib/libfl.la
```

**NOTE: Do not delete the unpacked sources**

## Contents

See the contents section of the 64-bit build of Flex for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./Flex64bit.md) Flex 64-bit | [HOME](../README.md) | TCL [>>](./TCL.md) |

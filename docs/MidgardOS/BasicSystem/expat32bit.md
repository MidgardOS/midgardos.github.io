# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./expat64bit.md) Expat 64-bit | [HOME](../README.md) | Net Utils [>>](./netutils.md) |

## Expat 32-bit

Name: Expat 32-bit<br />
Summary: An XML stream parsing library<br />
License: MIT<br />
Version: 2.7.3<br />
URL: [https://github.com/libexpat/libexpat/releases/download/R_2_7_3/expat-2.7.3.tar.xz](https://github.com/libexpat/libexpat/releases/download/R_2_7_3/expat-2.7.3.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 14 MiB<br />

## Configuration

To configure Expat 32-bit for install into the build root, run the following commands:

```bash
make distclean
./configure --prefix=/usr         \
            --libdir=/usr/lib     \
            --libexecdir=/usr/lib \
            --disable-static      \
            --docdir=/usr/share/doc/expat-2.7.3
```

## Compilation and Installation

To compile Expat 32-bit, run the following command:

```bash
make
```

Finally, to install Expat 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of Expat for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./expat64bit.md) Expat 64-bit | [HOME](../README.md) | Net Utils [>>](./netutils.md) |

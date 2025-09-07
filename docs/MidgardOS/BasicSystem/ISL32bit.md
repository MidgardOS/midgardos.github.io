# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL64bit.md) ISL 64-bit | [HOME](../README.md) | GNU LibIDN2 64-bit [>>](./libidn2-64bit.md) |

# Integer Set Library

Name: ISL 32-bit<br />
Summary: A library for handling Integer Sets bounded by linear constraints<br />
License: MIT<br />
Version: 0.27<br />
URL: [https://libisl.sourceforge.io/isl-0.27.tar.xz](https://libisl.sourceforge.io/isl-0.27.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 20 MiB<br />

## Configuration

To configure ISL 32-bit for install into our cross-compilation root, run the following command:

```bash
make distclean
CC="gcc -m32" CXX="g++ -m32" \
./configure --host=i686-midgardos-linux-gnu \
            --prefix=/usr                   \
            --libdir=/usr/lib               \
            --libexecdir=/usr/lib           \
            --disable-static                \
            --docdir=/usr/share/doc/isl-0.27
```

## Compilation and Installation

To compile ISL 32-bit, run the following commands:

```bash
make
```

Finally, to install ISL 32-bit into the cross-tools tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libisl.la
rm -fv DESTDIR/usr/lib/libisl.*-gdb.py
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU MPC for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL64bit.md) ISL 64-bit | [HOME](../README.md) | GNU LibIDN2 64-bit [>>](./libidn2-64bit.md) |

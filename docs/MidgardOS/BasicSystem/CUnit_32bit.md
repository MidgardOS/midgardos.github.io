# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./CUnit_64bit.md) CUnit 64-bit | [HOME](../README.md) | LibSEPol 64-bit [>>](./LibSEPol_64bit.md) |

## CUnit 32-bit

Name: CUnit 32-bit<br />
Summary: A library for writing unit tests in C<br />
License: LGPL v2.1<br />
Version: 2.1-3<br />
URL: [https://sourceforge.net/projects/cunit/files/CUnit/2.1-3/CUnit-2.1-3.tar.bz2](https://sourceforge.net/projects/cunit/files/CUnit/2.1-3/CUnit-2.1-3.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 544 KiB<br />

## Configuration

To configure CUnit 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" CXX="g++ -m32"                    \
./configure --prefix=/usr                       \
            --libdir=/usr/lib                   \
            --libexecdir=/usr/lib               \
            --disable-static                    \
            --docdir=/usr/share/doc/CUnit-2.1-3
```

## Compilation and Installation

To compile CUnit 32-bit, run the following command:

```bash
make
```

Finally, to install CUnit 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libcunit.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of the CUnit for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./CUnit_64bit.md) CUnit 64-bit | [HOME](../README.md) | LibSEPol 64-bit [>>](./LibSEPol_64bit.md) |

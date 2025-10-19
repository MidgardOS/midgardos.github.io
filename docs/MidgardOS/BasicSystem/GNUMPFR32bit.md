# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPFR64bit.md) GNU MPFR 64-bit | [HOME](../README.md) | GNU MPC 64-bit [>>](./GNUMPC64bit.md) |

# GNU Multi-Precision Floating-Point with Rounding Library

Name: MPFR 32-bit<br />
Summary: A C library for multi-precision floating-point calculations<br />
License: LGPL v3 or later<br />
Version: 4.2.2<br />
URL: [https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.2.tar.xz](https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.2.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 43 MiB<br />

## Configuration

To configure GNU MPFR 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu            \
            --prefix=/usr                       \
            --libdir=/usr/lib                   \
            --libexecdir=/usr/lib               \
            --disable-static                    \
            --enable-shared                     \
            --enable-thread-safe                \
            --with-gmp-include=/usr/include/m32 \
            --with-gmp-lib=/usr/lib             \
            --docdir=/usr/share/doc/mpfr-4.2.2
```

## Compilation and Installation

To compile GNU MPFR 32-bit, run the following command:

```bash
make
```

Finally, to install GNU MPFR 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libmpfr.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU MPFR for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPFR64bit.md) GNU MPFR 64-bit | [HOME](../README.md) | GNU MPC 64-bit [>>](./GNUMPC64bit.md) |

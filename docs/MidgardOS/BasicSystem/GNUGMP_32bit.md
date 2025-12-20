# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGMP_64bit.md) GNU GMP 64-bit | [HOME](../README.md) | GNU MPFR 64-bit [>>](./GNUMPFR_64bit.md) |

# GNU Math Precision Library

Name: GMP 32-bit<br />
Summary: A library for working with large numbers<br />
License: GPL v3 or later<br />
Version: 6.3.0<br />
URL: [http://ftp.gnu.org/pub/gnu/gmp/gmp-6.3.0.tar.xz](http://ftp.gnu.org/pub/gnu/gmp/gmp-6.3.0.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 54 MiB<br />

## Configuration

To configure GMP 32-bit for install into the build root, run the following commands:

```bash
make distclean
cp -v configfsf.guess config.guess
cp -v configfsf.sub   config.sub
ABI="32" \
CFLAGS="-m32 -O2 -pedantic -fomit-frame-pointer -mtune=generic -march=i686" \
CXXFLAGS="$CFLAGS" \
PKG_CONFIG_PATH="/usr/lib/pkgconfig" \
./configure                          \
    --host=i686-pc-linux-gnu         \
    --prefix=/usr                    \
    --libdir=/usr/lib                \
    --libexecdir=/usr/lib            \
    --disable-static                 \
    --enable-cxx                     \
    --includedir=/usr/include/m32/gmp
```

## Compilation and Installation

To compile GMP 32-bit, run the following commands:

```bash
sed -i 's/$(exec_prefix)\/include/$\(includedir\)/' Makefile
make
```

Next, run the test suite:

```bash
make check 2>&1 | tee gmp-check-log
awk '/# PASS:/{total+=$3} ; END{print total}' gmp-check-log
```

Finally, to install GMP 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib
cp -Rv DESTDIR/usr/include/m32/* /usr/include/m32/
rm -fv /usr/lib/lib{gmp,gmpxx}.la
```

## Contents

See the contents section of the 64-bit build of GNU GMP for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGMP_64bit.md) GNU GMP 64-bit | [HOME](../README.md) | GNU MPFR 64-bit [>>](./GNUMPFR_64bit.md) |

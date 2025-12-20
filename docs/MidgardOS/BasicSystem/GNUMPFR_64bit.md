# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGMP_32bit.md) GNU GMP 32-bit | [HOME](../README.md) | GNU MPFR 32-bit [>>](./GNUMPFR_32bit.md) |

# GNU Multi-Precision Floating-Point with Rounding Library

Name: MPFR 64-bit<br />
Summary: A C library for multi-precision floating-point calculations<br />
License: LGPL v3 or later<br />
Version: 4.2.2<br />
URL: [https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.2.tar.xz](https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.2.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 43 MiB<br />

## Configuration

To configure GNU MPFR 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --enable-shared         \
            --enable-thread-safe    \
            --docdir=/usr/share/doc/mpfr-4.2.2
```

## Compilation and Installation

To compile GNU MPFR 64-bit, run the following commands:

```bash
make
make html
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU MPFR 64-bit into the build tree, run the following commands:

```bash
make install
make install-html
rm -fv /usr/lib64/libmpfr.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libmpfr.so |
| Installed Directories | /usr/share/doc/mpfr-4.2.2 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGMP_32bit.md) GNU GMP 32-bit | [HOME](../README.md) | GNU MPFR 32-bit [>>](./GNUMPFR_32bit.md) |

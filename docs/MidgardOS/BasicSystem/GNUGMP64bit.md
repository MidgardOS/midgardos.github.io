# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) GNU Binutils | [HOME](../README.md) | GNU GMP 32-bit [>>](./GNUGMP32bit.md) |

# GNU Math Precision Library

Name: GMP 64-bit<br />
Summary: A library for working with large numbers<br />
License: GPL v3 or later<br />
Version: 6.3.0<br />
URL: [http://ftp.gnu.org/pub/gnu/gmp/gmp-6.3.0.tar.xz](http://ftp.gnu.org/pub/gnu/gmp/gmp-6.3.0.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 54 MiB<br />

## Configuration

To configure GMP 64-bit for install into the build root, run the following commands:

```bash
sed -i '/long long t1;/,+1s/()/(...)/' configure
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --enable-cxx                        \
            --disable-static                    \
            --docdir=/usr/share/doc/gmp-6.3.0
```

## Compilation and Installation

To compile GMP 64-bit, run the following commands:

```bash
make
make html
```

Next, run the test suite:

```bash
make check 2>&1 | tee gmp-check-log
awk '/# PASS:/{total+=$3} ; END{print total}' gmp-check-log
```

Finally, to install GMP 64-bit into the build tree, run the following commands:

```bash
make install
make install-html
rm -fv /usr/lib64/lib{gmp,gmpxx}.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libgmp.so and libgmpxx.so |
| Installed Directories | /usr/share/doc/gmp-6.3.0 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) GNU Binutils | [HOME](../README.md) | GNU GMP 32-bit [>>](./GNUGMP32bit.md) |

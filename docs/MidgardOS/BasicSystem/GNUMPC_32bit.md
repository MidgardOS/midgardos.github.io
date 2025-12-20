# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPC_64bit.md) GNU MPC 64-bit | [HOME](../README.md) | ISL 64-bit [>>](./ISL_64bit.md) |

# GNU Multi-Precision Calculation Library

Name: GNU MPC 32-bit<br />
Summary: A C library for multi-precision calculations<br />
License: LGPL v3 or later<br />
Version: 1.3.1<br />
URL: [https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.gz](https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 22 MiB<br />

## Configuration

To configure GNU MPC 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu        \
            --prefix=/usr                   \
            --libdir=/usr/lib               \
            --libexecdir=/usr/lib           \
            --disable-static                \
            --docdir=/usr/share/doc/mpc-1.3.1
```

## Compilation and Installation

To compile GNU MPC 32-bit, run the following command:

```bash
make
```

Finally, to install GNU MPC 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libmpc.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU MPC for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPC_64bit.md) GNU MPC 64-bit | [HOME](../README.md) | ISL 64-bit [>>](./ISL_64bit.md) |

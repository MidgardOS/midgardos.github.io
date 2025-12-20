# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPFR_32bit.md) GNU MPFR 32-bit | [HOME](../README.md) | GNU MPC 32-bit [>>](./GNUMPC_32bit.md) |

# GNU Multi-Precision Calculation Library

Name: GNU MPC 64-bit<br />
Summary: A C library for multi-precision calculations<br />
License: LGPL v3 or later<br />
Version: 1.3.1<br />
URL: [https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.gz](https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 22 MiB<br />

## Configuration

To configure GNU MPC 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --docdir=/usr/share/doc/mpc-1.3.1
```

## Compilation and Installation

To compile GNU MPC 64-bit, run the following commands:

```bash
make
make html
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU MPC 64-bit into the build tree, run the following commands:

```bash
make install
make install-html
rm -fv /usr/lib64/libmpc.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libmpc.so |
| Installed Directories | /usr/share/doc/mpc-1.3.1 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPFR_32bit.md) GNU MPFR 32-bit | [HOME](../README.md) | GNU MPC 32-bit [>>](./GNUMPC_32bit.md) |

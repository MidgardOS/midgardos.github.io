# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL_32bit.md) ISL 32-bit | [HOME](../README.md) | GNU LibIDN2 32-bit [>>](./libidn2_32bit.md) |

## GNU LibIDN2 64-bit

Name: GNU LibIDN2 64-bit<br />
Summary: Version 2 of the Internationalized Domain Name library<br />
License: LGPL v3 or later<br />
Version: 2.3.8<br />
URL: [https://ftp.gnu.org/gnu/libidn/libidn2-2.3.8.tar.gz](https://ftp.gnu.org/gnu/libidn/libidn2-2.3.8.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 21 MiB<br />

## Configuration

To configure GNU LibIDN2 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static
```

## Compilation and Installation

To compile GNU LibIDN2 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

The single test that fails is due to locale not being setup correctly for this test. It is safe to ignore.

Finally, to install GNU LibIDN2 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libidn2.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | idn2 |
| Installed Libraries | libidn2.so |
| Installed Directories | /usr/share/gtk-doc/html/libidn2 |

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL_32bit.md) ISL 32-bit | [HOME](../README.md) | GNU LibIDN2 32-bit [>>](./libidn2_32bit.md) |

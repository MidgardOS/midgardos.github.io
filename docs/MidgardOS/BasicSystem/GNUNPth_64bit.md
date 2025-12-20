# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibAssuan_32bit.md) GNU LibAssuan 32-bit | [HOME](../README.md) | GNU nPth 32-bit [>>](./GNUNPth_32bit.md) |

## GNU nPth 64-bit

Name: GNU nPth 64-bit<br />
Summary: A non-preemptive threading library implementation<br />
License: LGPL v2.1 or later<br />
Version: 1.8<br />
URL: [https://gnupg.org/ftp/gcrypt/npth/npth-1.8.tar.bz2](https://gnupg.org/ftp/gcrypt/npth/npth-1.8.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 96 KiB<br />

## Configuration

To configure GNU nPth 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/npth-1.8    \
            --disable-static
```

## Compilation and Installation

To compile GNU nPth 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU nPth 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libnpth.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libnpth.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibAssuan_32bit.md) GNU LibAssuan 32-bit | [HOME](../README.md) | GNU nPth 32-bit [>>](./GNUNPth_32bit.md) |

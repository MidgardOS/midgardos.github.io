# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibKSBA_32bit.md) GNU LibKSBA 32-bit | [HOME](../README.md) | GNU LibAssuan 32-bit [>>](./GNULibAssuan_32bit.md) |

## GNU LibAssuan 64-bit

Name: GNU LibAssuan 64-bit<br />
Summary: An IPC protocol library for GnuPG suite<br />
License: GPL v3 or later<br />
Version: 3.0.2<br />
URL: [https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.2.tar.bz2](https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.2.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 652 KiB<br />

## Configuration

To configure GNU LibAssuan 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                           \
            --libdir=/usr/lib64                     \
            --libexecdir=/usr/lib64                 \
            --docdir=/usr/share/doc/libassuan-3.0.2 \
            --disable-static
```

## Compilation and Installation

To compile GNU LibAssuan 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU LibAssuan 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libassuan.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | libassuan-config |
| Installed Libraries | libassuan.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibKSBA_32bit.md) GNU LibKSBA 32-bit | [HOME](../README.md) | GNU LibAssuan 32-bit [>>](./GNULibAssuan_32bit.md) |

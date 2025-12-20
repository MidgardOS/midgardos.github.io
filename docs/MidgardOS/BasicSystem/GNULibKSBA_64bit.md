# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibGCrypt_32bit.md) GNU LibGCrypt 32-bit | [HOME](../README.md) | GNU LibKSBA 32-bit [>>](./GNULibKSBA_32bit.md) |

## GNU LibKSBA 64-bit

Name: GNU LibKSBA 64-bit<br />
Summary: A library for manipulating X.509 certificates<br />
License: GPL v2 and later/LGPL v3 and later<br />
Version: 1.6.7<br />
URL: [https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2](https://gnupg.org/ftp/gcrypt/libksba/libksba-1.6.7.tar.bz2)<br />

Average Build Time: less than 0,1 SBU<br />
Used Install Space: 1.3 MiB<br />

## Configuration

To configure GNU LibKSBA 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                         \
            --libdir=/usr/lib64                   \
            --libexecdir=/usr/lib64               \
            --docdir=/usr/share/doc/libksba-1.6.7 \
            --disable-static
```

## Compilation and Installation

To compile GNU LibKSBA 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU LibKSBA 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libksba.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | ksba-config |
| Installed Libraries | libksba.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibGCrypt_32bit.md) GNU LibGCrypt 32-bit | [HOME](../README.md) | GNU LibKSBA 32-bit [>>](./GNULibKSBA_32bit.md) |

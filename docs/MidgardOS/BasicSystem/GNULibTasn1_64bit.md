# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNettle_32bit.md) GNU Nettle 32-bit | [HOME](../README.md) | GNU LibTasn1 32-bit [>>](./GNULibTasn1_32bit.md) |

## GNU LibTasn1 64-bit

Name: GNU LibTasn1 64-bit<br />
Summary: An ASN.1 implementation library<br />
License: GPL v3 or later/LGPL v2.1 or later<br />
Version: 4.20.0<br />
URL: [https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.20.0.tar.gz](https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.20.0.tar.gz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 708 KiB<br />

## Configuration

To configure GNU LibTasn1 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --enable-shared
```

## Compilation and Installation

To compile GNU LibTasn1 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU LibTasn1 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libtasn1.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | asn1Coding, asn1Decoding, and asn1Parser |
| Installed Libraries | libtasn1.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNettle_32bit.md) GNU Nettle 32-bit | [HOME](../README.md) | GNU LibTasn1 32-bit [>>](./GNULibTasn1_32bit.md) |

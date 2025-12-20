# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNPth_32bit.md) GNU nPth 32-bit | [HOME](../README.md) | GNU Nettle 32-bit [>>](./GNUNettle_32bit.md) |

## GNU Nettle 64-bit

Name: GNU Nettle 64-bit<br />
Summary: A low-level cryptography library<br />
License: LGPL v3 or later/GPL v2 or later<br />
Version: 3.10.2<br />
URL: [https://ftp.gnu.org/gnu/nettle/nettle-3.10.2.tar.gz](https://ftp.gnu.org/gnu/nettle/nettle-3.10.2.tar.gz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 3.8 MiB<br />

## Configuration

To configure GNU Nettle 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                           \
            --libdir=/usr/lib64                     \
            --libexecdir=/usr/lib64                 \
            --docdir=/usr/share/doc/nettle-3.10.2   \
            --disable-static                        \
            --enable-shared                         \
            --enable-fat
```

## Compilation and Installation

To compile GNU Nettle 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Nettle 64-bit into the build tree, run the following commands:

```bash
make install
for LIB in "libhogweed.so.6.11" "libnettle.so.8.11"; do
    chmod -v 755 /usr/lib64/$LIB
done
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | nettle-hash, nettle-lfib-stream, nettle-pbkdf2, pkcs1-conv, and sexp-conv |
| Installed Libraries | libhogweed.so and libnettle.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNPth_32bit.md) GNU nPth 32-bit | [HOME](../README.md) | GNU Nettle 32-bit [>>](./GNUNettle_32bit.md) |

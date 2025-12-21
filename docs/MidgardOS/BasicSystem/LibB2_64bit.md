# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LZO_32bit.md) LZO 32-bit | [HOME](../README.md) | LibB2 32-bit [>>](./LibB2_32bit.md) |

## LibB2 64-bit

Name: LibB2 64-bit<br />
Summary: A C library providing BLAKE2b, BLAKE2s, BLAKE2bp, BLAKE2sp hashes<br />
License: CC0 v1.0<br />
Version: 0.98.1<br />
URL: [https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.tar.gz](https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 52 KiB<br />

## Configuration

To configure LibB2 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                        \
            --libdir=/usr/lib64                  \
            --libexecdir=/usr/lib64              \
            --docdir=/usr/share/doc/libb2-0.98.1 \
            --disable-static                     \
            --disable-native                     \
            --enable-shared
```

## Compilation and Installation

To compile LibB2 64-bit, run the following command:

```bash
make
```

Next, run its test suite:

```bash
make check
```

Finally, to install LibB2 64-bit into the build tree, run the following command:

```bash
make install
install -v -m 644 -o root -g root src/blake2.h /usr/include/
rm -fv /usr/lib64/libb2.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libb2.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./LZO_32bit.md) LZO 32-bit | [HOME](../README.md) | LibB2 32-bit [>>](./LibB2_32bit.md) |

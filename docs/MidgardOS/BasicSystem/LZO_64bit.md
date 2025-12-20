# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./SGML-Common.md) SGML Common | [HOME](../README.md) | LZO 32-bit [>>](./LZO_32bit.md) |

## LZO 64-bit

Name: LZO 64-bit<br />
Summary: A fast real-time compression and decompression library<br />
License: GPL v2 or later<br />
Version: 2.10<br />
URL: [https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.tar.gz](https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.tar.gz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 880 KiB<br />

## Configuration

To configure LZO 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr            \
            --libdir=/usr/lib64      \
            --libexecdir=/usr/lib64  \
            --enable-shared          \
            --disable-static         \
            --docdir=/usr/share/doc/lzo-2.10
```

## Compilation and Installation

To compile LZO 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
make test
```

Finally, to install LZO 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/liblzo2.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | liblzo2.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./SGML-Common.md) SGML Common | [HOME](../README.md) | LZO 32-bit [>>](./LZO_32bit.md) |

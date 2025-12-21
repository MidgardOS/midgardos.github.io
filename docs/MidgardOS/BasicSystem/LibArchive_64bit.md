# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibB2_32bit.md) LibB2 32-bit | [HOME](../README.md) | LibArchive 32-bit [>>](./LibArchive_32bit.md) |

## LibArchive 64-bit

Name: LibArchive 64-bit<br />
Summary: A library that provides a single interface for working with various compression formats<br />
License: BSD 3-clause<br />
Version: 3.8.4<br />
URL: [https://github.com/libarchive/libarchive/releases/download/v3.8.4/libarchive-3.8.4.tar.xz](https://github.com/libarchive/libarchive/releases/download/v3.8.4/libarchive-3.8.4.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 4.5 MiB<br />

## Configuration

To configure LibArchive 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                            \
            --libdir=/usr/lib64                      \
            --libexecdir=/usr/lib64                  \
            --docdir=/usr/share/doc/libarchive-3.8.4 \
            --disable-static
```

## Compilation and Installation

To compile LibArchive 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install LibArchive 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libarchive.la
ln -sv bsdunzip /usr/bin/unzip
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | bsdcat, bsdcpio, bsdtar, and bsdunzip |
| Installed Libraries | libarchive.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./LZO_32bit.md) LZO 32-bit | [HOME](../README.md) | LibArchive 32-bit [>>](./LibArchive_32bit.md) |

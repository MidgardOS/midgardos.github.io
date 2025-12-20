# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPinEntry.md) GNU PinEntry | [HOME](../README.md) | LibXML2 32-bit [>>](./LibXML2_32bit.md) |

## LibXML2 64-bit

Name: LibXML2 64-bit<br />
Summary: An XML processing library<br />
License: MIT<br />
Version: 2.15.1<br />
URL: [https://download.gnome.org/sources/libxml2/2.15/libxml2-2.15.1.tar.xz](https://download.gnome.org/sources/libxml2/2.15/libxml2-2.15.1.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 5.2 MiB<br />

## Configuration

To configure LibXML2 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                           \
            --libdir=/usr/lib64                     \
            --libexecdir=/usr/lib64                 \
            --docdir=/usr/share/doc/libxml2-2.15.1  \
            --disable-static                        \
            --with-history
```

## Compilation and Installation

To compile LibXML2 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install LibXML2 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libxml2.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | xml2-config, xmlcatalog, and xmllint |
| Installed Libraries | libxml2.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPinEntry.md) GNU PinEntry | [HOME](../README.md) | LibXML2 32-bit [>>](./LibXML2_32bit.md) |

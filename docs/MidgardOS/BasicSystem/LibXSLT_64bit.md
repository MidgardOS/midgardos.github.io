# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXML2_32bit.md) LibXML2 32-bit | [HOME](../README.md) | LibXSLT 32-bit [>>](./LibXSLT_32bit.md) |

## LibXSLT 64-bit

Name: LibXSLT 64-bit<br />
Summary: An XSLT processor based on LibXML2<br />
License: MIT<br />
Version: 1.1.45<br />
URL: [https://gitlab.gnome.org/GNOME/libxslt/-/archive/v1.1.45/libxslt-v1.1.45.tar.bz2](https://gitlab.gnome.org/GNOME/libxslt/-/archive/v1.1.45/libxslt-v1.1.45.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 3.0 MiB<br />

## Configuration

To configure LibXSLT 64-bit for install into the build root, run the following commands:

```bash
NOCONFIGURE=1 ./autogen.sh
./configure --prefix=/usr                           \
            --libdir=/usr/lib64                     \
            --libexecdir=/usr/lib64                 \
            --docdir=/usr/share/doc/libxslt-1.1.45  \
            --with-crypto                           \
            --with-plugins
```

## Compilation and Installation

To compile LibXSLT 64-bit, run the following command:

```bash
make
```

Due to a circular dependency with LibXML2 and LibXSLT, the test suite cannot be run until LibXML2 is rebuilt later.

Finally, to install LibXSLT 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libexslt.la
rm -fv /usr/lib64/libxslt.la
rm -fv /usr/lib64/python3.13/site-packages/libxsltmod.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | xslt-config and xsltproc |
| Installed Libraries | libexslt.so and libxslt.so |
| Installed Plugins | libxsltmod and libxslt Python3 modules |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXML2_32bit.md) LibXML2 32-bit | [HOME](../README.md) | LibXSLT 32-bit [>>](./LibXSLT_32bit.md) |

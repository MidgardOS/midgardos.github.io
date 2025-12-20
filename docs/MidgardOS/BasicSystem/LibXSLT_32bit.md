# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXSLT_64bit.md) LibXSLT 64-bit | [HOME](../README.md) | SGML Common [>>](./SGML-Common.md) |

## LibXSLT 32-bit

Name: LibXSLT 32-bit<br />
Summary: An XSLT processor based on LibXML2<br />
License: MIT<br />
Version: 1.1.45<br />
URL: [https://gitlab.gnome.org/GNOME/libxslt/-/archive/v1.1.45/libxslt-v1.1.45.tar.bz2](https://gitlab.gnome.org/GNOME/libxslt/-/archive/v1.1.45/libxslt-v1.1.45.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 2.3 MiB<br />

## Configuration

To configure LibXSLT 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32 -march=i686" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr                           \
            --libdir=/usr/lib                       \
            --libexecdir=/usr/lib                   \
            --docdir=/usr/share/doc/libxslt-1.1.45  \
            --with-crypto                           \
            --with-plugins                          \
            --without-python
```

Note that python is disabled in 32bit builds, since Python is built 64-bit only for Midgard OS.

## Compilation and Installation

To compile LibXSLT 32-bit, run the following command:

```bash
make
```

Finally, to install LibXSLT 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libexslt.la
rm -fv DESTDIR/usr/lib/libxslt.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of the LibXSLT 32-bit for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXSLT_64bit.md) LibXSLT 64-bit | [HOME](../README.md) | SGML Common [>>](./SGML-Common.md) |

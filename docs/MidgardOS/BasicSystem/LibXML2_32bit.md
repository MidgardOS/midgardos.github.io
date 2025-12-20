# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXML2_64bit.md) LibXML2 64-bit | [HOME](../README.md) | LibXSLT 64-bit [>>](./LibXSLT_64bit.md) |

## LibXML2 32-bit

Name: LibXML2 32-bit<br />
Summary: An XML processing library<br />
License: MIT<br />
Version: 2.15.1<br />
URL: [https://download.gnome.org/sources/libxml2/2.15/libxml2-2.15.1.tar.xz](https://download.gnome.org/sources/libxml2/2.15/libxml2-2.15.1.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 5.2 MiB<br />

## Configuration

To configure LibXML2 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686"                           \
./configure --host=i686-pc-linux-gnu                \
            --prefix=/usr                           \
            --libdir=/usr/lib                       \
            --libexecdir=/usr/lib                   \
            --docdir=/usr/share/doc/libxml2-2.15.1  \
            --disable-static                        \
            --with-history
```

## Compilation and Installation

To compile LibXML2 32-bit, run the following command:

```bash
make
```

Finally, to install LibXML2 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libxml2.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of LibXML2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXML2_64bit.md) LibXML2 64-bit | [HOME](../README.md) | LibXSLT 64-bit [>>](./LibXSLT_64bit.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibArchive_64bit.md) LibArchive 64-bit | [HOME](../README.md) | Docbook 3.1 SGML DTDs [>>](./Docbook31-SGML-DTD.md) |

## LibArchive 32-bit

Name: LibArchive 32-bit<br />
Summary: A library that provides a single interface for working with various compression formats<br />
License: BSD 3-clause<br />
Version: 3.8.4<br />
URL: [https://github.com/libarchive/libarchive/releases/download/v3.8.4/libarchive-3.8.4.tar.xz](https://github.com/libarchive/libarchive/releases/download/v3.8.4/libarchive-3.8.4.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 4.5 MiB<br />

## Configuration

To configure LibArchive 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" \
./configure --prefix=/usr                            \
            --libdir=/usr/lib                        \
            --libexecdir=/usr/lib                    \
            --docdir=/usr/share/doc/libarchive-3.8.4 \
            --disable-static
```

## Compilation and Installation

To compile LibArchive 32-bit, run the following command:

```bash
make
```

Finally, to install LibArchive 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libarchive.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of LibArchive for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibArchive_64bit.md) LibArchive 64-bit | [HOME](../README.md) | Docbook 3.1 SGML DTDs [>>](./Docbook31-SGML-DTD.md) |

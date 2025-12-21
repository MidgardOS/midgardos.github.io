# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibB2_64bit.md) LibB2 64-bit | [HOME](../README.md) | LibArchive 64-bit [>>](./LibArchive_64bit.md) |

## LibB2 32-bit

Name: LibB2 32-bit<br />
Summary: A C library providing BLAKE2b, BLAKE2s, BLAKE2bp, BLAKE2sp hashes<br />
License: CC0 v1.0<br />
Version: 0.98.1<br />
URL: [https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.tar.gz](https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 52 KiB<br />

## Configuration

To configure LibB2 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" \
./configure --prefix=/usr                        \
            --libdir=/usr/lib                    \
            --libexecdir=/usr/lib                \
            --docdir=/usr/share/doc/libb2-0.98.1 \
            --disable-static                     \
            --disable-native                     \
            --enable-shared
```

## Compilation and Installation

To compile LibB2 32-bit, run the following command:

```bash
make
```

Finally, to install LibB2 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libb2.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of LibB2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibB2_64bit.md) LibB2 64-bit | [HOME](../README.md) | LibArchive 64-bit [>>](./LibArchive_64bit.md) |

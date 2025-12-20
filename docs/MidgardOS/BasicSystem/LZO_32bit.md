# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LZO_64bit.md) LZO 64-bit | [HOME](../README.md) | LibArchive 64-bit [>>](./LibArchive_64bit.md) |

## LZO 32-bit

Name: LZO 32-bit<br />
Summary: A fast real-time compression and decompression library<br />
License: GPL v2 or later<br />
Version: 2.10<br />
URL: [https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.tar.gz](https://www.oberhumer.com/opensource/lzo/download/lzo-2.10.tar.gz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 880 KiB<br />

## Configuration

To configure LZO 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --enable-shared          \
            --disable-static         \
            --docdir=/usr/share/doc/lzo-2.10
```

## Compilation and Installation

To compile LZO 32-bit, run the following command:

```bash
make
```

Finally, to install LZO 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/liblzo2.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of LZO for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LZO_64bit.md) LZO 64-bit | [HOME](../README.md) | LibArchive 64-bit [>>](./LibArchive_64bit.md) |

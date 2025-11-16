# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./File_64bit.md) File 64-bit | [HOME](../README.md) | GNU Readline 64-bit [>>](./GNUReadline_64bit.md) |

## File 32-bit

Name: File 32-bit<br />
Summary: A tool to determine the type of a file from heuristics<br />
License: BSD 2-Clause<br />
Version: 5.46<br />
URL: [http://ftp.astrom.com/pub/file/file-5.46.tar.xz](ftp://ftp.astron.com/pub/file/file-5.46.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 19 MiB<br />

## Configuration

To configure File 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32" ./configure \
    --prefix=/usr         \
    --libdir=/usr/lib     \
    --libexecdir=/usr/lib \
    --host=i686-pc-linux-gnu
```

## Compilation and Installation

To compile File 32-bit, run the following command:

```bash
make
```

Finally, to install File 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib
rm -fv /usr/lib/libmagic.la
```

## Contents

See the contents section of the 64-bit build of File for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./File_64bit.md) File 64-bit | [HOME](../README.md) | GNU Readline 64-bit [>>](./GNUReadline_64bit.md) |

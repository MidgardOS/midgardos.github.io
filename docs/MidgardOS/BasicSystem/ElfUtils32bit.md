# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ElfUtils64bit.md) ElfUtils 64-bit | [HOME](../README.md) | LibFFI 64-bit [>>](./libffi64bit.md) |

## ElfUtils 32-bit

Name: ElfUtils 32-bit<br />
Summary: A library and utilities for inspecting and modifying ELF files<br />
License: GPL v3/GPL v2/LGPL v2.1/GPD<br />
Version: 0.194<br />
URL: [https://sourceware.org/elfutils/ftp/0.194/elfutils-0.194.tar.bz2](https://sourceware.org/elfutils/ftp/0.194/elfutils-0.194.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 1 MiB<br />

## Configuration

To configure ElfUtils 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" CXX="g++ -m32"         \
./configure --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --host=i686-pc-linux-gnu \
            --disable-debuginfod     \
            --enable-libdebuginfod=dummy
```

## Compilation and Installation

To compile ElfUtils 32-bit, run the following commands:

```bash
make -C lib
make -C libelf
```

Finally, to install ElfUtils 32-bit into the build tree, run the following commands:

```bash
make -C libelf DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.a
cp -Rv DESTDIR/usr/lib/* /usr/lib/
install -vm644 config/libelf.pc /usr/lib/pkgconfig
```

## Contents

See the contents section of the 64-bit build of ElfUtils for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./ElfUtils64bit.md) ElfUtils 64-bit | [HOME](../README.md) | LibFFI 64-bit [>>](./libffi64bit.md) |

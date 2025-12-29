# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibPipeline_64bit.md) LibPipeline 64-bit | [HOME](../README.md) | GNU Make [>>](./GNUMake.md) |

## LibPipeline 32-bit

Name: LibPipeline 32-bit<br />
Summary: A C library for constructing command pipelines<br />
License: GPL v3.0 or later<br />
Version: 1.5.8<br />
URL: [https://download.savannah.nongnu.org/releases/%{name}/%{name}-%{version}.tar.gz](https://download.savannah.nongnu.org/releases/%{name}/%{name}-%{version}.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 316 KiB<br />

## Configuration

To configure LibPipeline 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32 -march=i686"                               \
./configure --host=i686-pc-linux-gnu                    \
            --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
            --docdir=/usr/share/doc/libpipeline-1.5.8   \
            --disable-static                            \
            --enable-shared                             \
            --enable-threads=posix                      \
            --enable-socketpair-pipe                    \
            --with-pic=yes                              \
            --with-gnu-ld
```

## Compilation and Installation

To compile LibPipeline 32-bit, run the following command:

```bash
make
```

Finally, to install LibPipeline 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libpipeline.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of LibPipeline for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibPipeline_64bit.md) LibPipeline 64-bit | [HOME](../README.md) | GNU Make [>>](./GNUMake.md) |

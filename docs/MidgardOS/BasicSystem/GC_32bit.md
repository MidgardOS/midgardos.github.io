# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GC_64bit.md) GC 64-bit | [HOME](../README.md) | W3M [>>](./W3m.md) |

## GC 32-bit

Name: GC 32-bit<br />
Summary: A library implementing memory garbage collection for C and C++<br />
License: MIT<br />
Version: 8.2.10<br />
URL: [https://github.com/bdwgc/bdwgc/releases/download/v8.2.10/gc-8.2.10.tar.gz](https://github.com/bdwgc/bdwgc/releases/download/v8.2.10/gc-8.2.10.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.5 MiB<br />

## Configuration

To configure GC 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" CXX="g++ -m32 -march=i686"    \
./configure --host=i686-pc-linux-gnu                    \
            --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
            --docdir=/usr/share/doc/gc-8.2.10           \
            --enable-cplusplus                          \
            --disable-gcj                               \
            --enable-shared                             \
            --disable-static
```

## Compilation and Installation

To compile GC 32-bit, run the following command:

```bash
make
```

Finally, to install GC 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libcord.la
rm -fv DESTDIR/usr/lib/libgc.la
rm -fv DESTDIR/usr/lib/libgccpp.la
rm -fv DESTDIR/usr/lib/libgctba.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GC for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GC_64bit.md) GC 64-bit | [HOME](../README.md) | W3M [>>](./W3m.md) |

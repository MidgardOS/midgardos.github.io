# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGDBM64bit.md) GNU GDBM 64-bit | [HOME](../README.md) | GNU GPerf [>>](./GNUGPerf.md) |

## GNU GDBM 32-bit

Name: GNU GDBM 32-bit<br />
Summary: The GNU Database Manager<br />
License: GPL v3<br />
Version: 1.26<br />
URL: [https://ftp.gnu.org/gnu/gdbm/gdbm-1.26.tar.gz](https://ftp.gnu.org/gnu/gdbm/gdbm-1.26.tar.gz)<br />

Average Build Time: less than 0.2 SBU<br />
Used Install Space: 13 MiB<br />

## Configuration

To configure GNU GDBM 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" CXX="gcc -m32" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --disable-static         \
            --enable-libgdbm-compat
```

## Compilation and Installation

To compile GNU GDBM 32-bit, run the following command:

```bash
make
```

Finally, to install GNU GDBM 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU GDBM for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGDBM64bit.md) GNU GDBM 64-bit | [HOME](../README.md) | GNU GPerf [>>](./GNUGPerf.md) |

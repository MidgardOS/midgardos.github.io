# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibtool64bit.md) GNU Libtool 64-bit | [HOME](../README.md) | GNU Automake [>>](./Automake.md) |

# GNU Libtool 32-bit

Name: GNU Libtool 32-bit<br />
Summary: A tool to build shared libraries<br />
License: GPL v3+<br />
Version: 2.5.4<br />
URL: [https://ftp.gnu.org/pub/gnu/libtool/libtool-2.5.4.tar.xz](https://ftp.gnu.org/pub/gnu/libtool/libtool-2.5.4.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.8 MiB<br />

## Configuration

To configure GNU Libtool 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-midgardos-linux-gnu \
            --prefix=/usr                   \
            --libdir=/usr/lib               \
            --sysconfdir=/etc               \
            --libexecdir=/usr/lib           \
            --disable-static
```

## Compilation and Installation

To compile GNU Libtool 32-bit, run the following command:

```bash
make
```

Finally, to install GNU Libtool 32-bit into the tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libltdl.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU Libtool for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibtool64bit.md) GNU Libtool 64-bit | [HOME](../README.md) | GNU Automake [>>](./Automake.md) |

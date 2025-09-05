# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUReadline64bit.md) GNU Readline 64-bit | [HOME](../README.md) | PCRE2 64-bit [>>](./PCRE264bit) |

## GNU Readline 32-bit

Name: GNU Readline 32-bit<br />
Summary: Libraries for command-line editing and history<br />
License: GPL v3 or later<br />
Version: 8.3<br />
URL: [https://ftp.gnu.org/gnu/readline/readline-8.3.tar.gz](https://ftp.gnu.org/gnu/readline/readline-8.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 17 MiB<br />

## Configuration

To configure GNU Readline 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32" ./configure           \
    --host=i686-midgardos-linux-gnu \
    --prefix=/usr                   \
    --libdir=/usr/lib               \
    --disable-static                \
    --with-curses
```

## Compilation and Installation

To compile GNU Readline 32-bit, run the following command:

```bash
make SHLIB_LIBS="-lncursesw"
```

Finally, to install GNU Readline 32-bit into the build tree, run the following command:

```bash
make SHLIB_LIBS="-lncursesw" DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib
```

## Contents

See the contents section of the 64-bit build of GNU Readline for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUReadline64bit.md) GNU Readline 64-bit | [HOME](../README.md) | PCRE2 64-bit [>>](./PCRE264bit) |

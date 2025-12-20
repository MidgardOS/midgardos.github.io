# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGettext_64bit.md) GNU Gettext 64-bit | [HOME](../README.md) | GNU Bison [>>](./GNUBison.md) |

## GNU Gettext 32-bit

Name: GNU Gettext 32-bit<br />
Summary: Utilities for i18n and l10n tasks<br />
License: GPL v3<br />
Version: 0.26<br />
URL: [https://ftp.gnu.org/gnu/gettext/gettext-0.26.tar.xz](https://ftp.gnu.org/gnu/gettext/gettext-0.26.tar.xz)<br />

Average Build Time: 2.6 SBU<br />
Used Install Space: 395 MiB<br />

## Configuration

To configure GNU Gettext 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" CXX="g++ -m32 -march=i686" \
./configure --host=i686-pc-linux-gnu                 \
            --prefix=/usr                            \
            --libdir=/usr/lib                        \
            --libexecdir=/usr/lib                    \
            --disable-static                         \
            --docdir=/usr/share/doc/gettext-0.26
```

## Compilation and Installation

To compile GNU Gettext 32-bit, run the following command:

```bash
make
```

Finally, to install GNU Gettext 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
chmod -v 0755 DESTDIR/usr/lib/preloadable_libintl.so
rm -fv DESTDIR/usr/lib/libasprintf.la
rm -fv DESTDIR/usr/lib/libgettext*.la
rm -fv DESTDIR/usr/lib/libtextstyle.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU Gettext for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGettext_64bit.md) GNU Gettext 64-bit | [HOME](../README.md) | GNU Bison [>>](./GNUBison.md) |

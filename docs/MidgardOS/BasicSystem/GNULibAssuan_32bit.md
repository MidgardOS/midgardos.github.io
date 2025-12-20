# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibAssuan_64bit.md) GNU LibAssuan 64-bit | [HOME](../README.md) | GNU NPth 64-bit [>>](./GNUNPth_64bit.md) |

## GNU LibAssuan 32-bit

Name: GNU LibAssuan 32-bit<br />
Summary: An IPC protocol library for GnuPG suite<br />
License: GPL v3 or later<br />
Version: 3.0.2<br />
URL: [https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.2.tar.bz2](https://gnupg.org/ftp/gcrypt/libassuan/libassuan-3.0.2.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 652 KiB<br />

## Configuration

To configure GNU LibAssuan 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32 -march=i686"                           \
./configure --host=i686-pc-linux-gnu                \
            --prefix=/usr                           \
            --libdir=/usr/lib                       \
            --libexecdir=/usr/lib                   \
            --docdir=/usr/share/doc/libassuan-3.0.2 \
            --disable-static
```

## Compilation and Installation

To compile GNU LibAssuan 32-bit, run the following command:

```bash
make
```

Finally, to install GNU LibAssuan 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU LibAssuan for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibAssuan_64bit.md) GNU LibAssuan 64-bit | [HOME](../README.md) | GNU NPth 64-bit [>>](./GNUNPth_64bit.md) |

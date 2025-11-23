# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNPth_64bit.md) GNU nPth 64-bit | [HOME](../README.md) | GNU Privacy Guard [>>](./GNUPrivacyGuard.md) |

## GNU nPth 32-bit

Name: GNU nPth 32-bit<br />
Summary: A non-preemptive threading library implementation<br />
License: LGPL v2.1 or later<br />
Version: 1.8<br />
URL: [https://gnupg.org/ftp/gcrypt/npth/npth-1.8.tar.bz2](https://gnupg.org/ftp/gcrypt/npth/npth-1.8.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 96 KiB<br />

## Configuration

To configure GNU nPth 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32 -march=i686"                       \
./configure --host=i686-pc-linux-gnu            \
            --prefix=/usr                       \
            --libdir=/usr/lib                   \
            --libexecdir=/usr/lib               \
            --docdir=/usr/share/doc/npth-1.8    \
            --disable-static
```

## Compilation and Installation

To compile GNU nPth 32-bit, run the following command:

```bash
make
```

Finally, to install GNU nPth 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU nPth for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNPth_64bit.md) GNU nPth 64-bit | [HOME](../README.md) | GNU Privacy Guard [>>](./GNUPrivacyGuard.md) |

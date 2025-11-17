# Section X - $SECTION_TITLE

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibGPG-Error_64bit.md) GNU LibGPG-Error 64-bit | [HOME](../README.md) | GNU LibGCrypt 64-bit [>>](./GNULibGCrypt_64bit.md) |

## GNU LibGPG-Error 32-bit

Name: GNU LibGPG-Error 32-bit<br />
Summary: A library that defines the various error states for GnuPG<br />
License: GPL v2.0 or later/LGPL v2.1 or later<br />
Version: 1.56<br />
URL: [https://gnupg.org/ftp/gcrypt/libgpg-error/libgpg-error-1.56.tar.bz2](https://gnupg.org/ftp/gcrypt/libgpg-error/libgpg-error-1.56.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 1.9 MiB<br />

## Configuration

To configure GNU LibGPG-Error 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686"               \
./configure --host=i686-pc-linux-gnu    \
            --prefix=/usr               \
            --libdir=/usr/lib           \
            --libexecdir=/usr/lib       \
            --disable-static
```

## Compilation and Installation

To compile GNU LibGPG-Error 32-bit, run the following command:

```bash
make
```

Finally, to install GNU LibGPG-Error 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libgpg-error.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of the GNU LibGPG-Error package for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibGPG-Error_64bit.md) GNU LibGPG-Error 64-bit | [HOME](../README.md) | GNU LibGCrypt 64-bit [>>](./GNULibGCrypt_64bit.md) |

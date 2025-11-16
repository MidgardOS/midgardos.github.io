# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXCrypt_64bit.md) LibXCrypt 64-bit | [HOME](../README.md) | LibCap-NG 64-bit [>>](./LibCap-NG_64bit.md) |

## LibXCrypt 32-bit

Name: LibXCrypt 32-bit<br />
Summary: An enhanced encryption library for a number of cryptographic hashing algorithms<br />
License: LGPL v2.1<br />
Version: 4.4.38<br />
URL: [https://github.com/besser82/libxcrypt/releases/download/v4.4.38/libxcrypt-4.4.38.tar.xz](https://github.com/besser82/libxcrypt/releases/download/v4.4.38/libxcrypt-4.4.38.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 12 MiB<br />

## Configuration

To configure LibXCrypt 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu    \
            --prefix=/usr               \
            --libdir=/usr/lib           \
            --libexecdir=/usr/lib       \
            --enable-hashes=all         \
            --enable-obsolete-api=suse  \
            --disable-static            \
            --disable-failure-tokens
```


## Compilation and Installation

To compile LibXCrypt 32-bit, run the following command:

```bash
make
```

Finally, to install LibXCrypt 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libcrypt.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of LibXCrypt for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXCrypt_64bit.md) LibXCrypt 64-bit | [HOME](../README.md) | LibCap-NG 64-bit [>>](./LibCap-NG_64bit.md) |

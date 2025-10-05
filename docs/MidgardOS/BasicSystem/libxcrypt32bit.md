# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libxcrypt64bit.md) libxcrypt 64-bit | [HOME](../README.md) | libcap-ng 64-bit [>>](./libcap-ng64bit.md) |

## libxcrypt 32-bit

Name: libxcrypt 32-bit<br />
Summary: An enhanced encryption library for a number of cryptographic hashing algorithms<br />
License: LGPL v2.1<br />
Version: 4.4.38<br />
URL: [https://github.com/besser82/libxcrypt/releases/download/v4.4.38/libxcrypt-4.4.38.tar.xz](https://github.com/besser82/libxcrypt/releases/download/v4.4.38/libxcrypt-4.4.38.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 12 MiB<br />

## Configuration

To configure libxcrypt 32-bit for install into the build root, run the following commands:

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

To compile libxcrypt 32-bit, run the following command:

```bash
make
```

Finally, to install libxcrypt 32-bit into the build tree, run the following commands:

```bash
cp -av .libs/libcrypt.so* /usr/lib/ &&
make install-pkgconfigDATA &&
ln -svf libxcrypt.pc /usr/lib/pkgconfig/libcrypt.pc
```

## Contents

See the contents section of the 64-bit build of libxcrypt for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libxcrypt64bit.md) libxcrypt 64-bit | [HOME](../README.md) | libcap-ng 64-bit [>>](./libcap-ng64bit.md) |

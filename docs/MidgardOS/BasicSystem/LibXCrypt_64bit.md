# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibCap_32bit.md) LibCap2 32-bit | [HOME](../README.md) | LibXCrypt 32-bit [>>](./LibXCrypt_32bit.md) |

## LibXCrypt 64-bit

Name: LibXCrypt 64-bit<br />
Summary: An enhanced encryption library for a number of cryptographic hashing algorithms<br />
License: LGPL v2.1<br />
Version: 4.4.38<br />
URL: [https://github.com/besser82/libxcrypt/releases/download/v4.4.38/libxcrypt-4.4.38.tar.xz](https://github.com/besser82/libxcrypt/releases/download/v4.4.38/libxcrypt-4.4.38.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 12 MiB<br />

## Configuration

To configure LibXCrypt 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr               \
            --libdir=/usr/lib64         \
            --libexecdir=/usr/lib64     \
            --enable-hashes=all         \
            --enable-obsolete-api=suse  \
            --disable-static            \
            --disable-failure-tokens
```

## Compilation and Installation

To compile LibXCrypt 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install LibXCrypt 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libcrypt.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libcrypt.so and libxcrypt.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibCap_32bit.md) LibCap2 32-bit | [HOME](../README.md) | LibXCrypt 32-bit [>>](./LibXCrypt_32bit.md) |

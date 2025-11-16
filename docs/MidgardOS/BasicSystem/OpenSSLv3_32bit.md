# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./OpenSSLv3_64bit.md) OpenSSL v3 64-bit | [HOME](../README.md) | ElfUtils 64-bit [>>](./ElfUtils_64bit.md) |

## OpenSSL v3 32-bit

Name: OpenSSL v3 32-bit<br />
Summary: The standard open source cryptography library and tools<br />
License: Apache v2.0<br />
Version: 3.6.0<br />
URL: [https://github.com/openssl/openssl/releases/download/openssl-3.6.0/openssl-3.6.0.tar.gz](https://github.com/openssl/openssl/releases/download/openssl-3.6.0/openssl-3.6.0.tar.gz)<br />

Average Build Time: 1.9 SBU<br />
Used Install Space: 35 MiB<br />

## Configuration

To configure OpenSSL v3 32-bit for install into the build root, run the following commands:

```bash
make distclean
./config --prefix=/usr               \
         --libdir=/usr/lib           \
         --openssldir=/etc/ssl       \
         shared                      \
         zlib-dynamic                \
         enable-camellia             \
         enable-fips                 \
         enable-ktls                 \
         enable-pie                  \
         enable-rfc3779              \
         enable-seed                 \
         no-afalgeng                 \
         no-atexit                   \
         no-ec2m                     \
         no-mdc2                     \
         linux-x86
perl configdata.pm --dump
```

## Compilation and Installation

To compile OpenSSL v3 32-bit, run the following command:

```bash
make
```

Finally, to install OpenSSL v3 32-bit into the build tree, run the following commands:

```bash
sed -i '/INSTALL_LIBS/s/libcrypto.a libssl.a//' Makefile
make MANSUFFIX=ssl DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of OpenSSL v3 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./OpenSSLv3_64bit.md) OpenSSL v3 64-bit | [HOME](../README.md) | ElfUtils 64-bit [>>](./ElfUtils_64bit.md) |

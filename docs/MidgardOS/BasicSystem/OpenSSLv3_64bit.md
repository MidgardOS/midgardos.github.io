# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./intltool.md) IntlTool | [HOME](../README.md) | OpenSSL v3 32-bit [>>](./OpenSSLv3_32bit.md) |

## OpenSSL v3 64-bit

Name: OpenSSL v3 64-bit<br />
Summary: The standard open source cryptography library and tools<br />
License: Apache v2.0<br />
Version: 3.6.0<br />
URL: [https://github.com/openssl/openssl/releases/download/openssl-3.6.0/openssl-3.6.0.tar.gz](https://github.com/openssl/openssl/releases/download/openssl-3.6.0/openssl-3.6.0.tar.gz)<br />

Average Build Time: 1.9 SBU<br />
Used Install Space: 35 MiB<br />

## Configuration

To configure OpenSSL v3 64-bit for install into the build root, run the following commands:

```bash
./config --prefix=/usr               \
         --libdir=/usr/lib64         \
         --openssldir=/etc/ssl       \
         shared                      \
         zlib-dynamic                \
         enable-camellia             \
         enable-ec_nistp_64_gcc_128  \
         enable-fips                 \
         enable-ktls                 \
         enable-pie                  \
         enable-rfc3779              \
         enable-seed                 \
         no-afalgeng                 \
         no-atexit                   \
         no-ec2m                     \
         no-mdc2
perl configdata.pm --dump
```

## Compilation and Installation

To compile OpenSSL v3 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
HARNESS_JOBS=$(nproc) make test
```

Finally, to install OpenSSL v3 64-bit into the build tree, run the following commands:

```bash
sed -i '/INSTALL_LIBS/s/libcrypto.a libssl.a//' Makefile
make MANSUFFIX=ssl install
mv -v /usr/share/doc/openssl /usr/share/doc/openssl-3.6.0
cp -vfr doc/* /usr/share/doc/openssl-3.6.0
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | c_rehash and openssl |
| Installed Libraries | libcrypto.so and libssl.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./intltool.md) IntlTool | [HOME](../README.md) | OpenSSL v3 32-bit [>>](./OpenSSLv3_32bit.md) |

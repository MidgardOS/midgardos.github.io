# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibTasn1_32bit.md) GNU LibTasn1 32-bit | [HOME](../README.md) | GNU TLS 32-bit [>>](./GNUTLS_32bit.md) |

## GNU TLS 64-bit

Name: GNU TLS 64-bit<br />
Summary: The GNU TLS and SSL library<br />
License: GPL v3 or later/LGPL v2.1 or later<br />
Version: 3.8.11<br />
URL: [https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.11.tar.xz](https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.11.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 20 MiB<br />

## Configuration

To configure GNU TLS 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
            --disable-static                            \
            --disable-rpath                             \
            --enable-shared                             \
            --disable-gcc-warnings                      \
            --enable-fips140-mode                       \
            --with-fips140-module-name="GnuTLS version" \
            --with-fips140-module-version="3.8.11"      \
            --enable-year2038                           \
            --without-p11-kit                           \
            --without-tpm                               \
            --disable-libdane                           \
            --disable-doc                               \
            --enable-manpages                           \
            gl_cv_func_printf_directive_n=yes           \
            gl_cv_func_printf_infinite_long_double=yes
```

**NOTE: The following features are not enabled due to missing dependences:<br />**
**- PKCS11 support<br />**
**- TPM and TPM2 support<br />**
**- KTLS support<br />**
**- Brotli compression support<br />**
**- DANE library support<br />**
**- DNSSEC support<br />**
**- TeX documentation is not built<br />**

The above features will be enabled when being rebuilt as an RPM.

## Compilation and Installation

To compile GNU TLS 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

NOTE: Two of the tests fail inside the build root. This is likely due to missing dependencies for DNSSEC and other features.

Finally, to install GNU TLS 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libgnutls.la && rm -fv /usr/lib64/libgnutlsxx.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | certtool, gnutls-cli, gnutls-cli-debug, gnutls-serv, ocsptool, and psktool |
| Installed Libraries | libgnutls.so and libgnutlsxx.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibTasn1_32bit.md) GNU LibTasn1 32-bit | [HOME](../README.md) | GNU TLS 32-bit [>>](./GNUTLS_32bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTLS_64bit.md) GNU TLS 64-bit | [HOME](../README.md) | GNU Privacy Guard [>>](./GNUPrivacyGuard.md) |

## GNU TLS 32-bit

Name: GNU TLS 32-bit<br />
Summary: The GNU TLS and SSL library<br />
License: GPL v3 or later/LGPL v2.1 or later<br />
Version: 3.8.11<br />
URL: [https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.11.tar.xz](https://www.gnupg.org/ftp/gcrypt/gnutls/v3.8/gnutls-3.8.11.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 20 MiB<br />

## Configuration

To configure GNU TLS 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32 -march=i686" CXX="g++ -m32 -march=i686"    \
./configure --host=i686-pc-linux-gnu                    \
            --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
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

## Compilation and Installation

To compile GNU TLS 32-bit, run the following command:

```bash
make
```

Finally, to install GNU TLS 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
for la_file in "libgnutls.la" "libgnutlsxx.la"; do
    rm -fv DESTDIR/usr/lib/$la_file
done
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU TLS for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTLS_64bit.md) GNU TLS 64-bit | [HOME](../README.md) | GNU Privacy Guard [>>](./GNUPrivacyGuard.md) |

# Section X - $SECTION_TITLE

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSS_64bit-pass1.md) Mozilla NSS 64-bit - pass 1 | [HOME](../README.md) | Python module LXML [>>](./Python313-LXML.md) |

## Mozilla NSS 32-bit - pass 1

Name: Mozilla NSS 32-bit<br />
Summary: The Netscape Network Security Services<br />
License: MPL v2.0<br />
Version: 3.121<br />
URL: [https://archive.mozilla.org/pub/security/nss/releases/NSS_3_121_RTM/src/nss-3.121.tar.gz](https://archive.mozilla.org/pub/security/nss/releases/NSS_3_121_RTM/src/nss-3_121.tar.gz)<br />

Average Build Time: 0.8 SBU<br />
Used Install Space: 28 KiB<br />

## Configuration

This package does not come with a traditional configuration script.

## Compilation and Installation

To compile Mozilla NSS 32-bit, run the following commands:

```bash
cd nss/
make clean
rm -rf ../DESTDIR
rm -rf ../dist
rm -rf ../tests_results
make -j1 BUILD_OPT=1                            \
  LIBDIR=/usr/lib                               \
  NSPR_INCLUDE_DIR=$(nspr-config --includedir)  \
  NSPR_LIB_DIR=/usr/lib                         \
  USE_SYSTEM_SQLITE=1                           \
  USE_SYSTEM_ZLIB=1                             \
  ZLIB_LIBS=-lz                                 \
  NSS_ENABLE_WERROR=0                           \
  NSS_USE_SYSTEM_SQLITE=1
```

Finally, to install Mozilla NSS 32-bit into the build tree, run the following commands:

```bash
cd ../dist
install -v -d -m755 -o root -g root DESTDIR
install -v -d -m755 -o root -g root DESTDIR/usr
install -v -d -m755 -o root -g root DESTDIR/usr/lib
install -v -d -m755 -o root -g root DESTDIR/usr/lib/pkgconfig
install -v -m755 Linux*/lib/*.so                           DESTDIR/usr/lib/
install -v -m644 Linux*/lib/{*.chk,libcrmf.a}              DESTDIR/usr/lib/
install -v -m644 Linux*/lib/pkgconfig/nss.pc               DESTDIR/usr/lib/pkgconfig/
unalias cp
cp -Rv DESTDIR/usr/lib/*          /usr/lib/
alias cp="cp -i"
cd -
```

## Contents

See the contents section of the 64-bit build of Mozilla NSS for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSS_64bit-pass1.md) Mozilla NSS 64-bit - pass 1 | [HOME](../README.md) | Python module LXML [>>](./Python313-LXML.md) |

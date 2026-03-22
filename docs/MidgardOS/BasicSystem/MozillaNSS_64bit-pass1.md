# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSPR_32bit.md) Mozilla NSPR 32-bit | [HOME](../README.md) | Mozilla NSS 32-bit [>>](./MozillaNSS_32bit-pass1.md) |

## Mozilla NSS 64-bit - Pass 1

Name: Mozilla NSS 64-bit<br />
Summary: The Netscape Network Security Services<br />
License: MPL v2.0<br />
Version: 3.121<br />
URL: [https://archive.mozilla.org/pub/security/nss/releases/NSS_3_121_RTM/src/nss-3.121.tar.gz](https://archive.mozilla.org/pub/security/nss/releases/NSS_3_121_RTM/src/nss-3_121.tar.gz)<br />

Average Build Time: 0.8 SBU<br />
Used Install Space: 28 KiB<br />

## Configuration

This package does not come with a traditional configuration script.

## Compilation and Installation

To compile Mozilla NSS 64-bit, run the following commands:

```bash
patch -Np1 -i ../patches/nss/nss-standalone-1.patch
cd nss
make -j1 BUILD_OPT=1                            \
  LIBDIR=/usr/lib64                             \
  NSPR_INCLUDE_DIR=$(nspr-config --includedir)  \
  NSPR_LIB_DIR=$(nspr-config --libdir)          \
  USE_SYSTEM_SQLITE=1                           \
  USE_SYSTEM_ZLIB=1                             \
  ZLIB_LIBS=-lz                                 \
  NSS_ENABLE_WERROR=0                           \
  NSS_USE_SYSTEM_SQLITE=1                       \
  $([ $(uname -m) = x86_64 ] && echo USE_64=1)
```

The tests cannot be run right now due to a circular dependency on p11-kit and make-ca. They will be run later after those packages are added.

Finally, to install Mozilla NSS 64-bit into the build tree, run the following command:

```bash
cd ../dist
install -v -d -m755 -o root -g root DESTDIR
install -v -d -m755 -o root -g root DESTDIR/usr
install -v -d -m755 -o root -g root DESTDIR/usr/{bin,include,lib64}
install -v -d -m755 -o root -g root DESTDIR/usr/include/nss
install -v -d -m755 -o root -g root DESTDIR/usr/lib64/pkgconfig
install -v -m755 Linux*/lib/*.so                           DESTDIR/usr/lib64
install -v -m644 Linux*/lib/{*.chk,libcrmf.a}              DESTDIR/usr/lib64
cp -v -RL {public,private}/nss/*                           DESTDIR/usr/include/nss
install -v -m755 Linux*/bin/{certutil,nss-config,pk12util} DESTDIR/usr/bin
install -v -m644 Linux*/lib/pkgconfig/nss.pc               DESTDIR/usr/lib64/pkgconfig
unalias cp
cp -Rv DESTDIR/usr/bin/*            /usr/bin/
install -d -v -m755 -o root -g root /usr/install/nss
cp -Rv DESTDIR/usr/include/nss/*    /usr/include/nss/
cp -Rv DESTDIR/usr/lib64/*          /usr/lib64/
alias cp="cp -i"
cd -
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | certutil, nss-config, and pk12util |
| Installed Libraries | libcrmf.a, libfreebl3.so, libfreeblpriv3.so, libnss3.so, libnssckbi.so, libnssckbi-testlib.so, libnssdbm3.so, libnsssysinit.so, libnssutil3.so, libpkcs11testmodule.so, libsmime3.so, libsoftokn3.so, and libssl3.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSPR_32bit.md) Mozilla NSPR 32-bit | [HOME](../README.md) | Mozilla NSS 32-bit - pass 1 [>>](./MozillaNSS_32bit-pass1.md) |

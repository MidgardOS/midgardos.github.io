# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSecComp_32bit.md) LibSecComp 32-bit | [HOME](../README.md) | Mozilla NSPR 32-bit [>>](./MozillaNSPR_32bit.md) |

## Mozilla NSPR 64-bit

Name: Mozilla NSPR 64-bit<br />
Summary: The Netscape Portable Runtime<br />
License: MPL v2<br />
Version: 4.38.2<br />
URL: [https://archive.mozilla.org/pub/nspr/releases/v4.38.2/src/nspr-4.38.2.tar.gz](https://archive.mozilla.org/pub/nspr/releases/v4.38.2/src/nspr-4.38.2.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.6 MiB<br />

## Configuration

To configure Mozilla NSPR 64-bit for install into the build root, run the following command:

```bash
sed -i '/^RELEASE/s|^|#|' pr/src/misc/Makefile.in &&
sed -i 's|$(LIBRARY) ||'  config/rules.mk
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/nspr-4.38.2 \
            --sysconfdir=/etc                   \
            --enable-ipv6                       \
            --with-mozilla                      \
            --with-pthreads                     \
            $([ $(uname -m) = x86_64 ] && echo --enable-64bit)
```

## Compilation and Installation

To compile Mozilla NSPR 64-bit, run the following command:

```bash
make
```

The test suite is not operational without Mozilla NSS side-by-side.

Finally, to install Mozilla NSPR 64-bit into the build tree, run the following command:

```bash
make install
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | nspr-config |
| Installed Libraries | libnspr4.so, libplc4.so, and libplds4.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSecComp_32bit.md) LibSecComp 32-bit | [HOME](../README.md) | Mozilla NSPR 32-bit [>>](./MozillaNSPR_32bit.md) |

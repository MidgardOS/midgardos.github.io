# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSPR_64bit.md) Mozilla NSPR 64-bit | [HOME](../README.md) | Mozilla NSS 64-bit [>>](./MozillaNSS_64bit.md) |

## Mozilla NSPR 32-bit

Name: Mozilla NSPR 32-bit<br />
Summary: The Netscape Portable Runtime<br />
License: MPL v2<br />
Version: 4.38.2<br />
URL: [https://archive.mozilla.org/pub/nspr/releases/v4.38.2/src/nspr-4.38.2.tar.gz](https://archive.mozilla.org/pub/nspr/releases/v4.38.2/src/nspr-4.38.2.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.6 MiB<br />

## Configuration

To configure Mozilla NSPR 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32 -march=i686" \
    CXX="g++ -m32 -march=i686" \
    ./configure --host=i686-pc-linux-gnu        \
            --prefix=/usr                       \
            --libdir=/usr/lib                   \
            --libexecdir=/usr/lib               \
            --docdir=/usr/share/doc/nspr-4.38.2 \
            --sysconfdir=/etc                   \
            --enable-ipv6                       \
            --with-mozilla                      \
            --with-pthreads
```

## Compilation and Installation

To compile Mozilla NSPR 32-bit, run the following command:

```bash
make
```

Finally, to install Mozilla NSPR 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of Mozilla NSPR for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSPR_64bit.md) Mozilla NSPR 64-bit | [HOME](../README.md) | Mozilla NSS 64-bit [>>](./MozillaNSS_64bit.md) |

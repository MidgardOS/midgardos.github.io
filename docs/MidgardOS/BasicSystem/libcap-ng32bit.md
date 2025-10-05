# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libcap-ng64bit.md) libcap-ng 64-bit | [HOME](../README.md) | libtirpc 64-bit [>>](./libtirpc64bit.md) |

## libcap-ng 32-bit

Name: libcap-ng 32-bit<br />
Summary: An alternate POSIX capabilities library<br />
License: GPL v2/LGPL v2.1<br />
Version: 0.8.5<br />
URL: [https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.tar.gz](https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 348 KiB<br />

## Configuration

To configure libcap-ng 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --disable-static
```

## Compilation and Installation

To compile libcap-ng 32-bit, run the following command:

```bash
make
```

Finally, to install libcap2 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of libcap-ng for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libcap-ng64bit.md) libcap-ng 64-bit | [HOME](../README.md) | libtirpc 64-bit [>>](./libtirpc64bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libnsl264bit.md) libnsl2 64-bit | [HOME](../README.md) | tcp_wrappers 64-bit [>>](./tcp_wrappers64bit.md) |

## libnsl2 32-bit

Name: libnsl2 32-bit<br />
Summary: The NIS/YP API library<br />
License: LGPL v2.1<br />
Version: 2.0.1<br />
URL: [https://github.com/thkukuk/libnsl/releases/download/v2.0.1/libnsl-2.0.1.tar.xz](https://github.com/thkukuk/libnsl/releases/download/v2.0.1/libnsl-2.0.1.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 168 KiB<br />

## Configuration

To configure the libnsl2 32-bit package for compilation, run the following command:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-midgardos-linux-gnu \
            --prefix=/usr                   \
            --libdir=/usr/lib               \
            --libexecdir=/usr/lib           \
            --sysconfdir=/etc               \
            --disable-static
```

## Compilation and Installation

To compile libnsl2 32-bit, run the following command:

```bash
make
```

Finally, to install libnsl2 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
cp -Rv DESTDIR/usr/lib/* /usr/lib
```

## Contents

See the contents section of the 64-bit build of libnsl2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libnsl264bit.md) libnsl2 64-bit | [HOME](../README.md) | tcp_wrappers 64-bit [>>](./tcp_wrappers64bit.md) |

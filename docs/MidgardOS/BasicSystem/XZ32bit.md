# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./XZ64bit.md) XZ 64-bit | [HOME](../README.md) | Lz4 64-bit [>>](./Lz464.md) |

## XZ 32-bit

Name: XZ 32-bit<br />
Summary: A Better Compression Tool<br />
License: Multiple FOSS licenses (0BSD/GPLv2+/GPLv3+)<br />
Version: 5.8.1<br />
URL: [https://github.com/tukaani-project/xz/releases/download/v5.8.1/xz-5.8.1.tar.xz](https://github.com/tukaani-project/xz/releases/download/v5.8.1/xz-5.8.1.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 24 MiB<br />

## Configuration

To configure XZ 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" ./configure \
    --host=i686-midgardos-linux-gnu \
    --prefix=/usr                   \
    --libdir=/usr/lib               \
    --libexecdir=/usr/lib           \
    --disable-static
```

## Compilation and Installation

To compile XZ 32-bit, run the following command:

```bash
make
```

Finally, to install XZ 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib
rm -rf DESTDIR
```

## Contents

See the contents section of the 64-bit build of XZ for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./XZ64bit.md) XZ 64-bit | [HOME](../README.md) | Lz4 64-bit [>>](./Lz464.md) |

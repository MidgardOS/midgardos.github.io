# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libffi64bit.md) LibFFI 64-bit | [HOME](../README.md) | 7zip CLI [>>](./7zip.md) |

## LibFFI 32-bit

Name: LibFFI 32-bit<br />
Summary: A foreign function interface library<br />
License: MIT<br />
Version: 3.5.2<br />
URL: [https://github.com/libffi/libffi/releases/download/v3.5.2/libffi-3.5.2.tar.gz](https://github.com/libffi/libffi/releases/download/v3.5.2/libffi-3.5.2.tar.gz)<br />

Average Build Time: 1.7 SBU<br />
Used Install Space: 168 KiB<br />

## Configuration

To configure LibFFI 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" CXX="g++ -m32"         \
./configure --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --host=i686-pc-linux-gnu \
            --disable-static         \
            --with-gcc-arch=i686
```

## Compilation and Installation

To compile LibFFI 32-bit, run the following command:

```bash
make
```

Finally, to install LibFFI 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libffi.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of the LibFFI 32-bit for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libffi64bit.md) LibFFI 64-bit | [HOME](../README.md) | 7zip CLI [>>](./7zip.md) |

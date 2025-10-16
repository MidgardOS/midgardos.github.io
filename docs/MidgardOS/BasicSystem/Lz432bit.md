# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Lz464bit.md) Lz4 64-bit | [HOME](../README.md) | ZStd 64-bit [>>](./ZStd64bit.md) |

## Lz4 32-bit

Name: Lz4 32-bit<br />
Summary: A lossless compression library<br />
License: BSD-2-Clause/GPLv2 or later<br />
Version: 1.10.0<br />
URL: [https://github.com/lz4/lz4/releases/download/v1.10.0/lz4-1.10.0.tar.gz](https://github.com/lz4/lz4/releases/download/v1.10.0/lz4-1.10.0.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 4.2 MiB<br />

## Configuration

The sources for Lz4 do not have a traditional configuration script.

## Compilation and Installation

To compile Lz4 32-bit, run the following commands:

```bash
make clean
CC="gcc -m32" make -j1 BUILD_STATIC=no
```

Finally, to install Lz4 32-bit into the build tree, run the following command:

```bash
make -j1 BUILD_STATIC=no PREFIX=/usr LIBDIR=/usr/lib DESTDIR=$(pwd)/m32 install && \
cp -a m32/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of Lz4 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./Lz464bit.md) Lz4 64-bit | [HOME](../README.md) | ZStd 64-bit [>>](./ZStd64bit.md) |

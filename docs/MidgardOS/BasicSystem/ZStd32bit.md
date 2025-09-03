# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ZStd64bit.md) ZStd 64-bit | [HOME](../README.md) | GNU libunistring 64-bit [>>](./libunistring64bit.md) |

## ZStd 32-bit

Name: ZStd 32-bit<br />
Summary: Real-time compression library<br />
License: BSD-3-Clause AND GPL-2.0-only<br />
Version: 1.5.7<br />
URL: [https://github.com/facebook/zstd/releases/download/v1.5.7/zstd-1.5.7.tar.gz](https://github.com/facebook/zstd/releases/download/v1.5.7/zstd-1.5.7.tar.gz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 86 MiB<br />

## Configuration

The ZStd sources don't have a standard configuration script.

## Compilation and Installation

To compile ZStd 32-bit, run the following command:

```bash
make clean
CC="gcc -m32" make prefix=/usr libdir=/usr/lib
```

Finally, to install ZStd 32-bit into the build tree, run the following command:

```bash
make prefix=/usr DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of ZStd for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./ZStd64bit.md) ZStd 64-bit | [HOME](../README.md) | GNU libunistring 64-bit [>>](./libunistring64bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./BZip264bit.md) BZip2 64-bit | [HOME](../README.md) | XZ 64-bit [>>](./XZ64bit.md) |

## BZip2 32-bit

Name: BZip2 32-bit<br />
Summary: An efficient file compression tool<br />
License: BSD-3-Clause<br />
Version: 1.0.8<br />
URL: [https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz](https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 7.3 MiB<br />

## Preparation

First, clean the previously extracted source tree to allow the 32-bit build to work:

```bash
make clean
```

## Compilation and Installation

After preparing the BZip2 sources for build, run the following commands to generate a dynamic 32-bit library:

```bash
sed -e "s/^CC=.*/CC=gcc -m32/" -i Makefile{,-libbz2_so}
make -f Makefile-libbz2_so
make libbz2.a
```

Finally, to install BZip2 64-bit into the build tree, run the following commands:

```bash
install -Dm755 libbz2.so.1.0.8 /usr/lib/libbz2.so.1.0.8
ln -sf libbz2.so.1.0.8 /usr/lib/libbz2.so
ln -sf libbz2.so.1.0.8 /usr/lib/libbz2.so.1
ln -sf libbz2.so.1.0.8 /usr/lib/libbz2.so.1.0
```

## Contents

See the contents section of the 64-bit build of BZip2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./BZip264bit.md) BZip2 64-bit | [HOME](../README.md) | XZ 64-bit [>>](./XZ64bit.md) |

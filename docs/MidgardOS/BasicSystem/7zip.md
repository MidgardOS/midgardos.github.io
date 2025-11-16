# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibFFI_32bit.md) LibFFI 32-bit | [HOME](../README.md) | SQLite3 64-bit [>>](./SQLite3_64bit.md) |

## 7zip

Name: 7zip<br />
Summary: A command line tool to manage archive files of various formats<br />
License: BSD-3-clause/LGPL v2.1<br />
Version: 25.01<br />
URL: [https://www.7-zip.org/a/7z2501-src.tar.xz](https://www.7-zip.org/a/7z2501-src.tar.xz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 2.8 MiB<br />

## Preparation

This package's archive does not have a parent directory for the archive. As such, it is recommended to unpack it into a directory inside the `/sources` directory:

```bash
mkdir -pv /sources/7z2501
cp -v 7z2501-src.tar.xz /sources/7z2501/
pushd /sources/7z2501
    tar xvf 7z2501-src.tar.xz
popd
```

There is no standard configuration script for this package.

## Compilation and Installation

To compile 7zip, run the following commands:

```bash
pushd /sources/7z2501/CPP/7zip/Bundle/Alone2
    make -j -f makefile.gcc
popd
```

The 7zip package does not include a test harness that runs on Linux systems.

Finally, to install 7zip's CLI tool into the build tree, run the following commands:

```bash
pushd /sources/7z2501/CPP/7zip/Bundle/Alone2/_o
    install -v -m 755 -o root -g root ./7zz /usr/bin/7zz
popd
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | 7zz |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibFFI_32bit.md) LibFFI 32-bit | [HOME](../README.md) | SQLite3 64-bit [>>](./SQLite3_64bit.md) |

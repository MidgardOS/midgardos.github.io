# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Kbd.md) Kbd | [HOME](../README.md) | LibPipeline 32-bit [>>](./LibPipeline_32bit.md) |

## LibPipeline 64-bit

Name: LibPipeline 64-bit<br />
Summary: A C library for constructing command pipelines<br />
License: GPL v3.0 or later<br />
Version: 1.5.8<br />
URL: [https://download.savannah.nongnu.org/releases/%{name}/%{name}-%{version}.tar.gz](https://download.savannah.nongnu.org/releases/%{name}/%{name}-%{version}.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 316 KiB<br />

## Configuration

To configure LibPipeline 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
            --docdir=/usr/share/doc/libpipeline-1.5.8   \
            --disable-static                            \
            --enable-shared                             \
            --enable-threads=posix                      \
            --enable-socketpair-pipe                    \
            --with-pic=yes                              \
            --with-gnu-ld
```

## Compilation and Installation

To compile LibPipeline 64-bit, run the following command:

```bash
make
```

The tests require the Check library, which will be built later.

Finally, to install LibPipeline 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libpipeline.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libpipeline.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./Kbd.md) Kbd | [HOME](../README.md) | LibPipeline 32-bit [>>](./LibPipeline_32bit.md) |

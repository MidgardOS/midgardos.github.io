# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./OpenSSLv3_32bit.md) OpenSSL v3 32-bit | [HOME](../README.md) | ElfUtils 32-bit [>>](./ElfUtils_32bit.md) |

## ElfUtils 64-bit

Name: ElfUtils 64-bit<br />
Summary: A library and utilities for inspecting and modifying ELF files<br />
License: GPL v3/GPL v2/LGPL v2.1/GPD<br />
Version: 0.194<br />
URL: [https://sourceware.org/elfutils/ftp/0.194/elfutils-0.194.tar.bz2](https://sourceware.org/elfutils/ftp/0.194/elfutils-0.194.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 1 MiB<br />

## Configuration

To configure ElfUtils 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --enable-libdebuginfod=dumm
```

## Compilation and Installation

To compile ElfUtils 64-bit, run the following commands:

```bash
make -C lib
make -C libelf
```

Next, run the test suite:

```bash
make -k check
```

Finally, to install ElfUtils 64-bit into the build tree, run the following command:

```bash
make -C libelf DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib64/*.a
cp -Rv DESTDIR/usr/include/* /usr/include/
cp -Rv DESTDIR/usr/lib64/* /usr/lib64/
install -vm644 config/libelf.pc /usr/lib64/pkgconfig
rm -rf DESTDIR
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libelf.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./OpenSSLv3_32bit.md) OpenSSL v3 32-bit | [HOME](../README.md) | ElfUtils 32-bit [>>](./ElfUtils_32bit.md) |

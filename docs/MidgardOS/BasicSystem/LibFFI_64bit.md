# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ElfUtils_32bit.md) ElfUtils 32-bit | [HOME](../README.md) | LibFFI 32-bit [>>](./LibFFI_32bit.md) |

## LibFFI 64-bit

Name: LibFFI 64-bit<br />
Summary: A foreign function interface library<br />
License: MIT<br />
Version: 3.5.2<br />
URL: [https://github.com/libffi/libffi/releases/download/v3.5.2/libffi-3.5.2.tar.gz](https://github.com/libffi/libffi/releases/download/v3.5.2/libffi-3.5.2.tar.gz)<br />

Average Build Time: 1.7 SBU<br />
Used Install Space: 168 KiB<br />

## Configuration

To configure LibFFI 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --with-gcc-arch=x86-64-v2
```

## Compilation and Installation

To compile LibFFI 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install LibFFI 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libffi.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libffi.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./ElfUtils_32bit.md) ElfUtils 32-bit | [HOME](../README.md) | LibFFI 32-bit [>>](./LibFFI_32bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libtirpc32bit.md) libtirpc 32-bit | [HOME](../README.md) | libnsl2 32-bit [>>](./libnsl232bit.md) |

## libnsl2 64-bit

Name: libnsl2 64-bit<br />
Summary: The NIS/YP API library<br />
License: LGPL v2.1<br />
Version: 2.0.1<br />
URL: [https://github.com/thkukuk/libnsl/releases/download/v2.0.1/libnsl-2.0.1.tar.xz](https://github.com/thkukuk/libnsl/releases/download/v2.0.1/libnsl-2.0.1.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 168 KiB<br />

## Configuration

To configure the libnsl2 64-bit package for compilation, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --sysconfdir=/etc       \
            --disable-static
```

## Compilation and Installation

To compile libnsl2 64-bit, run the following commands:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install libnsl2 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libnsl.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libnsl.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./libtirpc32bit.md) libtirpc 32-bit | [HOME](../README.md) | libnsl2 32-bit [>>](./libnsl232bit.md) |

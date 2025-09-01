# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GLibC32bit.md) GNU GLibC 32-bit | [HOME](../README.md) | ZLib 32-bit [>>](./ZLib32bit.md) |

# ZLib 64-bit

Name: zlib 64-bit<br />
Summary: A fast compression library<br />
License: Zlib<br />
Version: 1.3.1<br />
URL: [https://www.zlib.net/zlib-1.3.1.tar.xz](https://www.zlib.net/zlib-1.3.1.tar.xz)<br />

Average Build Time: 0.1 SBU
Used Install Space: 6.4 MiB

## Configuration

To configure ZLib 64-bit for installation into the system, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --shared
```

## Compilation and Installation

To compile ZLib 64-bit, run the following command:

```bash
make
```

Next, run the tests for the package:

```bash
make check
```

Finally, to install ZLib 64-bit, run the following command:

```bash
make install
rm -fv /usr/lib64/libz.a
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Libraries | libz.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GLibC32bit.md) GNU GLibC 32-bit | [HOME](../README.md) | ZLib 32-bit [>>](./ZLib32bit.md) |

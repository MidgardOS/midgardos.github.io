# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC_32bit.md) GNU GLibC 32-bit | [HOME](../README.md) | ZLib 32-bit [>>](./ZLib_32bit.md) |

# ZLib 64-bit

Name: zlib 64-bit<br />
Summary: A fast compression library<br />
License: Zlib<br />
Version: 1.3.1<br />
URL: [https://www.zlib.net/zlib-1.3.1.tar.xz](https://www.zlib.net/zlib-1.3.1.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 6.4 MiB<br />

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
| [<<](./GNUGLibC_32bit.md) GNU GLibC 32-bit | [HOME](../README.md) | ZLib 32-bit [>>](./ZLib_32bit.md) |

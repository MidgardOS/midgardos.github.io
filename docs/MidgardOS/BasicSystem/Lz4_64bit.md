# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./XZ_32bit.md) XZ 32-bit | [HOME](../README.md) | Lz4 32-bit [>>](./Lz4_32bit.md) |

## Lz4 64-bit

Name: Lz4 64-bit<br />
Summary: A lossless compression library<br />
License: BSD-2-Clause/GPLv2 or later<br />
Version: 1.10.0<br />
URL: [https://github.com/lz4/lz4/releases/download/v1.10.0/lz4-1.10.0.tar.gz](https://github.com/lz4/lz4/releases/download/v1.10.0/lz4-1.10.0.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 4.2 MiB<br />

## Configuration

The sources for Lz4 do not have a traditional configuration script.

## Compilation and Installation

To compile Lz4 64-bit, run the following command:

```bash
make -j1 BUILD_STATIC=no PREFIX=/usr
```

Next, run the test suite:

```bash
make -j1 check
```

Finally, to install Lz4 64-bit into the build tree, run the following command:

```bash
make BUILD_STATIC=no PREFIX=/usr LIBDIR=/usr/lib64 install
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | lz4, lz4c (link to lz4), lz4cat (link to lz4), and unlz4 (link to lz4) |
| Installed Libraries | liblz4.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./XZ_32bit.md) XZ 32-bit | [HOME](../README.md) | Lz4 32-bit [>>](./Lz4_32bit.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Lz4_32bit.md) Lz4 32-bit | [HOME](../README.md) | ZStd 32-bit [>>](./ZStd_32bit.md) |

## ZStd 64-bit

Name: ZStd 64-bit<br />
Summary: Real-time compression library<br />
License: BSD-3-Clause AND GPL-2.0-only<br />
Version: 1.5.7<br />
URL: [https://github.com/facebook/zstd/releases/download/v1.5.7/zstd-1.5.7.tar.gz](https://github.com/facebook/zstd/releases/download/v1.5.7/zstd-1.5.7.tar.gz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 86 MiB<br />

## Configuration

There is no standard configuration scripting for ZStd.

## Compilation and Installation

To compile ZStd 64-bit, run the following command:

```bash
make prefix=/usr libdir=/usr/lib64
```

Now, run the test suite:

```bash
make check
```

Finally, to install ZStd 64-bit into the build tree, run the following command:

```bash
make prefix=/usr libdir=/usr/lib64 install
rm -v /usr/lib64/libzstd.a
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | zstd, zstdcat (link to zstd), zstdgrep, zstdless, zstdmt (link to zstd), and unzstd (link to zstd) |
| Installed Libraries | libzstd.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./Lz4_32bit.md) Lz4 32-bit | [HOME](../README.md) | ZStd 32-bit [>>](./ZStd_32bit.md) |

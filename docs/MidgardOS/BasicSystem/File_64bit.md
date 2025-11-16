# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibUniString_32bit.md) libunistring 32-bit | [HOME](../README.md) | File 32-bit [>>](./File_32bit.md) |

## File 64-bit

Name: File 64-bit<br />
Summary: A tool to determine the type of a file from heuristics<br />
License: BSD 2-Clause<br />
Version: 5.46<br />
URL: [http://ftp.astrom.com/pub/file/file-5.46.tar.xz](ftp://ftp.astron.com/pub/file/file-5.46.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 19 MiB<br />

## Configuration

To configure File 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64
```

## Compilation and Installation

To compile File 64-bit, run the following command:

```bash
make
```

Now, run the test suite:

```bash
make check
```

Finally, to install File 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libmagic.la
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | file |
| Installed Libraries | libmagic.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibUniString_32bit.md) libunistring 32-bit | [HOME](../README.md) | File 32-bit [>>](./File_32bit.md) |

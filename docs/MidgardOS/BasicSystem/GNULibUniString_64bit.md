# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ZStd_32bit.md) ZStd 32-bit | [HOME](../README.md) | GNU libunistring 32-bit [>>](./GNULibUniString_32bit.md) |

## libunistring 64-bit

Name: libunistring 64-bit<br />
Summary: Unicode string handling in C<br />
License: GPL v3 or later<br />
Version: 1.3<br />
URL: [https://ftp.gnu.org/gnu/libunistring/libunistring-1.3.tar.xz](https://ftp.gnu.org/gnu/libunistring/libunistring-1.3.tar.xz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 58 MiB<br />

## Configuration

To configure libunistring 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --docdir=/usr/share/doc/libunistring-1.3
```

## Compilation and Installation

To compile libunistring 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install libunistring 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libunistring.la
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | N/A |
| Installed Libraries | libunistring.so |
| Installed Directories | /usr/include/unistring, /usr/share/doc/libunistring-1.3 |

| Navigation |||
| --- | --- | ---: |
| [<<](./ZStd_32bit.md) ZStd 32-bit | [HOME](../README.md) | GNU libunistring 32-bit [>>](./GNULibUniString_32bit.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./IPUtils.md) IP Utils | [HOME](../README.md) | CUnit 32-bit [>>](./CUnit_32bit.md) |

## CUnit 64-bit

Name: CUnit 64-bit<br />
Summary: A library for writing unit tests in C<br />
License: LGPL v2.1<br />
Version: 2.1-3<br />
URL: [https://sourceforge.net/projects/cunit/files/CUnit/2.1-3/CUnit-2.1-3.tar.bz2](https://sourceforge.net/projects/cunit/files/CUnit/2.1-3/CUnit-2.1-3.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 544 KiB<br />

## Configuration

To configure CUnit 64-bit for install into the build root, run the following commands:

```bash
autoreconf -fiv
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --disable-static                    \
            --docdir=/usr/share/doc/CUnit-2.1-3
```

## Compilation and Installation

To compile CUnit 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install CUnit 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libcunit.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libcunit.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./IPUtils.md) IP Utils | [HOME](../README.md) | CUnit 32-bit [>>](./CUnit_32bit.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGPerf.md) GNU GPerf | [HOME](../README.md) | Expat 32-bit [>>](./Expat_32bit.md) |

## Expat 64-bit

Name: Expat 64-bit<br />
Summary: An XML stream parsing library<br />
License: MIT<br />
Version: 2.7.3<br />
URL: [https://github.com/libexpat/libexpat/releases/download/R_2_7_3/expat-2.7.3.tar.xz](https://github.com/libexpat/libexpat/releases/download/R_2_7_3/expat-2.7.3.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 14 MiB<br />

## Configuration

To configure Expat 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --docdir=/usr/share/doc/expat-2.7.3
```

## Compilation and Installation

To compile Expat 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install Expat 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libexpat.la
install -v -m644 doc/*.{html,css} /usr/share/doc/expat-2.7.3
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | xmlwf |
| Installed Libraries | libexpat.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGPerf.md) GNU GPerf | [HOME](../README.md) | Expat 32-bit [>>](./Expat_32bit.md) |

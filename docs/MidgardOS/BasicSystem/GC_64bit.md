# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook-XSL-Stylesheets.md) Docbook XSL Stylesheets | [HOME](../README.md) | GC 32-bit [>>](./GC_32bit.md) |

## GC 64-bit

Name: GC 64-bit<br />
Summary: A library implementing memory garbage collection for C and C++<br />
License: MIT<br />
Version: 8.2.10<br />
URL: [https://github.com/bdwgc/bdwgc/releases/download/v8.2.10/gc-8.2.10.tar.gz](https://github.com/bdwgc/bdwgc/releases/download/v8.2.10/gc-8.2.10.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.5 MiB<br />

## Configuration

To configure GC 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/gc-8.2.10   \
            --enable-cplusplus                  \
            --disable-gcj                       \
            --enable-shared                     \
            --disable-static
```

## Compilation and Installation

To compile GC 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GC 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libcord.la
rm -fv /usr/lib64/libgc.la
rm -fv /usr/lib64/libgccpp.la
rm -fv /usr/lib64/libgctba.la
mv -v /usr/share/doc/gc /usr/share/doc/gc-8.2.10
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libcord.so, libgc.so, libgccpp.so, and libgctba.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook-XSL-Stylesheets.md) Docbook XSL Stylesheets | [HOME](../README.md) | GC 32-bit [>>](./GC_32bit.md) |

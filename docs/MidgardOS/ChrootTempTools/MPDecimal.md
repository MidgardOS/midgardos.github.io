# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Zlib.md) ZLib | [HOME](../README.md) | Python 3.15 [>>](./Python315.md) |

## MPDecimal

Name: MPDecimal<br />
Summary: A library of fast C/C++ arbitrary precision decimal floating point arithmetic routines<br />
License: BSD 2-Clause License<br />
Version: 4.0.1<br />
URL: [https://www.bytereef.org/software/mpdecimal/releases/mpdecimal-4.0.1.tar.gz](https://www.bytereef.org/software/mpdecimal/releases/mpdecimal-4.0.1.tar.gz)<br />

## Configuration

To configure MPDecimal for installation into the build root, run the following command:

```bash
./configure --prefix=/usr       \
            --libdir=/usr/lib64 \
            --disable-static    \
            --docdir=/usr/share/doc/mpdecimal-4.0.1
```

## Compilation and Installation

To compile MPDecimal, run the following command:

```bash
make
```

Finally, to install MPDecimal into the build tree, run the following command:

```bash
make install
```

More details about this package are covered later in the core system build.


| Navigation |||
| --- | --- | ---: |
| [<<](./Zlib.md) ZLib | [HOME](../README.md) | Python 3.15 [>>](./Python315.md) |

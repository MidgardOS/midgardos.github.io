| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGMP.md) | [HOME](./README.md) | [>>](./CrossCompileMPC.md) |

# GNU Multi-Precision Floating-Point with Rounding Library

Name: MPFR<br />
Summary: A C library for multi-precision floating-point calculations<br />
License: LGPL v3+<br />
Version: 4.2.0<br />
URL: [https://www.mpfr.org/mpfr-4.2.0/](https://www.mpfr.org/mpfr-4.2.0/)<br />

## Configuration

To configure GNU MPFR for install into our cross-compilation root, run the following command:

```bash
LDFLAGS="-Wl,-rpath,/cross-tools/lib" \
./configure --prefix=/cross-tools --disable-static --with-gmp=/cross-tools \
            --enable-thread-safe
```

## Compilation and Installation

To compile GMP, run the following command:

```bash
make
```

Finally, to install GMP into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGMP.md) | [HOME](./README.md) | [>>](./CrossCompileMPC.md) |

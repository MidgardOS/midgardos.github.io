| Navigation |||
| --- | --- | ---: |
| [<<](./GMP.md) | [HOME](../README.md) | [>>](./MPC.md) |

# GNU Multi-Precision Floating-Point with Rounding Library

Name: MPFR<br />
Summary: A C library for multi-precision floating-point calculations<br />
License: LGPL v3+<br />
Version: 4.2.0<br />
URL: [https://www.mpfr.org/mpfr-4.2.0/](https://www.mpfr.org/mpfr-4.2.0/)<br />

## Configuration

To configure GNU MPFR for install into our cross-compilation root, run the following command:

```bash
sed -e 's/+01,234,567/+1,234,567 /' \
    -e 's/13.10Pd/13Pd/'            \
    -i tests/tsprintf.c
LDFLAGS="-Wl,-rpath,/cross-tools/lib64 -L${BRFS}/cross-tools/lib64" \
./configure --prefix=/cross-tools --libdir=/cross-tools/lib64 --disable-static \
            --with-gmp=/cross-tools --enable-thread-safe
```

## Compilation and Installation

To compile GNU MPFR, run the following command:

```bash
make
```

Finally, to install GNU MPFR into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GMP.md) | [HOME](../README.md) | [>>](./MPC.md) |

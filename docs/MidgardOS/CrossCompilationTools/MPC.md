| Navigation |||
| --- | --- | ---: |
| [<<](./MPFR.md) | [HOME](../README.md) | [>>](./ISL.md) |

# GNU Multi-Precision Calculation Library

Name: MPC<br />
Summary: A C library for multi-precision calculations<br />
License: LGPL v3+<br />
Version: 1.3.1<br />
URL: [https://ftp.gnu.org/gnu/mpc](https://ftp.gnu.org/gnu/mpc)<br />

## Configuration

To configure GNU MPC for install into our cross-compilation root, run the following command:

```bash
LDFLAGS="-Wl,-rpath,/cross-tools/lib" \
./configure --prefix=/cross-tools --disable-static --with-gmp=/cross-tools \
            --with-mpfr=/cross-tools
```

## Compilation and Installation

To compile GNU MPC, run the following command:

```bash
make
```

Finally, to install GNU MPC into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./MPFR.md) | [HOME](../README.md) | [>>](./ISL.md) |

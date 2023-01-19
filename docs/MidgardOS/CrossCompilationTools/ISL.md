| Navigation |||
| --- | --- | ---: |
| [<<](./MPC.md) | [HOME](../README.md) | [>>](./GNUBinutils.md) |

# Integer Set Library

Name: ISL<br />
Summary: A library for handling Integer Sets bounded by linear constraints<br />
License: MIT<br />
Version: 0.25<br />
URL: [https://libisl.sourceforge.io](https://libisl.sourceforge.io)<br />

## Configuration

To configure ISL for install into our cross-compilation root, run the following command:

```bash
LDFLAGS="-Wl,-rpath,/cross-tools/lib" \
./configure --prefix=/cross-tools --disable-static --with-gmp-prefix=/cross-tools
```

## Compilation and Installation

To compile ISL, run the following command:

```bash
make
```

Finally, to install ISL into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./MPC.md) | [HOME](../README.md) | [>>](./GNUBinutils.md) |

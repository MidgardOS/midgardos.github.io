| Navigation |||
| --- | --- | ---: |
| [<<](./TempToolsMPC.md) | [HOME](./README.md) | [>>](./TempToolsZLib.md) |

# Integer Set Library

Name: ISL<br />
Summary: A library for handling Integer Sets bounded by linear constraints<br />
License: MIT<br />
Version: 0.25<br />
URL: [https://libisl.sourceforge.io](https://libisl.sourceforge.io)<br />

## Configuration

To configure ISL for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/tools --build=${BRFS_HOST} --host=${BRFS_TARGET} --libdir=/tools/lib64 \
    --enable-shared --enable-static
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
| [<<](./TempToolsMPC.md) | [HOME](./README.md) | [>>](./TempToolsZLib.md) |

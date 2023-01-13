| Navigation |||
| --- | --- | ---: |
| [<<](./TempToolsOverview.md) | [HOME](./README.md) | [>>](./TempToolsMPFR.md) |

# GNU Math Precision Library

Name: GMP<br />
Summary: A library for working with large numbers<br />
License: GPL v3+<br />
Version: 6.2.1<br />
URL: [http://ftp.gnu.org/pub/gnu/gmp](http://ftp.gnu.org/pub/gnu/gmp)<br />

## Configuration

To configure GMP for install into our cross-compilation root, run the following command:

```bash
CC_FOR_BUILD=gcc \
./configure --prefix=/tools --build=${BRFS_HOST} --host=${BRFS_TARGET} \
    --libdir=/tools/lib64 --enable-cxx --enable-fat
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
| [<<](./TempToolsOverview.md) | [HOME](./README.md) | [>>](./TempToolsMPFR.md) |

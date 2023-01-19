| Navigation |||
| --- | --- | ---: |
| [<<](./ISL.md) | [HOME](../README.md) | [>>](./GNUGCCp1.md) |

# GNU Binutils

Name: binutils<br />
Summary: A suite of tools for working with executable files<br />
License: LGPL v3.0+<br />
Version: 2.39<br />
URL: [https://ftp.gnu.org/gnu/binutils](https://ftp.gnu.org/gnu/binutils)<br />

## Configuration

To configure GNU Binutils for install into our cross-compilation root, run the following command:

```bash
mkdir -p build
cd build
AR=ar AS=as \
../configure \
    --prefix=/cross-tools --host=${BRFS_HOST} \
    --target=${BRFS_TARGET} --with-sysroot=${BRFS} \
    --with-lib-path=/tools/lib:/tools/lib64 --disable-nls --disable-static \
    --enable-64-bit-bfd --enable-gold=yes --enable-plugins --enable-threads \
    --disable-werror
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Binutils, run the following command:

```bash
make
```

Finally, to install GNU Binutils into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./ISL.md) | [HOME](../README.md) | [>>](./GNUGCCp1.md) |

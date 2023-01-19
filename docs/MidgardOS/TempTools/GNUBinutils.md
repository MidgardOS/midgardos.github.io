| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib.md) | [HOME](../README.md) | [>>](./GNUGCC.md) |

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
../configure --prefix=/tools --libdir=/tools/lib64 --with-lib-path=/tools/lib64:/tools/lib \
    --build=${BRFS_HOST} --host=${BRFS_TARGET} --target=${BRFS_TARGET} --disable-nls \
    --enable-shared --enable-64-bit-bfd --enable-gold=yes --enable-plugins --with-system-zlib \
    --enable-threads
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
| [<<](./ZLib.md) | [HOME](../README.md) | [>>](./GNUGCC.md) |

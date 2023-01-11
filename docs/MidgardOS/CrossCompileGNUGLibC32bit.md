| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileISL.md) | [HOME](./README.md) | [>>](./CrossCompileGNUGCCP1.md) |

# GNU C Library - 32-bit

Name: glibc-32bit<br />
Summary: The GNU C language runtime library - 32-bit<br />
License: GPL v2.0+/LGPL 2.1+<br />
Version: 2.36<br />
URL: [https://ftp.gnu.org/gnu/glibc](https://ftp.gnu.org/gnu/glibc)<br />

## Configuration

To configure GNU C Library 32-bit build for install into our cross-compilation root, run the following command:

```bash
mkdir -v glibc-build
cd glibc-build
BUILD_CC="gcc" CC="${BRFS_TARGET}-gcc ${BUILD32}" \
AR="${BRFS_TARGET}-ar" RANLIB="${BRFS_TARGET}-ranlib" \
../configure \
    --prefix=/tools \
    --host=${BRFS_TARGET32} \
    --build=${BRFS_HOST} \
    --enable-kernel=4.1.2 \
    --with-binutils=/cross-tools/bin \
    --with-headers=/tools/include \
    --enable-obsolete-rpc
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU C Library 32-bit build, run the following command:

```bash
make
```

Finally, to install the GNU C Library 32-bit build into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileISL.md) | [HOME](./README.md) | [>>](./CrossCompileGNUGCCP1.md) |

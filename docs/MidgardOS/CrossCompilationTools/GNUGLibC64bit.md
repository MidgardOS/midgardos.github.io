| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC32bit.md) | [HOME](../README.md) | [>>](./GNUGCCp2.md) |

# GNU C Library - 64-bit

Name: glibc-64bit<br />
Summary: The GNU C language runtime library - 64-bit<br />
License: GPL v2.0+/LGPL 2.1+<br />
Version: 2.36<br />
URL: [https://ftp.gnu.org/gnu/glibc](https://ftp.gnu.org/gnu/glibc)<br />

## Configuration

To configure GNU C Library 32-bit build for install into our cross-compilation root, run the following command:

```bash
mkdir -v glibc-build
cd glibc-build
echo "libc_cv_slibdir=/tools/lib64" >> config.cache
BUILD_CC="gcc" CC="${BRFS_TARGET}-gcc ${BUILD64}" \
AR="${BRFS_TARGET}-ar" RANLIB="${BRFS_TARGET}-ranlib" \
../configure \
    --prefix=/tools \
    --host=${BRFS_TARGET} \
    --build=${BRFS_HOST} \
    --libdir=/tools/lib64 \
    --enable-kernel=4.1.2 \
    --with-binutils=/cross-tools/bin \
    --with-headers=/tools/include \
    --enable-obsolete-rpc \
    --cache-file=config.cache
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
| [<<](./GNUGLibC32bit.md) | [HOME](../README.md) | [>>](./GNUGCCp2.md) |

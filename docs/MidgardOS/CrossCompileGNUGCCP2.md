| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGNUGLibC64bit.md) | [HOME](./README.md) | [>>](./BasicToolsOverview.md) |

# GNU Compiler Collection - Pass 2 - Dynamically Linked

Name: gcc<br />
Summary: A suite of compiler tools<br />
License: GPL v3.0+<br />
Version: 12.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc](https://ftp.gnu.org/gnu/gcc)<br />

## Pre-Configuration

Unlike earlier cross-compilation tools, GCC requires some extra steps to allow the binaries generated
to utilize the earlier packages, such as binutils, etc.

First, apply the following patch linked [here](./patches/gcc-12.2.0-specs.patch):

```bash
patch -Np1 -i ../gcc-12.2.0-specs.patch
```

Next, modify the StartFile spec to look in `tools`:

```bash
echo -en '\n#undef STANDARD_STARTFILE_PREFIX_1\n#define STANDARD_STARTFILE_PREFIX_1 "/tools/lib/"\n' >> gcc/config/linux.h
echo -en '\n#undef STANDARD_STARTFILE_PREFIX_2\n#define STANDARD_STARTFILE_PREFIX_2 ""\n' >> gcc/config/linux.h
```

## Configuration

To configure the GNU Compiler Collection for install into our cross-compilation root, run the following command:

```bash
mkdir -p gcc-build
cd gcc-build
AR=ar \
LDFLAGS="-Wl,-rpath,/cross-tools/lib" \
../configure \
    --prefix=/cross-tools \
    --build=${BRFS_HOST} \
    --host=${BRFS_HOST} \
    --target=${BRFS_TARGET} \
    --with-sysroot=${BRFS} \
    --with-local-prefix=/tools \
    --with-native-system-header-dir=/tools/include \
    --disable-static \
    --with-mpfr=/cross-tools \
    --with-gmp=/cross-tools \
    --with-mpc=/cross-tools \
    --with-isl=/cross-tools \
    --enable-languages=c,c++ \
    --with-glibc-version=2.36
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Compiler Collection, run the following command:

```bash
make AS_FOR_TARGET="${BRFS_TARGET}-as" LD_FOR_TARGET="${BRFS_TARGET}-ld"
```

Finally, to install GNU Compiler Collection into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGNUGLibC64bit.md) | [HOME](./README.md) | [>>](./BasicToolsOverview.md) |

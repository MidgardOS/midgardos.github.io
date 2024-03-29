| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) | [HOME](../README.md) | [>>](./GNUGLibC32bit.md) |

# GNU Compiler Collection - Pass 1 - Static, No Threads

Name: gcc<br />
Summary: A suite of compiler tools<br />
License: GPL v3.0+<br />
Version: 12.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc](https://ftp.gnu.org/gnu/gcc)<br />

## Pre-Configuration

Unlike earlier cross-compilation tools, GCC requires some extra steps to allow the binaries generated
to utilize the earlier packages, such as binutils, etc.

## Configuration

To configure the GNU Compiler Collection for install into our cross-compilation root, run the following command:

```bash
mkdir -p gcc-build
cd gcc-build
LDFLAGS="-Wl,-rpath,/cross-tools/lib64 -L/cross-tools/lib64" \
mlist=m64,m32,mx32 \
../configure                  \
    --target=$BRFS_TARGET                          \
    --prefix=$BRFS/tools                           \
    --with-glibc-version=2.37                      \
    --with-sysroot=$BRFS                           \
    --with-newlib                                  \
    --without-headers                              \
    --enable-default-pie                           \
    --enable-default-ssp                           \
    --enable-initfini-array                        \
    --disable-nls                                  \
    --disable-shared                               \
    --enable-multilib --with-multilib-list=$mlist  \
    --disable-decimal-float                        \
    --disable-threads                              \
    --disable-libatomic                            \
    --disable-libgomp                              \
    --disable-libquadmath                          \
    --disable-libssp                               \
    --disable-libvtv                               \
    --disable-libstdcxx                            \
    --with-mpfr=/cross-tools                       \
    --with-gmp=/cross-tools                        \
    --with-mpc=/cross-tools                        \
    --with-isl=/cross-tools                        \
    --enable-languages=c,c++
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Compiler Collection, run the following command:

```bash
make
```

Next, to install GNU Compiler Collection into the cross-tools tree, run the following command:

```bash
make install
```

Finally, to allow other builds to be successful, we need to create a dummy system header for `limits.h`, as the one that
GCC installs for itself is a partial header that references the system `limits.h`. At this point, that header is not yet
installed. While this works for building GLibC (the next package in the list), this will break with other tools later in
the stack.

```bash
cd ..
cat gcc/limitx.h gcc/glimits.h gcc/limity.h > \
  `dirname $(${BRFS_TARGET}-gcc -print-libgcc-file-name)`/install-tools/include/limits.h
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) | [HOME](../README.md) | [>>](./GNUGLibC32bit.md) |

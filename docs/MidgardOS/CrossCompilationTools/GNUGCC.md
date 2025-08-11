| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) | [HOME](../README.md) | [>>](./GNUGLibC64bit.md) |

# GNU Compiler Collection - Pass 1 - Static, No Threads

Name: gcc<br />
Summary: A suite of compiler tools<br />
License: GPL v3.0+<br />
Version: 15.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc](https://ftp.gnu.org/gnu/gcc)<br />

## Pre-Configuration

Unlike earlier cross-compilation tools, GCC requires some extra steps to allow the binaries generated
to utilize the earlier packages, such as binutils, etc. Much of the change requires MPC, MPFR, and GMP sources
installed into the source tree to build correctly. To do so, run the following commands:

```bash
tar xvf ../mpfr-4.2.2.tar.xz
mv -v mpfr-4.2.2 mpfr
tar xvf ../gmp-6.3.0.tar.xz
mv -v gmp-6.3.0 gmp
tar xvf ../mpc-1.3.1.tar.gz
mv -v mpc-1.3.1 mpc
```

Next, we need to apply a small fix to allow older 32bit pre-compiled Linux binaries and Windows PE 32bit binaries to run correctly on MidgardOS. To do so, run the following command:

```bash
sed '/STACK_REALIGN_DEFAULT/s/0/(!TARGET_64BIT \&\& TARGET_SSE)/' \
      -i gcc/config/i386/i386.h
```

## Configuration

To configure the GNU Compiler Collection for install into our cross-compilation root, run the following command:

```bash
mkdir -v build && cd build
../configure                                       \
    --target=$BRFS_TARGET                          \
    --prefix=$BRFS/tools                           \
    --with-glibc-version=2.42                      \
    --with-sysroot=$BRFS                           \
    --with-newlib                                  \
    --without-headers                              \
    --enable-default-pie                           \
    --enable-default-ssp                           \
    --enable-initfini-array                        \
    --with-bugurl="https://github.com/MidgardOS/MidgardOS/issues" \
	--with-pkgversion="MidgardOS"                  \
    --disable-nls                                  \
    --disable-shared                               \
    --enable-multilib                              \
    --with-arch-32=x86-64-v2                       \
    --with-arch=x86-64-v2                          \
    --with-tune=generic                            \
    --disable-decimal-float                        \
    --disable-threads                              \
    --disable-libatomic                            \
    --disable-libgomp                              \
    --disable-libquadmath                          \
    --disable-libssp                               \
    --disable-libvtv                               \
    --disable-libstdcxx                            \
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
cd /tools/bin
for command in "addr2lin" "ar" "as" 'c++' 'c++filt' "cpp" "elfedit" 'g++' "gcc" "gcc-15.2.0" "gcc-ar" "gcc-nm" "gcc-ranlib" \
               "gconv" "gconv-dump" "gconv-tool" "gprof" "ld" "ld.bfd" "lto-dump" "nm" "objcopy" "objdump" "ranlib" \
               "readelf" "size" "strings" "strip"; do
    ln -sv x86_64-unknown-linux-gnu-$command $command
done
cd -
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) | [HOME](../README.md) | [>>](./GNUGLibC64bit.md) |

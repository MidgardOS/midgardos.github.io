# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) GNU Binutils | [HOME](../README.md) | Entering the Build Root[>>](./ChrootBuildingTempTools.md) |

## GNU Compiler Collection - Pass 2

Name: gcc<br />
Summary: A suite of compiler tools<br />
License: GPL v3.0+<br />
Version: 15.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc/gcc-15.2.0.tar.xz](https://ftp.gnu.org/gnu/gcc/gcc-15.2.0.tar.xz)<br />

## Pre-Configuration

Like before with the cross-compilation tools, GCC requires some extra steps to allow the binaries generated
to utilize the earlier packages, such as the new binutils, etc.

```bash
tar xvf ../mpfr-4.2.2.tar.xz
mv -v mpfr-4.2.2 mpfr
tar xvf ../gmp-6.3.0.tar.xz
mv -v gmp-6.3.0 gmp
tar xvf ../mpc-1.3.1.tar.gz
mv -v mpc-1.3.1 mpc
```

Next, we need to apply a small fix to allow older 32-bit pre-compiled Linux binaries and Windows PE 32bit binaries to run correctly on MidgardOS. To do so, run the following command:

```bash
sed '/STACK_REALIGN_DEFAULT/s/0/(!TARGET_64BIT \&\& TARGET_SSE)/' \
      -i gcc/config/i386/i386.h
```

Finally, we need to override the build rules to allow libgcc and libstdc++ so that they can be built with POSIX threads enabled:

```bash
sed '/thread_header =/s/@.*@/gthr-posix.h/' \
    -i libgcc/Makefile.in libstdc++-v3/include/Makefile.in
```

## Configuration

To configure the GNU Compiler Collection for install into our cross-compilation root, run the following command:

```bash
mkdir -v build && cd build
mlist=m32,m64                                      \
../configure                                       \
    --build=$BRFS_HOST                             \
    --host=$BRFS_TARGET                            \
    --target=$BRFS_TARGET                          \
    --prefix=/usr                                  \
    --libdir=/usr/lib64                            \
    --libexecdir=/usr/lib64                        \
    --with-build-sysroot=$BRFS                     \
    --enable-default-pie                           \
    --enable-default-ssp                           \
    --with-glibc-version=2.42                      \
    --with-sysroot=$BRFS                           \
    --disable-nls                                  \
    --with-bugurl="https://github.com/MidgardOS/MidgardOS/issues" \
	--with-pkgversion="MidgardOS"                  \
    --enable-multilib                              \
    --with-multilib-list=$mlist                    \
    --with-arch-32=x86-64-v2                       \
    --with-arch=x86-64-v2                          \
    --with-tune=generic                            \
    --disable-libatomic                            \
    --disable-libgomp                              \
    --disable-libquadmath                          \
    --disable-libsanitizer                         \
    --disable-libssp                               \
    --disable-libvtv                               \
    --enable-languages=c,c++                       \
    LDFLAGS_FOR_TARGET=-L$PWD/$BRFS_TARGET/libgcc
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Compiler Collection, run the following command:

```bash
make
```

Next, to install GNU Compiler Collection into the cross-tools tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

Finally, to allow other builds to be successful, we need to create a dummy system header for `limits.h`, as the one that
GCC installs for itself is a partial header that references the system `limits.h`. At this point, that header is not yet
installed. While this works for building GLibC (the next package in the list), this will break with other tools later in
the stack.

```bash
cd $BRFS/usr/bin
ln -sv x86_64-unknown-linux-gnu-gcc-15.2.0 gcc-15.2.0
ln -sv gcc cc
cd -
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) GNU Binutils | [HOME](../README.md) | Entering the Build Root [>>](./ChrootBuildingTempTools.md) |

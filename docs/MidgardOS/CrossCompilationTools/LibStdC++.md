# Section 2 - Cross-Compilation Toolchain

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC32bit.md) GNU GLibC - 32-bit | [HOME](../README.md) | Temporary Tools Overview [>>](../TempTools/Overview.md) |

## GNU Compiler Collection - LibStdC++

Name: libstdc++<br />
Summary: The C++ language runtime libraries<br />
License: GPL v3.0+<br />
Version: 15.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc/gcc-15.2.0.tar.xz](https://ftp.gnu.org/gnu/gcc/gcc-15.2.0.tar.xz)<br />

## Configuration

To configure LibStdC++ from the the GNU Compiler Collection into our cross-compilation root, run the following command:

```bash
mkdir -v build && cd build
../libstdc++-v3/configure   \
    --host=$BRFS_TARGET     \
    --build=$BRFS_HOST      \
    --prefix=/usr           \
    --enable-multilib       \
    --disable-nls           \
    --disable-libstdcxx-pch \
    --with-gxx-include-dir=/tools/$BRFS_TARGET/include/c++/15.2.0
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile LibStdC++ from the GNU Compiler Collection, run the following command:

```bash
make
```

Next, to install LibStdC++ from the GNU Compiler Collection into the cross-tools tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

Finally, remove the libtool archive files from the build root for libstdc++ and others since they can cause issues with other builds. To do so, run the following command:

```bash
rm -v $BRFS/usr/lib/lib{stdc++{,exp,fs},supc++}.la
rm -v $BRFS/usr/lib64/lib{stdc++{,exp,fs},supc++}.la
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC32bit.md) GNU GLibC - 32-bit | [HOME](../README.md) | Temporary Tools Overview [>>](../TempTools/Overview.md) |

# Section 2 - Cross-Compilation Toolchain

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC64bit.md) GNU GLibC - 64-bit | [HOME](../README.md) | GNU C++ Runtime [>>](./LibStdC++.md) |

## GNU C Library - 32-bit

Name: glibc-32bit<br />
Summary: The GNU C language runtime library - 32-bit<br />
License: GPL v2.0+/LGPL 2.1+<br />
Version: 2.36<br />
URL: [https://ftp.gnu.org/gnu/glibc](https://ftp.gnu.org/gnu/glibc)<br />

## Configuration

To configure GNU C Library 32-bit build for install into our cross-compilation root, run the following command from inside the previously created `build` directory:

```bash
make clean
find .. -name "*.a" -delete
CC="$BRFS_TARGET-gcc -m32" CXX="$BRFS_TARGET-g++ -m32" \
../configure \
    --prefix=/tools                   \
    --host=${BRFS_TARGET32}           \
    --build=${BRFS_HOST}              \
    --enable-kernel=5.4               \
    --with-headers=$BRFS/usr/include  \
    --enable-bind-now                 \
    --enable-stackguard-randomization \
    --enable-tunables                 \
    --with-bugurl="https://github.com/MidgardOS/MidgardOS/Issues" \
    --libdir=/usr/lib                 \
    --libexecdir=/usr/lib             \
    libc_cv_slibdir=/usr/lib
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU C Library 32-bit build, run the following command:

```bash
make
```

Finally, to install the GNU C Library 32-bit build into the cross-tools tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -a DESTDIR/usr/lib $BRFS/usr/ ||:
install -v -m 644 DESTDIR/tools/include/gnu/{lib-names,stubs}-32.h $BRFS/usr/include/gnu
ln -sfv ../lib64/ld-linux-x86-64.so.2 ${BRFS}/lib64/ld-lsb-x86-64.so.3
```

If an operation not permitted error comes up around preserving times for the `/MidgardOS/usr/lib`, please ignore this, as the `lib` directory already exists.

## Validation

It is critically important to validate that the new GLibC 32-bit build is working as expected in tandom with the temporary compiler instance in the system. To do so, run the following commands:

```bash
echo 'int main(){}' > dummy.c
$BRFS_TARGET-gcc -m32 dummy.c
readelf -l a.out | grep '/ld-linux'
```

This should return the following output:
```
      [Requesting program interpreter: /lib/ld-linux.so.2]
```

If this is successful, you can move on to the next step in the guide.

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC64bit.md) GNU GLibC - 64-bit | [HOME](../README.md) | GNU C++ Runtime [>>](./LibStdC++.md) |

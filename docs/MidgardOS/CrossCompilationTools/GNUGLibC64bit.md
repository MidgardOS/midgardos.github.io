| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) | [HOME](../README.md) | [>>](./GNUGLibC32bit.md) |

# GNU C Library - 64-bit

Name: glibc-64bit<br />
Summary: The GNU C language runtime library - 64-bit<br />
License: GPL v2.0+/LGPL 2.1+<br />
Version: 2.36<br />
URL: [https://ftp.gnu.org/gnu/glibc](https://ftp.gnu.org/gnu/glibc)<br />

## Configuration

To configure GNU C Library 32-bit build for install into our cross-compilation root, run the following command:

```bash
patch -Np1 -i ../glibc-2.42-fhs-1.patch
mkdir -v build && cd build
echo "rootsbindir=/usr/sbin" > configparms
../configure \
    --prefix=/usr                      \
    --host=${BRFS_TARGET}              \
    --build=${BRFS_HOST}               \
    --libdir=/usr/lib64                \
    --libexecdir=/usr/lib64            \
    --enable-bind-now                  \
    --enable-multi-arch                \
    --enable-cet                       \
    --enable-stackguard-randomization  \
    --enable-tunables                  \
    --with-bugurl="https://github.com/MidgardOS/MidgardOS/Issues" \
    libc_cv_slibdir=/usr/lib64         \
    --enable-kernel=5.4
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU C Library 32-bit build, run the following command:

```bash
make
```

Finally, to install the GNU C Library 32-bit build into the cross-tools tree, run the following command:

```bash
make DESTDIR=${BRFS} install
sed '/RTLDLIST=/s@/usr@@g' -i $BRFS/usr/bin/ldd
```

## Validating the Installation so Far

To ensure that the basic cross-compilation tool chain is installed correctly, run the following commands:

```bash
echo 'int main(){}' | $BRFS_TARGET-gcc -x c - -v -Wl,--verbose &> dummy.log
readelf -l a.out | grep ': /lib'
```

These two commands should return the following output:
```
      [Requesting program interpreter: /lib64/ld-linux-x86-64.so.2]
```

Next, make sure that the correct start files are being used:
```bash
grep -E -o "$LFS/lib.*/S?crt[1in].*succeeded" dummy.log
```

This should output something like the following:
```
/lib/../lib64/Scrt1.o succeeded
/lib/../lib64/crti.o succeeded
/lib/../lib64/crtn.o succeeded
```

Now verify that the compiler is using the headers in the build root:
```bash
grep -B3 "^ $BRFS/usr/include" dummy.log
```

This should output something like the following:
```
 /MidgardOS/tools/lib/gcc/x86_64-unknown-linux-gnu/15.2.0/include
 /MidgardOS/usr/local/include
 /MidgardOS/tools/lib/gcc/x86_64-unknown-linux-gnu/15.2.0/include-fixed
 /MidgardOS/usr/include
```

Now verify that the new linker is being used and is using the correct search paths:
```bash
grep 'SEARCH.*/usr/lib' dummy.log |sed 's|; |\n|g'
```

This should output something like the following:
```
SEARCH_DIR("=/tools/x86_64-unknown-linux-gnu/lib64")
SEARCH_DIR("=/usr/local/lib64")
SEARCH_DIR("=/lib64")
SEARCH_DIR("=/usr/lib64")
SEARCH_DIR("=/tools/x86_64-unknown-linux-gnu/lib")
SEARCH_DIR("=/usr/local/lib")
SEARCH_DIR("=/lib")
SEARCH_DIR("=/usr/lib");
```

Next, ensure we are using the correct C library:
```bash
grep "/lib.*/libc.so.6 " dummy.log
```

This should output the following line:
```
attempt to open /MidgardOS/usr/lib64/libc.so.6 succeeded
```

Finally, validate that the compiler is using the correct dynamic linker:
```bash
grep found dummy.log
```

This should output the following string:
```
found ld-linux-x86-64.so.2 at /MidgardOS/usr/lib64/ld-linux-x86-64.so.2
```

If any of the above validation steps fail to return similar output, the build root is unusable and needs rebuilt after validating that all steps have been followed exactly as explained in this guide. It is imperative that all of the validation steps work as expected before going any farther.

If everything works as expected, delete the current build directory, but do NOT remove the unpacked sources, as they will be used in the next step.

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) | [HOME](../README.md) | [>>](./GNUGLibC32bit.md) |

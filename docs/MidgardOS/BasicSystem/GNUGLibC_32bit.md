# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC_64bit.md) GNU GLibC 64-bit | [HOME](../README.md) | ZLib 64-bit [>>](./ZLib_64bit.md) |

## GNU GLibC 32-bit

Name: GNU GLibC 32-bit<br />
Summary: The GNU C language runtime library - 64-bit<br />
License: GPL v3 or later<br />
Version: 2.42<br />
URL: [https://ftp.gnu.org/gnu/glibc/glibc-2.42.tar.xz](https://ftp.gnu.org/gnu/glibc/glibc-2.42.tar.xz)<br />

Average Build Time: 6 SBU<br />
Used Install Space: 1.2 GiB<br />

## Configuration

To configure GNU GLibC 32-bit for install into the build root, run the following command:

```bash
rm -rf ./*
find .. -name "*.a" -delete
CC="gcc -m32" CXX="g++ -m32" \
../configure                                                         \
      --prefix=/usr                                                  \
      --host=i686-pc-linux-gnu                                       \
      --build=$(../scripts/config.guess)                             \
      --libdir=/usr/lib                                              \
      --libexecdir=/usr/lib                                          \
      --disable-werror                                               \
      libc_cv_slibdir=/usr/lib                                       \
      --enable-stackguard-randomization                              \
      --enable-tunables                                              \
      --with-bugurl="https://github.com/MidgardOS/MidgardOS/Issues"  \
      --enable-kernel=5.4
```

## Compilation and Installation

To compile GNU GLibC 32-bit, run the following command:

```bash
make
```

Finally, to install GNU GLibC 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -a DESTDIR/usr/lib/* /usr/lib/
install -vm644 DESTDIR/usr/include/gnu/{lib-names,stubs}-32.h /usr/include/gnu/
```

Once the 32-bit libraries are installed, it is a good idea to validate that the C runtime is working as expected. To do so, run the following commands:

```bash
echo 'int main(){}' > dummy.c
gcc -m32 dummy.c
readelf -l a.out | grep '/ld-linux'
```

If everything worked correctly, the output from the above commands should look like the following:

```
      [Requesting program interpreter: /lib/ld-linux.so.2]
```

## Contents

See the previous 64-bit build of GLibC to read about which libraries have been installed.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGLibC_64bit.md) GNU GLibC 64-bit | [HOME](../README.md) | ZLib 64-bit [>>](./ZLib_64bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Shadow_32bit-pass1.md) Shadow Utils 32-bit | [HOME](../README.md) | NCurses 64-bit [>>](./NCurses_64bit.md) |

## GNU Compiler Collection

Name: GNU Compiler Collection<br />
Summary: A suite of compiler tools<br />
License: GPL v3.0+<br />
Version: 15.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc/gcc-15.2.0.tar.xz](https://ftp.gnu.org/gnu/gcc/gcc-15.2.0.tar.xz)<br />

Average Build Time: 46 SBU<br />
Used Install Space: 1.8 GiB<br />

## Pre-Configuration

To allow older 32-bit pre-compiled Linux binaries and Windows PE 32-bit executables to run correctly under MidgardOS, we need to apply a small fix to the i386 C header. To do so, run the following command:

```bash
sed '/STACK_REALIGN_DEFAULT/s/0/(!TARGET_64BIT \&\& TARGET_SSE)/' \
      -i gcc/config/i386/i386.h
```

## Configuration

To configure the GNU Compiler Collection for install into the build root, run the following commands:

```bash
mkdir -v build && cd build
../configure                                                        \
    --prefix=/usr                                                   \
    --libdir=/usr/lib64                                             \
    --libexecdir=/usr/lib64                                         \
    --enable-languages=c,c++                                        \
    --enable-default-pie                                            \
    --enable-default-ssp                                            \
    --with-glibc-version=2.42                                       \
    --with-bugurl="https://github.com/MidgardOS/MidgardOS/issues"   \
    --with-pkgversion="MidgardOS"                                   \
    --enable-multilib                                               \
    --with-arch-32=x86-64-v2                                        \
    --with-arch=x86-64-v2                                           \
    --with-tune=generic                                             \
    --disable-bootstrap                                             \
    --disable-fixincludes                                           \
    --with-system-zlib
```

## Compilation and Installation

To compile the GNU Compiler Collection, run the following command:

```bash
make
```

Now, run the test suite:

```bash
sed -e '/cpython/d' -i ../gcc/testsuite/gcc.dg/plugin/plugin.exp
chown -R tester .
su tester -c "PATH=$PATH make -k check"
../contrib/test_summary | grep -A7 Summ
chown -R root:root .
```

Note, it is not expected that all tests will pass. GCC has some "future" looking tests for features under development in the tree. It is know that tests related to `pr90579.c` will fail.

Finally, to install the GNU Compiler Collection into the build tree, run the following commands:

```bash
make install
chown -v -R root:root /usr/lib64/gcc/$(gcc -dumpmachine)/15.2.0/include{,-fixed}
ln -svr /usr/bin/cpp /usr/lib64
ln -sv gcc.1 /usr/share/man/man1/cc.1
ln -sfv ../../lib64/gcc/$(gcc -dumpmachine)/15.2.0/liblto_plugin.so /usr/lib64/bfd-plugins/
```

Now that the last rebuild of GCC is complete before it is rebuilt for packaging as an RPM, we need to ensure that compiling and linking work as expected. To perform the first sanity check, run the following commands:

```bash
echo 'int main(){}' | cc -x c - -v -Wl,--verbose &> dummy.log
readelf -l a.out | grep ': /lib'
```

This should output the following:
```
[Requesting program interpreter: /lib64/ld-linux-x86-64.so.2]
```

Now, ensure that the compiler uses the correct start files:

```bash
grep -E -o '/usr/lib.*/S?crt[1in].*succeeded' dummy.log
```

This should output the following:
```
/usr/lib64/gcc/x86_64-pc-linux-gnu/15.2.0/../../../../lib64/Scrt1.o succeeded
/usr/lib64/gcc/x86_64-pc-linux-gnu/15.2.0/../../../../lib64/crti.o succeeded
/usr/lib64/gcc/x86_64-pc-linux-gnu/15.2.0/../../../../lib64/crtn.o succeeded
```

Next, check that the compiler is using the correct header files:

```bash
grep -B4 '^ /usr/include' dummy.log
```

This should print the following output:
```
#include <...> search starts here:
 /usr/lib64/gcc/x86_64-pc-linux-gnu/15.2.0/include
 /usr/local/include
 /usr/lib64/gcc/x86_64-pc-linux-gnu/15.2.0/include-fixed
 /usr/include
```

After this, check that the compiler is using the right linker with the correct search paths:

```bash
grep 'SEARCH.*/usr/lib' dummy.log |sed 's|; |\n|g'
```

This should return the following:
```
SEARCH_DIR("/usr/x86_64-pc-linux-gnu/lib64")
SEARCH_DIR("/usr/local/lib64")
SEARCH_DIR("/lib64")
SEARCH_DIR("/usr/lib64")
SEARCH_DIR("/usr/x86_64-pc-linux-gnu/lib")
SEARCH_DIR("/usr/local/lib")
SEARCH_DIR("/lib")
SEARCH_DIR("/usr/lib");
```

Now, ensure that the correct libc is being used:

```bash
grep "/lib.*/libc.so.6 " dummy.log
```

This should output the following:
```
attempt to open /usr/lib64/libc.so.6 succeeded
```

And finally, validate that GCC is using the correct dynamic linker:

```bash
grep found dummy.log
```

This should print the following:
```
found ld-linux-x86-64.so.2 at /usr/lib64/ld-linux-x86-64.so.2
```

Once everything is validated, clean up our test files:

```bash
rm -v a.out dummy.log
```

Finally, move a mis-placed file:

```bash
install -v -d -m 755 -o root -g root /usr/share/gdb/auto-load/usr/lib
mv -v /usr/lib/*gdb.py /usr/share/gdb/auto-load/usr/lib
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | c++, cc (link to gcc), cpp, g++, gcc, gcc-ar, gcc-nm, gcc-ranlib, gcov, gcov-dump, gcov-tool, and lto-dump |
| Installed Libraries | libasan.{a,so}, libatomic.{a,so}, libcc1.so, libgcc.a, libgcc_eh.a, libgcc_s.so, libgcov.a, libgomp.{a,so}, libhwasan.{a,so}, libitm.{a,so}, liblsan.{a,so}, liblto_plugin.so, libquadmath.{a,so}, libssp.{a,so}, libssp_nonshared.a, libstdc++.{a,so}, libstdc++exp.a, libstdc++fs.a, libsupc++.a, libtsan.{a,so}, and libubsan.{a,so} |

| Navigation |||
| --- | --- | ---: |
| [<<](./Shadow_32bit-pass1.md) Shadow Utils 32-bit | [HOME](../README.md) | NCurses 64-bit [>>](./NCurses_64bit.md) |

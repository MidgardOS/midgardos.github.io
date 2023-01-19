| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) | [HOME](../README.md) | [>>](./NCurses.md) |

# GNU Compiler Collection

Name: gcc<br />
Summary: A suite of compiler tools<br />
License: GPL v3.0+<br />
Version: 12.2.0<br />
URL: [https://ftp.gnu.org/gnu/gcc](https://ftp.gnu.org/gnu/gcc)<br />

## Pre-Configuration

Like before with the cross-compilation tools, GCC requires some extra steps to allow the binaries generated
to utilize the earlier packages, such as the new temporary tools tree like binutils, etc.

First, apply the following patch linked [here](./patches/gcc-12.2.0-specs.patch):

```bash
patch -Np1 -i ../gcc-12.2.0-specs.patch
```

Next, modify the StartFile spec to look in `tools`:

```bash
echo -en '\n#undef STANDARD_STARTFILE_PREFIX_1\n#define STANDARD_STARTFILE_PREFIX_1 "/tools/lib/"\n' >> gcc/config/linux.h
echo -en '\n#undef STANDARD_STARTFILE_PREFIX_2\n#define STANDARD_STARTFILE_PREFIX_2 ""\n' >> gcc/config/linux.h
```

Next, run the following `sed` command to ensure that the compilation will not run the `fixincludes` script, which will
break the temporary tools build environment:

```bash
cp -v gcc/Makefile.in{,.orig}
sed 's@\./fixinc\.sh@-c true@' gcc/Makefile.in.orig > gcc/Makefile.in
```

## Configuration

To configure the GNU Compiler Collection for install into our cross-compilation root, run the following command:

```bash
mkdir -pv gcc-build
cd gcc-build
../configure --prefix=/tools --libdir=/tools/lib64 --build=${BRFS_HOST} --host=${BRFS_HOST} --target=${BRFS_TARGET} \
    --with-local-prefix=/tools --with-mpfr=/tools --with-gmp=/tools --with-mpc=/tools --with-isl=/tools \
    --with-native-system-header-dir=/tools/include --enable-languages=c,c++
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Compiler Collection, run the following command:

```bash
make AS_FOR_TARGET="${AS}" LD_FOR_TARGET="${LD}"
```

Finally, to install GNU Compiler Collection into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBinutils.md) | [HOME](../README.md) | [>>](./NCurses.md) |

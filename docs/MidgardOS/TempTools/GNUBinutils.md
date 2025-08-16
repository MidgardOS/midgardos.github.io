# Section X - $SECTION_TITLE

| Navigation |||
| --- | --- | ---: |
| [<<](./Xz.md) Xz | [HOME](../README.md) | GNU Compiler Collection - pass 2 [>>](./GNUGCC.md) |

## GNU Binutils

Name: binutils<br />
Summary: A suite of tools for working with executable files<br />
License: LGPL v3.0+<br />
Version: 2.45<br />
URL: [https://ftp.gnu.org/gnu/binutils](https://ftp.gnu.org/gnu/binutils)<br />

## Configuration

To configure GNU Binutils for install into our cross-compilation root, run the following command:

```bash
sed '6031s/$add_dir//' -i ltmain.sh
mkdir -pv build && cd build
../configure --prefix=/usr                   \
             --libdir=/usr/lib64             \
             --libexecdir=/usr/lib64         \
             --build=${BRFS_HOST}            \
             --host=${BRFS_TARGET}           \
             --disable-nls                   \
             --enable-shared                 \
             --enable-gprofng=no             \
             --enable-64-bit-bfd             \
             --enable-new-dtags              \
             --enable-default-hash-style=gnu
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Binutils, run the following command:

```bash
make
```

Finally, to install GNU Binutils into the cross-tools tree, run the following command:

```bash
make DESTDIR=$BRFS install
rm -v $BRFS/usr/lib64/lib{bfd,ctf,ctf-nobfd,opcodes,sframe}.{a,la}
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./Xz.md) Xz | [HOME](../README.md) | GNU Compiler Collection - pass 2 [>>](./GNUGCC.md) |

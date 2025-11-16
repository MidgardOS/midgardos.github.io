# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./PkgConf.md) PkgConf | [HOME](../README.md) | GNU GMP 64-bit [>>](./GNUGMP_64bit.md) |

## GNU Binutils

Name: binutils<br />
Summary: A suite of tools for working with executable files<br />
License: LGPL v3.0+<br />
Version: 2.45<br />
URL: [https://ftp.gnu.org/gnu/binutils/binutils-2.45.tar.xz](https://ftp.gnu.org/gnu/binutils/binutils-2.45.tar.xz)<br />

Average Build Time: 1.6 SBU<br />
Used Install Space: 832 MiB<br />

## Configuration

To configure GNU Binutils for install into the build root, run the following commands:

```bash
mkdir -pv build && cd build
../configure --prefix=/usr           \
             --libdir=/usr/lib64     \
             --libexecdir=/usr/lib64 \
             --sysconfdir=/etc       \
             --enable-ld=default     \
             --enable-plugins        \
             --enable-shared         \
             --disable-werror        \
             --enable-64-bit-bfd     \
             --enable-new-dtags      \
             --with-system-zlib      \
             --enable-default-hash-style=gnu
```

Stay in the build directory until this package is installed.

## Compilation and Installation

To compile GNU Binutils, run the following command:

```bash
make tooldir=/usr
```

Next, run the test suite:

```bash
make -k check
grep '^FAIL:' $(find -name '*.log')
```

Finally, to install GNU Binutils into the build tree, run the following commands:

```bash
make tooldir=/usr install
rm -rfv /usr/lib64/lib{bfd,ctf,ctf-nobfd,gprofng,opcodes,sframe}.a /usr/share/doc/gprofng/
rm -rfv /usr/lib64/lib{bfd,ctf,ctf-nobfd,gprofng,opcodes,sframe}.la
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | addr2line, ar, as, c++filt, dwp, elfedit, gprof, gprofng, ld, ld.bfd, nm, objcopy, objdump, ranlib, readelf, size, strings, and strip |
| Installed Libraries | libbfd.so, libctf.so, libctf-nobfd.so, libgprofng.so, libopcodes.so, and libsframe.so |
| Installed Directories | /usr/lib/ldscripts |

| Navigation |||
| --- | --- | ---: |
| [<<](./PkgConf.md) PkgConf | [HOME](../README.md) | GNU GMP 64-bit [>>](./GNUGMP_64bit.md) |

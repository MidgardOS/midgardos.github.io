# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUFindutils.md) GNU Findutils | [HOME](../README.md) | GNU Grep [>>](./GNUGrep.md) |

## GNU Awk

Name: GNU Awk<br />
Summary: GNU Awk is a CLI tool for processing text data<br />
License: GPL v3 or later<br />
Version: 5.4.1<br />
URL: [https://ftp.gnu.org/gnu/gawk/gawk-5.4.1.tar.xz](https://ftp.gnu.org/gnu/gawk/gawk-5.4.1.tar.xz)<br />

## Pre-configuration

To fix a bug in the behaviour of GNU Awk that breaks compilation of future tools, apply the following patch:

```bash
patch -Np1 -i ../patches/gawk/gawk-5.4.1-fix-alignment-bug.patch
```

## Configuration

To configure GNU Awk for installation into the build root, run the following command:

```bash
sed -i 's/extras//' Makefile.in
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --host=$BRFS_TARGET     \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Awk, run the following command:

```bash
make
```

Finally, to install GNU Awk into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package are covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUFindutils.md) GNU Findutils | [HOME](../README.md) | GNU Grep [>>](./GNUGrep.md) |

# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMake.md) GNU Make | [HOME](../README.md) | GNU Sed [>>](./GNUSed.md) |

## GNU Patch

Name: GNU Patch<br />
Summary: A tool for making and applying differential patch files<br />
License: GPL v3 or later<br />
Version: 2.8<br />
URL: [https://ftp.gnu.org/gnu/patch/patch-2.8.tar.xz](https://ftp.gnu.org/gnu/patch/patch-2.8.tar.xz)<br />

## Configuration

To configure GNU Patch for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --host=$BRFS_TARET      \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Patch, run the following command:

```bash
make
```

Finally, to install GNU Patch into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMake.md) GNU Make | [HOME](../README.md) | GNU Sed [>>](./GNUSed.md) |

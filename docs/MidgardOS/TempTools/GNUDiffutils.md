# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUCoreutils.md) | [HOME](../README.md) | [>>](./File.md) |

## GNU Diffutils

Name: GNU Difftils<br />
Summary: Tools for displaying differences between files and directories<br />
License: GPL v3 or later<br />
Version: 3.9<br />
URL: [https://ftp.gnu.org/gnu/diffutils/diffutils-3.9.tar.xz](https://ftp.gnu.org/gnu/diffutils/diffutils-3.9.tar.xz)<br />

## Configuration

To configure GNU Diffutils for install into the build root, run the following command:

```bash
./configure --prefix=/usr                 \
            --host=$BRFS_TARGET           \
            gl_cv_func_strcasecmp_works=y \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Diffutils, run the following command:

```bash
make
```

Finally, to install GNU Diffutils into the build tree, run the following command:

```bash
make DESTDIR=${BRFS} install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUCoreutils.md) | [HOME](../README.md) | [>>](./File.md) |

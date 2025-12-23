# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUCoreutils.md) GNU Coreutils | [HOME](../README.md) | GNU AWK [>>](./GAWK.md) |

## GNU Diffutils

Name: GNU Diffutils<br />
Summary: Tools for displaying differences between files and directories<br />
License: GPL v3 or later<br />
Version: 3.12<br />
URL: [https://ftp.gnu.org/gnu/diffutils/diffutils-3.12.tar.xz](https://ftp.gnu.org/gnu/diffutils/diffutils-3.12.tar.xz)<br />

Average Build Time: 0.5 SBU<br />
Used Install Space: 2.8 MiB<br />

## Configuration

To configure GNU Diffutils for install into the build root, run the following command:

```bash
./configure --prefix=/usr                           \
            --libdir=/usr/lib64                     \
            --libexecdir=/usr/lib64                 \
            --docdir=/usr/share/doc/diffutils-3.12
```

## Compilation and Installation

To compile GNU Diffutils, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Diffutils into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | cmp, diff, diff3, and sdiff |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUCoreutils.md) GNU Coreutils | [HOME](../README.md) | GNU AWK [>>](./GAWK.md) |

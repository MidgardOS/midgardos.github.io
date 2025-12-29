# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMake.md) GNU Make | [HOME](../README.md) | GNU Tar [>>](./GNUTar.md) |

## GNU Patch

Name: GNU Patch<br />
Summary: A tool for applying differential patches to text files<br />
License: GPL v3.0 or later<br />
Version: 2.8<br />
URL: [https://ftp.gnu.org/gnu/patch/patch-2.8.tar.xz](https://ftp.gnu.org/gnu/patch/patch-2.8.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 700 KiB<br />

## Configuration

To configure GNU Patch for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/patch-2.8   \
            --sysconfdir=/etc
```

## Compilation and Installation

To compile GNU Patch, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Patch into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | patch |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMake.md) GNU Make | [HOME](../README.md) | GNU Tar [>>](./GNUTar.md) |

# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GAWK.md) | [HOME](../README.md) | [>>](./GNUGzip.md) |

## GNU Grep

Name: GNU Grep<br />
Summary: CLI tool for searching through text files<br />
License: GPL v3 or later<br />
Version: 3.9<br />
URL: [https://ftp.gnu.org/gnu/grep/grep-3.9.tar.xz](https://ftp.gnu.org/gnu/grep/grep-3.9.tar.xz)<br />

## Configuration

To configure GNU Grep for install into the build root, run the following command:

```bash
./configure --prefix=/usr       \
            --host=$BRFS_TARGET \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Grep, run the following command:

```bash
make
```

Finally, to install GNU Grep into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GAWK.md) | [HOME](../README.md) | [>>](./GNUGzip.md) |

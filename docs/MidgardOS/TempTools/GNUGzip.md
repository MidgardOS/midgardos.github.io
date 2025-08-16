# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGrep.md) | [HOME](../README.md) | [>>](./GNUMake.md) |

## GNU Gzip

Name: GNU Gzip<br />
Summary: The GNU compression and decompression tools<br />
License: GPL v3 or later<br />
Version: 1.14<br />
URL: [https://ftp.gnu.org/gnu/gzip/gzip-1.14.tar.xz](https://ftp.gnu.org/gnu/gzip/gzip-1.14.tar.xz)<br />

## Configuration

To configure GNU Gzip for install into the build root, run the following command:

```bash
./configure --prefix=/usr --host=$BRFS_TARGET
```

## Compilation and Installation

To compile GNU Gzip, run the following command:

```bash
make
```

Finally, to install GNU Gzip into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGrep.md) | [HOME](../README.md) | [>>](./GNUMake.md) |

# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUSed.md) GNU Sed | [HOME](../README.md) | Xz [>>](./Xz.md) |

## GNU Tar

Name: GNU Tar<br />
Summary: The Tape Archive Tool<br />
License: GPL v3 or later<br />
Version: 1.35<br />
URL: [https://ftp.gnu.org/gnu/tar/tar-1.35.tar.xz](https://ftp.gnu.org/gnu/tar/tar-1.35.tar.xz)<br />

## Configuration

To configure GNU Tar for install into the build root, run the following command:

```bash
./configure --prefix=/usr       \
            --host=$BRFS_TARGET \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Tar, run the following command:

```bash
make
```

Finally, to install GNU Tar into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUSed.md) GNU Sed | [HOME](../README.md) | Xz [>>](./Xz.md) |

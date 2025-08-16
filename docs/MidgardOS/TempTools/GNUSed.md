# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPatch.md) GNU Patch | [HOME](../README.md) | GNU Tar [>>](./GNUTar.md) |

## GNU Sed

Name: GNU Sed<br />
Summary: The GNU Stream Editor<br />
License: GPL v3 or later<br />
Version: 4.9<br />
URL: [https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz](https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz)<br />

## Configuration

To configure GNU Sed for install into the build root, run the following command:

```bash
./configure --prefix=/usr       \
            --host=$BRFS_TARGET \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Sed, run the following command:

```bash
make
```

Finally, to install GNU Sed into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPatch.md) GNU Patch | [HOME](../README.md) | GNU Tar [>>](./GNUTar.md) |

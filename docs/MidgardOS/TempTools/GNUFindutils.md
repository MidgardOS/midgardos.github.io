# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./File.md) File | [HOME](../README.md) | GNU Awk [>>](./GAWK.md) |

## GNU Findutils

Name: GNU Findutils<br />
Summary: CLI tools for searching for files based on various criteria<br />
License: $PKG_LICENSE<br />
Version: 4.10.0<br />
URL: [https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz](https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz)<br />

## Configuration

To configure GNU Findutils for install into the build root, run the following command:

```bash
./configure --prefix=/usr                   \
            --localstatedir=/var/lib/locate \
            --host=$BRFS_TARGET             \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Findutils, run the following command:

```bash
make
```

Finally, to install GNU Findutils into the cross-tools tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./File.md) File | [HOME](../README.md) | GNU Awk [>>](./GAWK.md) |

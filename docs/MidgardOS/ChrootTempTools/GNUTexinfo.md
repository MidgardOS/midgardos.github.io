# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313.md) Python 3.13 | [HOME](../README.md) | Util-Linux 64-bit [>>](./UtilLinux64bit.md) |

## GNU Texinfo

Name: GNU Texinfo<br />
Summary: Tools for reading, writing, and converting info documents<br />
License: GPL v3 or later<br />
Version: 7.2<br />
URL: [https://ftp.gnu.org/gnu/texinfo/texinfo-7.2.tar.xz](https://ftp.gnu.org/gnu/texinfo/texinfo-7.2.tar.xz)<br />

## Configuration

To configure GNU Texinfo for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64
```

## Compilation and Installation

To compile GNU Texinfo, run the following command:

```bash
make
```

Finally, to install GNU Texinfo into the build tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313.md) Python 3.13 | [HOME](../README.md) | Util-Linux 64-bit [>>](./UtilLinux64bit.md) |

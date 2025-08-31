# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGettext.md) GNU Gettext | [HOME](../README.md) | Perl [>>](./Perl.md) |

## GNU Bison

Name: GNU Bison<br />
Summary: The GNU Parser Generator<br />
License: $PKG_LICENSE<br />
Version: 3.8.2<br />
URL: [https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz](https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz)<br />

## Configuration

To configure GNU Bison for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --docdir=/usr/share/doc/bison-3.8.2
```

## Compilation and Installation

To compile GNU Bison, run the following command:

```bash
make
```

Finally, to install GNU Bison into the build tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGettext.md) GNU Gettext | [HOME](../README.md) | Perl [>>](./Perl.md) |

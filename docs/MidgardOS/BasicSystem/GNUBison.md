# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGettext_32bit.md) GNU Gettext 32-bit | [HOME](../README.md) | GNU Grep [>>](./GNUGrep.md) |

## GNU Bison

Name: GNU Bison<br />
Summary: The GNU Parser Generator<br />
License: GPL v3<br />
Version: 3.8.2<br />
URL: [https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz](https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz)<br />

Average Build Time: 2.1 SBU<br />
Used Install Space: 63 MiB<br />

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

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Bison into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | bison, yacc |
| Installed Libraries | liby.a |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGettext_32bit.md) GNU Gettext 32-bit | [HOME](../README.md) | GNU Grep [>>](./GNUGrep.md) |

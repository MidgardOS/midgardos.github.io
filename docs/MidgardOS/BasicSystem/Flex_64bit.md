# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBc.md) GNU Bc | [HOME](../README.md) | Flex 32-bit [>>](./Flex_32bit.md) |

## Flex 64-bit

Name: Flex 64-bit<br />
Summary: The fast lexer<br />
License: BSD-3-Clause<br />
Version: 2.6.4<br />
URL: [https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz](https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 33 MiB<br />

## Configuration

To configure Flex 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                      \
            --libdir=/usr/lib64                \
            --libexecdir=/usr/lib64            \
            --docdir=/usr/share/doc/flex-2.6.4 \
            --disable-static
```

## Compilation and Installation

To compile Flex 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install Flex 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libfl.la
ln -sv flex   /usr/bin/lex
ln -sv flex.1 /usr/share/man/man1/lex.1
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | flex, flex++ (link to flex), and lex (link to flex) |
| Installed Libraries | libfl.so |
| Installed Directories | /usr/share/doc/flex-2.6.4 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBc.md) GNU Bc | [HOME](../README.md) | Flex 32-bit [>>](./Flex_32bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUM4.md.md) GNU M4 | [HOME](../README.md) | Flex 64-bit [>>](./Flex64bit.md) |

## GNU Bc

Name: GNU Bc<br />
Summary: An arbitrary precision processing language<br />
License: GPL v3 or later<br />
Version: 1.08.2<br />
URL: [https://ftp.gnu.org/gnu/bc/bc-1.08.2.tar.lz](https://ftp.gnu.org/gnu/bc/bc-1.08.2.tar.lz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 7.8 MiB<br />

## Configuration

To configure GNU Bc for install into the build root, run the following command:

```bash
./configure --prefix=/usr
```

## Compilation and Installation

To compile GNU Bc, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Bc into the build tree, run the following command:

```bash
make install
```

| Contents ||
| --- | --- |
| Installed Programs | bc and dc |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUM4.md.md) GNU M4 | [HOME](../README.md) | Flex 64-bit [>>](./Flex64bit.md) |

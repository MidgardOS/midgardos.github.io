# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBison.md) GNU Bison | [HOME](../README.md) | GNU Bash [>>](./GNUBash.md) |

## GNU Grep

Name: GNU Grep<br />
Summary: CLI tool for searching through text files<br />
License: GPL v3 or later<br />
Version: 3.12<br />
URL: [https://ftp.gnu.org/gnu/grep/grep-3.12.tar.xz](https://ftp.gnu.org/gnu/grep/grep-3.12.tar.xz)<br />

Average Build Time: 0.6 SBU<br />
Used Install Space: 48 MiB<br />

## Configuration

To configure GNU Grep for install into the build root, run the following commands:

```bash
sed -i "s/echo/#echo/" src/egrep.sh
./configure --prefix=/usr --libdir=/usr/lib64 --libexecdir=/usr/lib64
```

## Compilation and Installation

To compile GNU Grep, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Grep into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | egrep, fgrep, grep |

| Navigation |||
| --- | --- | ---: |
| [<<](./$PREVIOUS_PAGE.md) $PREVIOUS_PKG_FULL_NAME | [HOME](../README.md) | $NEXT_PKG_FULL_NAME [>>](./$NEXT_PAGE.md) |

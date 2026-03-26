# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Pytest.md) Python module Pytest | [HOME](../README.md) | GTK-Doc [>>](./GTK-Doc.md) |

## GNU Which

Name: GNU Which<br />
Summary: A tool for finding the fully-qualified path to an installed binary command<br />
License: GPL v3<br />
Version: 2.23<br />
URL: [https://ftp.gnu.org/gnu/which/which-2.23.tar.gz](https://ftp.gnu.org/gnu/which/which-2.23.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.2 MiB<br />

## Configuration

To configure GNU Which for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64
```

## Compilation and Installation

To compile GNU Which, run the following command:

```bash
make
```

This package does not contain a test suite.

Finally, to install GNU Which into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | which |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Pytest.md) Python module Pytest | [HOME](../README.md) | GTK-Doc [>>](./GTK-Doc.md) |

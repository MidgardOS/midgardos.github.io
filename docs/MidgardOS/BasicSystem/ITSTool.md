# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-LXML.md) Python module LXML | [HOME](../README.md) | Python module Pygments [>>](./Python313-Pygments.md) |

## ITS Tool

Name: ITS Tool<br />
Summary: A tool that converts XML documents to PO and MO to create translated XML files using the W3C ITS.<br />
License: GPL v3.0 or later<br />
Version: 2.0.7<br />
URL: [https://github.com/itstool/itstool/archive/refs/tags/2.0.7.tar.gz](https://github.com/itstool/itstool/archive/refs/tags/2.0.7.tar.gz)<br />

Average Build Time: SBU<br />
Used Install Space: <br />

## Compilation and Installation

To compile $PKG_FULL_NAME, run the following command:

```bash
patch -Np1 -i ../patches/itstool/itstool-2.0.7-regular-expressions-quoting-bug.patch
patch -Np1 -i ../patches/itstool/itstool-2.0.7-lxml-1.patch
PYTHON=/usr/bin/python3 ./autogen.sh --prefix=/usr
make
```

Finally, to install $PKG_FULL_NAME into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | itstool |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-LXML.md) Python module LXML | [HOME](../README.md) | Python module Pygments [>>](./Python313-Pygments.md) |

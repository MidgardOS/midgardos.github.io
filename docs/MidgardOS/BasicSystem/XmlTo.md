# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./W3m.md) W3M | [HOME](../README.md) | SE CilC [>>](./SECilC.md) |

## XmlTo

Name: XmlTo<br />
Summary: A script to convert XML to various formats<br />
License: GPL v2<br />
Version: 0.0.29<br />
URL: [https://pagure.io/xmlto/archive/0.0.29/xmlto-0.0.29.tar.gz](https://pagure.io/xmlto/archive/0.0.29/xmlto-0.0.29.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 244 KiB<br />

## Configuration

To configure XmlTo for install into the build root, run the following commands:

```bash
rm -f xmlif/xmlif.c
autoreconf -fiv
export XML_CATALOG_FILES="/etc/xml/catalog"
XMLTO_BASH_PATH=/bin/bash \
./configure --prefix=/usr \
            --with-webbrowser=w3m
```

## Compilation and Installation

To compile XmlTo, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install XmlTo into the build tree, run the following command:

```bash
make install
unset XML_CATALOG_FILES
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | xmlif and xmlto |

| Navigation |||
| --- | --- | ---: |
| [<<](./W3m.md) W3M | [HOME](../README.md) | SE CilC [>>](./SECilC.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXSLT_32bit.md) LibXSLT 32-bit | [HOME](../README.md) | Docbook 3.1 DTD [>>](./Docbook31-DTD.md) |

## SGML Common

Name: SGML Common<br />
Summary: Basic tools for managing SGML catalogues<br />
License: GPL v2 or later<br />
Version: 0.6.3<br />
URL: [https://sourceware.org/ftp/docbook-tools/new-trials/SOURCES/sgml-common-0.6.3.tgz](https://sourceware.org/ftp/docbook-tools/new-trials/SOURCES/sgml-common-0.6.3.tgz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 384 KiB<br />

## Preparation

The SGML Common package needs a patch to add the manpage to the installation for the `install-catalog` script. To apply it, run the following command:

```bash
patch -Np1 -i  ../patches/sgml-common/sgml-common-0.6.3-manpage-fix.patch
autoreconf -fi
```

## Configuration

To configure SGML Common for install into the build root, run the following command:

```bash
./configure --prefix=/usr --sysconfdir=/etc
```

## Compilation and Installation

To compile SGML Common, run the following command:

```bash
make
```

This package does not come with a test suite.

Finally, to install SGML Common into the build tree, run the following command:

```bash
make docdir=/usr/share/doc install
install-catalog --add /etc/sgml/sgml-ent.cat \
    /usr/share/sgml/sgml-iso-entities-8879.1986/catalog &&
install-catalog --add /etc/sgml/sgml-docbook.cat \
    /etc/sgml/sgml-ent.cat
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | install-catalog and sgmlwhich |
| Installed Files | /etc/sgml/sgml.conf |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibXSLT_32bit.md) LibXSLT 32-bit | [HOME](../README.md) | Docbook 3.1 DTD [>>](./Docbook31-DTD.md) |

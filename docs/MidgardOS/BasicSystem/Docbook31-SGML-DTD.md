# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibArchive_32bit.md) LibArchive 32-bit | [HOME](../README.md) | Docbook 4.5 SGML DTD [>>](./Docbook45-SGML-DTD.md) |

## Docbook 3.1 SGML DTDs

Name: Docbook 3.1 SGML DTDs<br />
Summary: The document definitions of SGML data<br />
License: BSD 3-clause and MIT<br />
Version: 3.1<br />
URL: [https://www.docbook.org/sgml/3.1/docbk31.zip](https://www.docbook.org/sgml/3.1/docbk31.zip)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 626 KiB<br />

## Preparation

This package is distributed as a PkZip file. Unlike other packages in the system that come as tar.gz, tar.bz2. or tar.xz, the PkZip format does not create a directory for the extracted files to be uncompressed into. To ensure that the `/sources` tree remains clean, run the following commands to extract the content:

```bash
mkdir docbk31
cd docbk31
cp ../docbk31.zip ./
unzip docbk31.zip
```

Once extracted, remain in this directory until this package is installed.

## Installation

This package contains no configuration script, nor requires any compilation.

To install Docbook 3.1 SGML DTDs into the build tree, run the following commands:

```bash
sed -i -e '/ISO 8879/d' -e 's|DTDDECL "-//OASIS//DTD DocBook V3.1//EN"|SGMLDECL|g' docbook.cat
install -v -d -m755 -o root -g root /usr/share/sgml/docbook/sgml-dtd-3.1
install -v -m644 -o root -g root docbook.cat /usr/share/sgml/docbook/sgml-dtd-3.1/catalog
for FILE in *.dtd *.mod *.dcl; do
    install -v -m 644 -o root -g root $FILE /usr/share/sgml/docbook/sgml-dtd-3.1/
done
install-catalog --add /etc/sgml/sgml-docbook-dtd-3.1.cat /usr/share/sgml/docbook/sgml-dtd-3.1/catalog
install-catalog --add /etc/sgml/sgml-docbook-dtd-3.1.cat /etc/sgml/sgml-docbook.cat
cat >> /usr/share/sgml/docbook/sgml-dtd-3.1/catalog << "EOF"
  -- Begin Single Major Version catalog changes --

PUBLIC "-//Davenport//DTD DocBook V3.0//EN" "docbook.dtd"

  -- End Single Major Version catalog changes --
EOF
```

## Contents

| Contents | |
| --- | --- |
| Installed Files | SGML DTD and MOD files |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibArchive_32bit.md) LibArchive 32-bit | [HOME](../README.md) | Docbook 4.5 SGML DTD [>>](./Docbook45-SGML-DTD.md) |

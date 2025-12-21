# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook45-SGML-DTD.md) Docbook 4.5 SGML DTDs | [HOME](../README.md) | Docbook 5.0 XML DTDs [>>](./Docbook50-XML-DTD.md) |

## Docbook 4.5 XML DTDs

Name: Docbook 4.5 XML DTDs<br />
Summary: The document definitions for XML data<br />
License: BSD 3-clause and MIT<br />
Version: 4.5<br />
URL: [https://www.docbook.org/xml/4.5/docbook-xml-4.5.zip](https://www.docbook.org/xml/4.5/docbook-xml-4.5.zip)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 626 KiB<br />

## Preparation

This package is distributed as a PkZip file. Unlike other packages in the system that come as tar.gz, tar.bz2. or tar.xz, the PkZip format does not create a directory for the extracted files to be uncompressed into. To ensure that the `/sources` tree remains clean, run the following commands to extract the content:

```bash
mkdir docbook-xml-4.5
cd docbook-xml-4.5
cp ../docbook-xml-4.5.zip ./
unzip docbook-xml-4.5.zip
```

Once extracted, remain in this directory until this package is installed.

## Installation

This package contains no configuration script, nor requires any compilation.

To install Docbook 4.5 XML DTDs into the build tree, run the following commands:

```bash
install -v -d -m755 -o root -g root /usr/share/xml/docbook/xml-dtd-4.5
install -v -d -m755 -o root -g root /etc/xml
cp -v -af --no-preserve=ownership docbook.cat *.dtd ent/ *.mod /usr/share/xml/docbook/xml-dtd-4.5

if [[ ! -e /etc/xml/docbook ]]; then
    xmlcatalog --noout --create /etc/xml/docbook
fi

xmlcatalog --noout --add "public" "-//OASIS//DTD DocBook XML V4.5//EN"                                  \
    "http://www.oasis-open.org/docbook/xml/4.5/docbookx.dtd" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//DTD DocBook XML CALS Table Model V4.5//EN"                 \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/calstblx.dtd" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//DTD XML Exchange Table Model 19990315//EN"                 \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/soextblx.dtd" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//ELEMENTS DocBook XML Information Pool V4.5//EN"            \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/dbpoolx.mod" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//ELEMENTS DocBook XML Document Hierarchy V4.5//EN"          \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/dbhierx.mod" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//ELEMENTS DocBook XML HTML Tables V4.5//EN"                 \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/htmltblx.mod" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//ENTITIES DocBook XML Notations V4.5//EN"                   \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/dbnotnx.mod" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//ENTITIES DocBook XML Character Entities V4.5//EN"          \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/dbcentx.mod" /etc/xml/docbook
xmlcatalog --noout --add "public" "-//OASIS//ENTITIES DocBook XML Additional General Entities V4.5//EN" \
    "file:///usr/share/xml/docbook/xml-dtd-4.5/dbgenent.mod" /etc/xml/docbook
xmlcatalog --noout --add "rewriteSystem" "http://www.oasis-open.org/docbook/xml/4.5"                    \
    "file:///usr/share/xml/docbook/xml-dtd-4.5" /etc/xml/docbook
xmlcatalog --noout --add "rewriteURI" "http://www.oasis-open.org/docbook/xml/4.5"                       \
    "file:///usr/share/xml/docbook/xml-dtd-4.5" /etc/xml/docbook

if [[ ! -e /etc/xml/catalog ]]; then
    xmlcatalog --noout --create /etc/xml/catalog
fi

xmlcatalog --noout --add "delegatePublic" "-//OASIS//ENTITIES DocBook XML"      \
    "file:///etc/xml/docbook" /etc/xml/catalog
xmlcatalog --noout --add "delegatePublic" "-//OASIS//DTD DocBook XML"           \
    "file:///etc/xml/docbook" /etc/xml/catalog
xmlcatalog --noout --add "delegateSystem" "http://www.oasis-open.org/docbook/"  \
    "file:///etc/xml/docbook" /etc/xml/catalog
xmlcatalog --noout --add "delegateURI" "http://www.oasis-open.org/docbook/"     \
    "file:///etc/xml/docbook" /etc/xml/catalog

for DTDVERSION in "4.1.2" "4.2" "4.3" "4.4"; do
  xmlcatalog --noout --add "public" "-//OASIS//DTD DocBook XML V$DTDVERSION//EN"                    \
    "http://www.oasis-open.org/docbook/xml/$DTDVERSION/docbookx.dtd" /etc/xml/docbook

  xmlcatalog --noout --add "rewriteSystem" "http://www.oasis-open.org/docbook/xml/$DTDVERSION"      \
    "file:///usr/share/xml/docbook/xml-dtd-4.5" /etc/xml/docbook

  xmlcatalog --noout --add "rewriteURI" "http://www.oasis-open.org/docbook/xml/$DTDVERSION"         \
    "file:///usr/share/xml/docbook/xml-dtd-4.5" /etc/xml/docbook

  xmlcatalog --noout --add "delegateSystem" "http://www.oasis-open.org/docbook/xml/$DTDVERSION/"    \
    "file:///etc/xml/docbook" /etc/xml/catalog

  xmlcatalog --noout --add "delegateURI" "http://www.oasis-open.org/docbook/xml/$DTDVERSION/"       \
    "file:///etc/xml/docbook" /etc/xml/catalog
done
```

## Contents

| Contents | |
| --- | --- |
| Installed Files | XML DTD, ENT, and MOD files |

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook45-SGML-DTD.md) Docbook 4.5 SGML DTDs | [HOME](../README.md) | Docbook 5.0 XML DTDs [>>](./Docbook50-XML-DTD.md) |

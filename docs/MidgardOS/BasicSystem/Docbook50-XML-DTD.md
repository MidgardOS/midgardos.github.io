# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook45-XML-DTD.md) Docbook 4.5 XML DTDs | [HOME](../README.md) | Docbook XSL Stylesheets [>>](./Docbook-XSL-Stylesheets.md) |

## Docbook 5.0 XML DTDs

Name: Docbook 5.0 XML DTDs<br />
Summary: The document definitions for XML data<br />
License: BSD 3-clause and MIT<br />
Version: 5.0<br />
URL: [https://www.docbook.org/xml/5.0/docbook-5.0.zip](https://www.docbook.org/xml/5.0/docbook-5.0.zip)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 626 KiB<br />

## Preparation

This package is distributed as a PkZip file. Unlike other packages in the system that come as tar.gz, tar.bz2. or tar.xz, the PkZip format does not create a directory for the extracted files to be uncompressed into. To ensure that the `/sources` tree remains clean, run the following commands to extract the content:

```bash
mkdir docbook-5.0
cd docbook-5.0
cp ../docbook-5.0.zip ./
unzip docbook-5.0.zip
```

Once extracted, remain in this directory until this package is installed.

## Installation

This package contains no configuration script, nor requires any compilation.

To install Docbook 5.0 XML DTDs into the build tree, run the following commands:

```bash
install -v -d -m755 -o root -g root /usr/share/xml/docbook/schema/{dtd,rng,sch,xsd}/5.0
install -v -m644 -o root -g root dtd/* /usr/share/xml/docbook/schema/dtd/5.0
install -v -m644 -o root -g root rng/* /usr/share/xml/docbook/schema/rng/5.0
install -v -m644 -o root -g root sch/* /usr/share/xml/docbook/schema/sch/5.0
install -v -m644 -o root -g root xsd/* /usr/share/xml/docbook/schema/xsd/5.0

xmlcatalog --noout --create /usr/share/xml/docbook/schema/dtd/5.0/catalog.xml
xmlcatalog --noout --add "public" "-//OASIS//DTD DocBook XML 5.0//EN"                           \
  "docbook.dtd" /usr/share/xml/docbook/schema/dtd/5.0/catalog.xml
xmlcatalog --noout --add "system" "http://www.oasis-open.org/docbook/xml/5.0/dtd/docbook.dtd"   \
  "docbook.dtd" /usr/share/xml/docbook/schema/dtd/5.0/catalog.xml

xmlcatalog --noout --create /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/rng/docbook.rng"                     \
  "docbook.rng" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/rng/docbook.rng"      \
  "docbook.rng" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/rng/docbookxi.rng"                   \
  "docbookxi.rng" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/rng/docbookxi.rng"    \
  "docbookxi.rng" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/rng/docbook.rnc"                     \
  "docbook.rnc" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/rng/docbook.rnc"      \
  "docbook.rnc" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/rng/docbookxi.rnc"                   \
  "docbookxi.rnc" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/rng/docbookxi.rnc"    \
  "docbookxi.rnc" /usr/share/xml/docbook/schema/rng/5.0/catalog.xml

xmlcatalog --noout --create /usr/share/xml/docbook/schema/sch/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/sch/docbook.sch"                     \
  "docbook.sch" /usr/share/xml/docbook/schema/sch/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/sch/docbook.sch"      \
  "docbook.sch" /usr/share/xml/docbook/schema/sch/5.0/catalog.xml

xmlcatalog --noout --create /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/xsd/docbook.xsd"                     \
  "docbook.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/xsd/docbook.xsd"      \
  "docbook.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/xsd/docbookxi.xsd"                   \
  "docbookxi.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/xsd/docbookxi.xsd"    \
  "docbookxi.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/xsd/xlink.xsd"                       \
  "xlink.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/xsd/xlink.xsd"        \
   "xlink.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://docbook.org/xml/5.0/xsd/xml.xsd"                         \
   "xml.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml
xmlcatalog --noout --add "uri" "http://www.oasis-open.org/docbook/xml/5.0/xsd/xml.xsd"          \
   "xml.xsd" /usr/share/xml/docbook/schema/xsd/5.0/catalog.xml

if [[ ! -e /etc/xml/catalog ]]; then
    install -v -d -m755 -o root -g root /etc/xml
    xmlcatalog --noout --create /etc/xml/catalog
fi

xmlcatalog --noout --add "delegatePublic" "-//OASIS//DTD DocBook XML 5.0//EN" \
  "file:///usr/share/xml/docbook/schema/dtd/5.0/catalog.xml" /etc/xml/catalog
xmlcatalog --noout --add "delegateSystem" "http://docbook.org/xml/5.0/dtd/"   \
  "file:///usr/share/xml/docbook/schema/dtd/5.0/catalog.xml" /etc/xml/catalog
xmlcatalog --noout --add "delegateURI" "http://docbook.org/xml/5.0/dtd/"      \
  "file:///usr/share/xml/docbook/schema/dtd/5.0/catalog.xml" /etc/xml/catalog
xmlcatalog --noout --add "delegateURI" "http://docbook.org/xml/5.0/rng/"      \
  "file:///usr/share/xml/docbook/schema/rng/5.0/catalog.xml" /etc/xml/catalog
xmlcatalog --noout --add "delegateURI" "http://docbook.org/xml/5.0/sch/"      \
  "file:///usr/share/xml/docbook/schema/sch/5.0/catalog.xml" /etc/xml/catalog
xmlcatalog --noout --add "delegateURI" "http://docbook.org/xml/5.0/xsd/"      \
  "file:///usr/share/xml/docbook/schema/xsd/5.0/catalog.xml" /etc/xml/catalog
```

## Contents

| Contents | |
| --- | --- |
| Installed Files | XML DTD, RNC, RNG, SCH, and XSD files |

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook45-XML-DTD.md) Docbook 4.5 XML DTDs | [HOME](../README.md) | Docbook XSL Stylesheets [>>](./Docbook-XSL-Stylesheets.md) |

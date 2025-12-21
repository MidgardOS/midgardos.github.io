# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook50-XML-DTD.md) Docbook 5.0 XML DTDs | [HOME](../README.md) |  [>>](./) |

## Docbook XSL Stylesheets

Name: Docbook XSL Stylesheets<br />
Summary: The XSL Stylesheets for tranforming XML data<br />
License: BSD 3-clause and MIT<br />
Version: 1.79.2<br />
URL 1: [https://github.com/docbook/xslt10-stylesheets/releases/download/release/1.79.2/docbook-xsl-nons-1.79.2.tar.bz2](https://github.com/docbook/xslt10-stylesheets/releases/download/release/1.79.2/docbook-xsl-nons-1.79.2.tar.bz2)<br />
URL 2: [https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-1.79.2.tar.bz2](https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-1.79.2.tar.bz2)<br />
URL 3: [https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-doc-1.79.2.tar.bz2](https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-doc-1.79.2.tar.bz2)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 52 MiB<br />

## Preparation for the Non-namespaced XSL Stylesheets

To fix an issue where a stack overflow can occur when processing long strings due to deep recursion, apply the following patch using the below command while in the unpacked directory of Docbook XSL NoNS:

```bash
patch -Np1 -i ../patches/docbook-xsl-stylesheets/docbook-xsl-nons-1.79.2-stack_fix-1.patch
```

Now, unpack the documentation:

```bash
tar -xf ../docbook-xsl-doc-1.79.2.tar.bz2 --strip-components=1
```

## Installation

This package contains no configuration script, nor requires any compilation. The test suite is unable to be tested at this time due to missing dependencies in the system root. The test suite will be enabled later after the host's packages are built as RPMs.

To install Docbook XSL Stylesheets without namespacing into the build tree, run the following commands:

```bash
install -v -m755 -d /usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2
cp -Rv VERSION assembly common eclipse epub epub3 extensions fo        \
       highlighting html htmlhelp images javahelp lib manpages params  \
       profiling roundtrip slides template tests tools webhelp website \
       xhtml xhtml-1_1 xhtml5                                          \
        /usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2
ln -sv VERSION /usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2/VERSION.xsl
install -v -m644 -D README /usr/share/doc/docbook-xsl-nons-1.79.2/README.txt
install -v -m644    RELEASE-NOTES* NEWS* /usr/share/doc/docbook-xsl-nons-1.79.2
install -d -v -m 755 -o root -g root /usr/share/doc/docbook-xsl-nons-1.79.2
cp -Rv doc/* /usr/share/doc/docbook-xsl-nons-1.79.2

if [[ ! -d /etc/xml ]]; then
    install -v -m755 -d -o root -g root /etc/xml
fi
if [[ ! -f /etc/xml/catalog ]]; then
    xmlcatalog --noout --create /etc/xml/catalog
fi

xmlcatalog --noout --add "rewriteSystem" "http://cdn.docbook.org/release/xsl-nons/1.79.2"       \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "https://cdn.docbook.org/release/xsl-nons/1.79.2"      \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "http://cdn.docbook.org/release/xsl-nons/1.79.2"          \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "https://cdn.docbook.org/release/xsl-nons/1.79.2"         \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "http://cdn.docbook.org/release/xsl-nons/current"      \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "https://cdn.docbook.org/release/xsl-nons/current"     \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "http://cdn.docbook.org/release/xsl-nons/current"         \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "https://cdn.docbook.org/release/xsl-nons/current"        \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "http://docbook.sourceforge.net/release/xsl/current"   \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "http://docbook.sourceforge.net/release/xsl/current"      \
           "/usr/share/xml/docbook/xsl-stylesheets-nons-1.79.2" /etc/xml/catalog
```

Now that the non-namespaced XSL stylesheets are installed, install the namespaced ones. First, unpack the namespaced XSL stylesheets into their own directory and then apply the same patch as used above

```bash
install -v -d -m755 -o root -g root /usr/share/xml/docbook/xsl-stylesheets-1.79.2
cp -Rv VERSION assembly common eclipse epub epub3 extensions fo         \
        highlighting html htmlhelp images javahelp lib manpages params  \
        profiling roundtrip slides template tests tools webhelp website \
        xhtml xhtml-1_1 xhtml5                                          \
        /usr/share/xml/docbook/xsl-stylesheets-1.79.2
ln -sv VERSION /usr/share/xml/docbook/xsl-stylesheets-1.79.2/VERSION.xsl

if [[ ! -d /etc/xml ]]; then
    install -v -m755 -d -o root -g root /etc/xml;
fi
if [[ ! -f /etc/xml/catalog ]]; then
    xmlcatalog --noout --create /etc/xml/catalog
fi

xmlcatalog --noout --add "rewriteSystem" "http://cdn.docbook.org/release/xsl/1.79.2"                \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "https://cdn.docbook.org/release/xsl/1.79.2"               \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "http://cdn.docbook.org/release/xsl/1.79.2"                   \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "https://cdn.docbook.org/release/xsl/1.79.2"                  \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "http://cdn.docbook.org/release/xsl/current"               \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "https://cdn.docbook.org/release/xsl/current"              \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "http://cdn.docbook.org/release/xsl/current"                  \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "https://cdn.docbook.org/release/xsl/current"                 \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteSystem" "http://docbook.sourceforge.net/release/xsl-ns/current"    \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
xmlcatalog --noout --add "rewriteURI" "http://docbook.sourceforge.net/release/xsl-ns/current"       \
           "/usr/share/xml/docbook/xsl-stylesheets-1.79.2" /etc/xml/catalog
```

## Contents

| Contents | |
| --- | --- |
| Installed Files | XSL stylesheets |

| Navigation |||
| --- | --- | ---: |
| [<<](./Docbook50-XML-DTD.md) Docbook 5.0 XML DTDs | [HOME](../README.md) |  [>>](./) |

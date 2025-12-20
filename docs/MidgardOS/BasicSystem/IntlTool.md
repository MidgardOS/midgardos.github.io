# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl-XML-Parser.md) Perl-XML-Parser | [HOME](../README.md) | OpenSSL v3 64-bit [>>](./OpenSSLv3_64bit.md) |

## IntlTool

Name: IntlTool<br />
Summary: A tool to convert and various i18n strings from files<br />
License: GPL v2<br />
Version: 0.51.0<br />
URL: [https://launchpad.net/intltool/trunk/0.51.0/+download/intltool-0.51.0.tar.gz](https://launchpad.net/intltool/trunk/0.51.0/+download/intltool-0.51.0.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 168 KiB<br />

## Configuration

To configure IntlTool for install into the build root, run the following commands:

```bash
sed -i 's:\\\${:\\\$\\{:' intltool-update.in
./configure --prefix=/usr
```

## Compilation and Installation

To compile IntlTool, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install IntlTool into the build tree, run the following commands:

```bash
make install
install -v -Dm 644 doc/I18N-HOWTO /usr/share/doc/intltool-0.51.0/I18N-HOWTO
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | intltool-extract, intltool-merge, intltool-prepare, intltool-update, and intltoolize |

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl-XML-Parser.md) Perl-XML-Parser | [HOME](../README.md) | OpenSSL v3 64-bit [>>](./OpenSSLv3_64bit.md) |

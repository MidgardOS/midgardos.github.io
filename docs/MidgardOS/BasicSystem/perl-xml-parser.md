# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./perl.md) Perl | [HOME](../README.md) | IntlTool [>>](./intltool.md) |

## perl-XML-Parser

Name: perl-XML-Parser<br />
Summary: An XML parser for Perl based on the Expat library<br />
License: Artistic<br />
Version: 2.47<br />
URL: [https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.tar.gz](https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 512 KiB<br />

## Configuration

To configure perl-XML-Parser for install into the build root, run the following command:

```bash
perl Makefile.PL
```

## Compilation and Installation

To compile perl-XML-Parser, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make test
```

Finally, to install perl-XML-Parser into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | Expat.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./perl.md) Perl | [HOME](../README.md) | IntlTool [>>](./intltool.md) |

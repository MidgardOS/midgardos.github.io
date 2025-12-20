# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | IntlTool [>>](./IntlTool.md) |

## Perl-XML-Parser

Name: Perl-XML-Parser<br />
Summary: An XML parser for Perl based on the Expat library<br />
License: Artistic<br />
Version: 2.47<br />
URL: [https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.tar.gz](https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 512 KiB<br />

## Configuration

To configure the Perl XML-Parser for install into the build root, run the following command:

```bash
perl Makefile.PL
```

## Compilation and Installation

To compile the Perl XML-Parser, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make test
```

Finally, to install the Perl XML-Parser into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | Expat.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | IntlTool [>>](./IntlTool.md) |

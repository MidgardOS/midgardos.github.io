# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./PCRE2_32bit.md) PCRE2 32-bit | [HOME](../README.md) | GNU Bc [>>](./GNUBc.md) |

## GNU M4

Name: GNU M4<br />
Summary: A small macro processor runtime often used in Makefiles<br />
License: GPL v3<br />
Version: 1.4.20<br />
URL: [https://ftp.gnu.org/gnu/m4/m4-1.4.20.tar.xz](https://ftp.gnu.org/gnu/m4/m4-1.4.20.tar.xz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 60 MiB<br />

## Configuration

To configure GNU M4 for install into the build root, run the following commands:

```bash
sed 's/\[\[__nodiscard__]]//' -i lib/config.hin
sed 's/test-stdalign\$(EXEEXT) //' -i tests/Makefile.in
./configure --prefix=/usr
```

## Compilation and Installation

To compile GNU M4, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU M4 into the build tree, run the following command:

```bash
make install
```

| Contents ||
| --- | --- |
| Installed Programs | m4 |

| Navigation |||
| --- | --- | ---: |
| [<<](./PCRE2_32bit.md) PCRE2 32-bit | [HOME](../README.md) | GNU Bc [>>](./GNUBc.md) |

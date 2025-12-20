# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibtool_32bit.md) GNU Libtool 32-bit | [HOME](../README.md) | Audit 64-bit [>>](./Audit_64bit.md) |

# GNU Automake

Name: GNU Automake<br />
Summary: A tool for automatically generating GNU-style Makefile.in files<br />
License: GPL v3+<br />
Version: 1.18.1<br />
URL: [https://ftp.gnu.org/pub/gnu/automake/automake-1.18.1.tar.xz](https://ftp.gnu.org/pub/gnu/automake/automake-1.18.1.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.5 MiB<br />

## Configuration

To configure GNU Automake for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --sysconfdir=/etc       \
            --libexecdir=/usr/lib64
```

## Compilation and Installation

To compile GNU Automake, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

There are 29 tests that will fail with this release.

Finally, to install GNU Automake into the cross-tools tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | aclocal, aclocal-1.18, automake, and automake-1.18 |
| Installed Files | /usr/share/aclocal/, /usr/share/aclocal-1.18, /usr/share/automake-1.18/ |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibtool_32bit.md) GNU Libtool 32-bit | [HOME](../README.md) | Audit 64-bit [>>](./Audit_64bit.md) |

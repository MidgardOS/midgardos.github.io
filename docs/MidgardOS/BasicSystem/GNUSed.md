# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses_32bit.md) NCurses 32-bit | [HOME](../README.md) | PSMisc [>>](./PSMisc.md) |

## GNU Sed

Name: GNU Sed<br />
Summary: The GNU Stream Editor<br />
License: GPL v3<br />
Version: 4.9<br />
URL: [https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz](https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 30 MiB<br />

## Configuration

To configure GNU Sed for install into the build root, run the following command:

```bash
./configure --prefix=/usr
```

## Compilation and Installation

To compile GNU Sed, run the following commands:

```bash
make
make html
```

Now, run the test suite:

```bash
make check
```

Finally, to install GNU Sed into the build tree, run the following command:

```bash
make install
install -d -m755           /usr/share/doc/sed-4.9
install -m644 doc/sed.html /usr/share/doc/sed-4.9
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | sed |

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses_32bit.md) NCurses 32-bit | [HOME](../README.md) | PSMisc [>>](./PSMisc.md) |

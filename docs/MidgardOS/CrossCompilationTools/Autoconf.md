| Navigation |||
| --- | --- | ---: |
| [<<](./NCursesTic.md) | [HOME](../README.md) | [>>](./PkgConf.md) |

# GNU Autoconf

Name: GNU Autoconf<br />
Summary: A tool for automatically configuring source code for build<br />
License: GPL v3+<br />
Version: 2.71<br />
URL: [http://ftp.gnu.org/pub/gnu/autoconf](http://ftp.gnu.org/pub/gnu/autoconf)<br />

## Configuration

To configure GNU Autoconf for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/cross-tools
```

## Compilation and Installation

To compile GNU Autoconf, run the following command:

```bash
make
```

Finally, to install GNU Autoconf into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./NCursesTic.md) | [HOME](../README.md) | [>>](./PkgConf.md) |

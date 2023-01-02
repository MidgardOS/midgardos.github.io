| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileLinuxHeaders.md) | [HOME](./README.md) | [>>](./CrossCompileNCursesTic.md) |

# GNU File

Name: GNU M4<br />
Summary: A macro processing language often used in Makefiles<br />
License: GPL v3<br />
Version: 1.4.19<br />
URL: [http://ftp.gnu.org/gnu/m4/](http://ftp.gnu.org/gnu/m4/)<br />

## Configuration

To configure GNU M4 for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/cross-tools
```

## Compilation and Installation

To compile GNU M4, run the following command:

```bash
make
```

Finally, to install GNU M4 into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileLinuxHeaders.md) | [HOME](./README.md) | [>>](./CrossCompileNCursesTic.md) |

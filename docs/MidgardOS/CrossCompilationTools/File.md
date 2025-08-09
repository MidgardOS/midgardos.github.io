| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib.md) | [HOME](../README.md) | [>>](./LinuxHeaders.md) |

# GNU File

Name: File<br />
Summary: A tool to determine the type of a file from heuristics<br />
License: BSD 2-Clause<br />
Version: 5.44<br />
URL: [http://ftp.astrom.com/pub/file/](ftp://ftp.astron.com/pub/file/)<br />

## Configuration

To configure File for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/cross-tools --libdir=/cross-tools/lib64 --enable-shared
```

## Compilation and Installation

To compile File, run the following command:

```bash
make
```

Finally, to install File into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib.md) | [HOME](../README.md) | [>>](./LinuxHeaders.md) |

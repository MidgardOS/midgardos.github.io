| Navigation |||
| --- | --- | ---: |
| [<<](./Overview.md) | [HOME](../README.md) | [>>](./File.md) |

# Z Lib

Name: zlib<br />
Summary: A fast compression library<br />
License: Zlib<br />
Version: 1.2.13<br />
URL: [http://www.zlib.net/](http://zlib.net/)<br />

## Configuration

To configure ZLib for installation into the cross-tools tree, run the following command:

```bash
./configure --shared --prefix=/cross-tools --libdir=/cross-tools/lib64
```

## Compilation and Installation

To compile zlib, run the following command:

```bash
make
```

Finally, to install zlib, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./Overview.md) | [HOME](../README.md) | [>>](./File.md) |

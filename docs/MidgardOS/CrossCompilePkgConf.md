| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileAutoconf.md) | [HOME](./README.md) | [>>](./CrossCompileGMP.md) |

# PkgConf

Name: PkgConf<br />
Summary: A software compilation and linker metadata tool and library<br />
License: ISC<br />
Version: 1.9.3<br />
URL: [https://github.com/pkgconf/pkgconf/archive/refs/tags](https://github.com/pkgconf/pkgconf/archive/refs/tags/)<br />

## Configuration

To configure PkgConf for install into our cross-compilation root, run the following command:

```bash
autoreconf -fiv
./configure --disable-static --prefix=/cross-tools
```

## Compilation and Installation

To compile PkgConf, run the following command:

```bash
make
```

Finally, to install PkgConf into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileAutoconf.md) | [HOME](./README.md) | [>>](./CrossCompileGMP.md) |

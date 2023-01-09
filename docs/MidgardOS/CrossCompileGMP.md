| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompilePkgConf.md) | [HOME](./README.md) | [>>](./CrossCompileMPFR.md) |

# GNU Math Procision Library

Name: GMP<br />
Summary: A library for working with large numbers<br />
License: GPL v3+<br />
Version: 6.2.1<br />
URL: [http://ftp.gnu.org/pub/gnu/gmp](http://ftp.gnu.org/pub/gnu/gmp)<br />

## Configuration

To configure GMP for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/cross-tools --disable-static --enable-cxx --enable-fat
```

## Compilation and Installation

To compile GMP, run the following command:

```bash
make
```

Finally, to install GMP into the cross-tools tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompilePkgConf.md) | [HOME](./README.md) | [>>](./CrossCompileMPFR.md) |

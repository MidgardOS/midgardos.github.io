| Navigation |||
| --- | --- | ---: |
| [<<](./PkgConf.md) | [HOME](../README.md) | [>>](./MPFR.md) |

# GNU Math Precision Library

Name: GMP<br />
Summary: A library for working with large numbers<br />
License: GPL v3+<br />
Version: 6.2.1<br />
URL: [http://ftp.gnu.org/pub/gnu/gmp](http://ftp.gnu.org/pub/gnu/gmp)<br />

## Configuration

To configure GMP for install into our cross-compilation root, run the following command:

The copying of certain components in the source tree is to remove undesired optimizations.

```bash
cp configfsf.guess config.guess
cp configfsf.sub config.sub
./configure --prefix=/cross-tools --libdir=/cross-tools/lib64 --disable-static --enable-cxx
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
| [<<](./PkgConf.md) | [HOME](../README.md) | [>>](./MPFR.md) |

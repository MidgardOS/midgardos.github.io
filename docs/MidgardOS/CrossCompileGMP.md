| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompilePkgConf.md) | [HOME](./README.md) | [>>](./CrossCompileMPFR.md) |

# GNU Math Procision Library

Name: GMP<br />
Summary: A for working with large numbers<br />
License: GPL v3+<br />
Version: 6.2.1<br />
URL: [http://ftp.gnu.org/pub/gnu/gmp](http://ftp.gnu.org/pub/gnu/gmp)<br />

## Configuration

To configure GMP for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/cross-tools --disable-static --enable-cxx --enable-fat \
            --host=zen2-pc-linux-gnu
```

The host flag is to ensure that the package will be built to run on x86_64 rev 2 or better systems. By default, GMP's configure script will probe the CPU and may build the library too optimized for later use.

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

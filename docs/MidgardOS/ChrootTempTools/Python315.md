# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | GNU Texinfo [>>](./GNUTexinfo.md) |

## Python 3.15

Name: Python 3.15<br />
Summary: The Python scripting language<br />
License: Python-2.0<br />
Version: 3.15.2<br />
URL: [https://www.python.org/ftp/python/3.15.2/Python-3.15.2.tar.xz](https://www.python.org/ftp/python/3.15.2/Python-3.15.2.tar.xz)<br />

## Configuration

To configure Python 3.15 for installation into the build root, run the following command:

```bash
./configure --prefix=/usr              \
            --libdir=/usr/lib64        \
            --libexecdir=/usr/lib64    \
            --enable-shared            \
            --with-platlibdir=lib64    \
            --without-ensurepip        \
            --without-static-libpython
```

## Compilation and Installation

To compile Python 3.15, run the following command:

```bash
make
```

Finally, to install Python 3.15 into the build tree, run the following command:

```bash
make install
```

More details about this package are covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | GNU Texinfo [>>](./GNUTexinfo.md) |

# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | GNU Texinfo [>>](./GNUTexinfo.md) |

## Python 3.13

Name: Python 3.13<br />
Summary: The Python scripting language<br />
License: Python-2.0<br />
Version: 3.13.7<br />
URL: [https://www.python.org/ftp/python/3.13.7/Python-3.13.7.tar.xz](https://www.python.org/ftp/python/3.13.7/Python-3.13.7.tar.xz)<br />

## Configuration

To configure Python 3.13 for install into the build root, run the following command:

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

To compile Python 3.13, run the following command:

```bash
make
```

Finally, to install Python 3.13 into the build tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | GNU Texinfo [>>](./GNUTexinfo.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./SQLite3_32bit.md) SQLite3 32-bit | [HOME](../README.md) | Python module Flit-Core [>>](./Python313-Flit-Core.md) |

## Python 3.13

Name: Python 3.13<br />
Summary: The Python scripting language<br />
License: Python-2.0<br />
Version: 3.13.7<br />
URL: [https://www.python.org/ftp/python/3.13.7/Python-3.13.7.tar.xz](https://www.python.org/ftp/python/3.13.7/Python-3.13.7.tar.xz)<br />

Average Build Time: 2.0 SBU<br />
Used Install Space: 337 MiB<br />

## Configuration

To configure Python 3.13 for install into the build root, run the following command:

```bash
./configure --prefix=/usr                                               \
            --libdir=/usr/lib64                                         \
            --libexecdir=/usr/lib64                                     \
            --enable-ipv6                                               \
            --with-system-expat                                         \
            --with-system-ffi                                           \
            --enable-optimizations                                      \
            --enable-shared                                             \
            --with-lto                                                  \
            --with-platlibdir=lib64                                     \
            --with-ssl-default-suites=openssl                           \
            --with-wheel-pkg-dir=/usr/lib64/python3.13/vendor-packages  \
            --enable-loadable-sqlite-extensions                         \
            --without-static-libpython
```

## Compilation and Installation

To compile Python 3.13, run the following command:

```bash
make
```

Next, run the test suite with environment variables to control timeouts of some tests that can hang indefinitely:\

```bash
make test TESTOPTS="--timeout 120"
```

Finally, to install Python 3.13 into the build tree, run the following commands:

```bash
make install
install -v -dm755 /usr/share/doc/python-3.13.7/html
tar --strip-components=1  \
    --no-same-owner       \
    --no-same-permissions \
    -C /usr/share/doc/python-3.13.7/html \
    -xvf ../python-3.13.7-docs-html.tar.bz2
```

To allow further Python modules to be installed during the pre-RPM builds and avoid unnecessary warnings and avoid future warnings around newer versions of `pip3`, run the following command to suppress these warnings:

```bash
cat > /etc/pip.conf << EOF
[global]
root-user-action = ignore
disable-pip-version-check = true
EOF
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | idle3, pip3, pydoc3, python3, and python3-config |
| Installed Libraries | libpython3.13.so and libpython3.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./SQLite3_32bit.md) SQLite3 32-bit | [HOME](../README.md) | Python module Flit-Core [>>](./Python313-Flit-Core.md) |

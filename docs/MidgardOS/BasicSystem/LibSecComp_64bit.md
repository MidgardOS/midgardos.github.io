# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Jinja2.md) Python module Jinja2 | [HOME](../README.md) | LibSecComp 32-bit [>>](./LibSecComp_32bit.md) |

## LibSecComp 64-bit

Name: LibSecComp 64-bit<br />
Summary: A platform independent interface to the Linux syscall filtering mechanism<br />
License: LGPL v2.1 or later<br />
Version: 2.6.0<br />
URL: [https://github.com/seccomp/libseccomp/releases/download/v2.6.0/libseccomp-2.6.0.tar.gz](https://github.com/seccomp/libseccomp/releases/download/v2.6.0/libseccomp-2.6.0.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 3.6 MiB<br />

## Configuration

To configure LibSecComp 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
            --docdir=/usr/share/doc/libseccomp-2.6.0    \
            --sysconfdir=/etc                           \
            --enable-pic                                \
            --enable-shared                             \
            --enable-python
```

Note, we cannot disable the build of the static library at configuration time since it is needed to build the Python components.

## Compilation and Installation

To compile LibSecComp 64-bit, run the following command:

```bash
make
```

The tests at this point will not run without `valgrind` and LibBPF. These tools will be installed later.

Finally, to install LibSecComp 64-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib64/*.a
rm -fv DESTDIR/usr/lib64/*.la
cp -Rv DESTDIR/usr/* /usr/
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | scmp_sys_resolver |
| Installed Libraries | libseccomp.so |
| Installed Plugins | seccomp Python module |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Jinja2.md) Python module Jinja2 | [HOME](../README.md) | LibSecComp 32-bit [>>](./LibSecComp_32bit.md) |

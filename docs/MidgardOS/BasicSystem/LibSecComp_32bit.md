# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSecComp_64bit.md) LibSecComp 64-bit | [HOME](../README.md) | JSON-C 64-bit [>>](./JSONC_64bit.md) |

## LibSecComp 32-bit

Name: LibSecComp 32-bit<br />
Summary: A platform independent interface to the Linux syscall filtering mechanism<br />
License: LGPL v2.1 or later<br />
Version: 2.6.0<br />
URL: [https://github.com/seccomp/libseccomp/releases/download/v2.6.0/libseccomp-2.6.0.tar.gz](https://github.com/seccomp/libseccomp/releases/download/v2.6.0/libseccomp-2.6.0.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 3.6 MiB<br />

## Configuration

To configure LibSecComp 32-bit for install into the build root, run the following command:

```bash
rm -rf DESTDIR
make distclean
./configure --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
            --docdir=/usr/share/doc/libseccomp-2.6.0    \
            --sysconfdir=/etc                           \
            --disable-static                            \
            --disable-python                            \
            --enable-pic                                \
            --enable-shared
```

## Compilation and Installation

To compile LibSecComp 32-bit, run the following command:

```bash
```

Finally, to install LibSecComp 32-bit into the build tree, run the following command:

```bash
```

## Contents

See the contents section of the 64-bit build of the LibSecComp 32-bit for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSecComp_64bit.md) LibSecComp 64-bit | [HOME](../README.md) | JSON-C 64-bit [>>](./JSONC_64bit.md) |

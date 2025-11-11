# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libsepol_64bit.md) LibSEPol 64-bit | [HOME](../README.md) | LibSELinux 64-bit [>>](./libselinux_64bit.md) |

## LibSEPol 32-bit

Name: LibSEPol 32-bit<br />
Summary: A library for manipulating SELinux binary policies<br />
License: LGPL v2.1<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/libsepol-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/libsepol-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 3.2 MiB<br />

## Configuration

This package has no traditional configuration script.

## Compilation and Installation

To compile LibSEPol 32-bit, run the following commands:

```bash
make clean
CC="gcc -m32 -march=i686" make
```

Finally, to install LibSEPol 32-bit into the build tree, run the following commands:

```bash
PREFIX=/usr LIBDIR=/usr/lib SHLIBDIR=/usr/lib DESTDIR=$PWD/DESTDIR make install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

Other libraries and programs of the suite of tools for SELinux require linking against the static library for LibSEPol, as such it is not removed.

## Contents

See the contents section of the 64-bit build of LibSEPol for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libsepol_64bit.md) LibSEPol 64-bit | [HOME](../README.md) | LibSELinux 64-bit [>>](./libselinux_64bit.md) |

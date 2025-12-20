# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSELinux_64bit.md) LibSELinux 64-bit | [HOME](../README.md) | LibSEManage 64-bit [>>](./LibSEManage_64bit.md) |

## LibSELinux 32-bit

Name: LibSELinux 32-bit<br />
Summary: SELinux runtime and utilities<br />
License: Public Domain<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/libselinux-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/libselinux-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.5 MiB<br />

## Configuration

The package does not have a traditional configuration script.

## Compilation and Installation

To compile LibSELinux 32-bit, run the following commands:

```bash
make clean
make CC="gcc -m32 -march=i686"
```

Finally, to install LibSELinux 32-bit into the build tree, run the following command:

```bash
PREFIX=/usr LIBDIR=/usr/lib SHLIBDIR=/usr/lib DESTDIR=$PWD/DESTDIR make install
```

Other libraries and programs of the suite of tools for SELinux require linking against the static library for LibSELinux, as such it is not removed.

## Contents

See the contents section of the 64-bit build of LibSELinux for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSELinux_64bit.md) LibSELinux 64-bit | [HOME](../README.md) | LibSEManage 64-bit [>>](./LibSEManage_64bit.md) |

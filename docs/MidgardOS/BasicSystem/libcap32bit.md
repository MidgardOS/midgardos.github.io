# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libcap64bit.md) libcap2 64-bit | [HOME](../README.md) | libxcrypt 64-bit [>>](./libxcrypt64bit.md) |

## libcap2 32-bit

Name: libcap2 32-bit<br />
Summary: Userspace tools for POSIX 1003.1e privilege capabilities<br />
License: BSD-3-Clause or GPL v2.0<br />
Version: 2.76<br />
URL: [https://www.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.76.tar.xz](https://www.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.76.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 3.1 MiB<br />

## Configuration

The package does not have a traditional configuration script.

## Compilation and Installation

To compile libcap2 32-bit, run the following commands:

```bash
make distclean
make CC="gcc -m32 -march=i686"
```

Finally, to install libcap2 32-bit into the build tree, run the following commands:

```bash
make CC="gcc -m32 -march=i686" lib=lib prefix=$PWD/DESTDIR/usr -C libcap install
cp -Rv DESTDIR/usr/lib/* /usr/lib
sed -e "s|^libdir=.*|libdir=/usr/lib|" -i /usr/lib/pkgconfig/lib{cap,psx}.pc
chmod -v 755 /usr/lib/libcap.so.2.76
```

## Contents

See the contents section of the 64-bit build of libcap2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./libcap64bit.md) libcap2 64-bit | [HOME](../README.md) | libxcrypt 64-bit [>>](./libxcrypt64bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib64bit.md) ZLib 64-bit | [HOME](../README.md) | BZip2 64-bit [>>](./BZip264bit.md) |

# ZLib 32-bit

Name: ZLib 32-bit<br />
Summary: A fast compression library<br />
License: Zlib<br />
Version: 1.3.1<br />
URL: [https://www.zlib.net/zlib-1.3.1.tar.xz](https://www.zlib.net/zlib-1.3.1.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 6.4 MiB<br />

## Configuration

To configure ZLib 32-bit for installation into the system, run the following command:

```bash
make distclean
CFLAGS+=" -m32" CXXFLAGS+=" -m32" \
./configure --prefix=/usr         \
            --libdir=/usr/lib     \
            --shared
```

## Compilation and Installation

To compile ZLib 32-bit, run the following command:

```bash
make
```

Now run the tests:

```bash
make check
```

Finally, to install ZLib 32-bit, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
rm -fv /usr/lib64/libz.a
```

## Contents

See the information for the 64-bit version for details about what is installed.

| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib64bit.md) ZLib 64-bit | [HOME](../README.md) | BZip2 64-bit [>>](./BZip264bit.md) |

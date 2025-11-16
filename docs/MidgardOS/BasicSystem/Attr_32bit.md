# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Attr_64bit.md) Attr 64-bit | [HOME](../README.md) | ACL 64-bit [>>](./ACL_64bit.md) |

## Attr 32-bit

Name: Attr 32-bit<br />
Summary: Utilities to administer the extended attributes of filesystem objects<br />
License: GPL v2 or later<br />
Version: 2.5.2<br />
URL: [https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.tar.xz](https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 4.1 MiB<br />

## Configuration

To configure Attr 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --sysconfdir=/etc        \
            --disable-static
```

## Compilation and Installation

To compile Attr 32-bit, run the following command:

```bash
make
```

Finally, to install Attr 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libattr.la
cp -Rv DESTDIR/usr/lib/* /usr/lib
```

## Contents

See the contents section of the 64-bit build of Attr for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./Attr_64bit.md) Attr 64-bit | [HOME](../README.md) | ACL 64-bit [>>](./ACL_64bit.md) |

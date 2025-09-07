# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ACL64bit.md) ACL 64-bit | [HOME](../README.md) | libcap 64-bit [>>](./libcap64bit.md) |

## ACL 32-bit

Name: ACL 32-bit<br />
Summary: Utilities to administer POSIX ACLs of filesystem objects<br />
License: GPL v2 or later<br />
Version: 2.3.2<br />
URL: [https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.tar.xz](https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 6.5 MiB<br />

## Configuration

To configure ACL 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32" ./configure \
    --prefix=/usr         \
    --disable-static      \
    --libdir=/usr/lib     \
    --libexecdir=/usr/lib \
    --host=i686-midgardos-linux-gnu
```

## Compilation and Installation

To compile ACL 32-bit, run the following command:

```bash
make
```

Finally, to install ACL 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libacl.la
cp -Rv DESTDIR/usr/lib/* /usr/lib
```

## Contents

See the contents section of the 64-bit build of ACL for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./ACL64bit.md) ACL 64-bit | [HOME](../README.md) | libcap 64-bit [>>](./libcap64bit.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibIDN2_32bit.md) GNU LibIDN2 32-bit | [HOME](../README.md) | Attr 32-bit [>>](./Attr_32bit.md) |

## Attr 64-bit

Name: Attr 64-bit<br />
Summary: Utilities to administer the extended attributes of filesystem objects<br />
License: GPL v2 or later<br />
Version: 2.5.2<br />
URL: [https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.tar.xz](https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 4.1 MiB<br />

## Configuration

To configure Attr 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --sysconfdir=/etc       \
            --docdir=/usr/share/doc/attr-2.5.2
```

## Compilation and Installation

To compile Attr 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

The `root/getattr.test` seems to fail if the host OS uses SELinux. For now, ignore this error, since it will be resolved later when SELinux functionality is added to the system.

Finally, to install Attr 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libattr.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | attr, getfattr, and setfattr |
| Installed Libraries | libattr.so |
| Installed Directories | /usr/include/attr and /usr/share/doc/attr-2.5.2 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibIDN2_32bit.md) GNU LibIDN2 32-bit | [HOME](../README.md) | Attr 32-bit [>>](./Attr_32bit.md) |

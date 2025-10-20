# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Attr32bit.md) Attr 32-bit | [HOME](../README.md) | ACL 32-bit [>>](./ACL32bit.md) |

## ACL 64-bit

Name: ACL 64-bit<br />
Summary: Utilities to administer POSIX ACLs of filesystem objects<br />
License: GPL v2 or later<br />
Version: 2.3.2<br />
URL: [https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.tar.xz](https://download.savannah.nongnu.org/releases/acl/acl-2.3.2.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 6.5 MiB<br />

## Configuration

To configure ACL 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --sysconfdir=/etc       \
            --disable-static        \
            --docdir=/usr/share/doc/acl-2.3.2
```

## Compilation and Installation

To compile ACL 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Several tests fail due to missing SELinux support if the host OS is SELinux aware, and other issues due to being run inside a chroot environment. For now, ignore.

Finally, to install ACL 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libacl.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | chacl, getfacl, and setfacl |
| Installed Libraries | libacl.so |
| Installed Directories | /usr/include/acl and /usr/share/doc/acl-2.3.2 |

| Navigation |||
| --- | --- | ---: |
| [<<](./Attr32bit.md) Attr 32-bit | [HOME](../README.md) | ACL 32-bit [>>](./ACL32bit.md) |

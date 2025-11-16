# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSELinux_32bit.md) LibSELinux 32-bit | [HOME](../README.md) | LibSEManage 32-bit [>>](./LibSEManage_32bit.md) |

## LibSEManage 64-bit

Name: LibSEManage 64-bit<br />
Summary: SELinux policy management library<br />
License: LGPL v2.1<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/libsemanage-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/libsemanage-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.7 MiB<br />

## Configuration

The package does not have a traditional configuration script.

## Compilation and Installation

To compile LibSEManage 64-bit, run the following commands:

```bash
grep /usr/libexec . -rl | xargs sed -i "s|/usr/libexec|/usr/lib64|g"
make
```

The tests for this package cannot be run at this time due to an unfulfilled dependency on `secilc`.

Finally, to install LibSEManage 64-bit into the build tree, run the following command:

```bash
PREFIX=/usr LIBDIR=/usr/lib64 LIBEXECDIR=/usr/lib64 SHLIBDIR=/usr/lib64 make install
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libsemanage.so |
| Installed Files | /etc/selinux/semanage.conf |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSELinux_32bit.md) LibSELinux 32-bit | [HOME](../README.md) | LibSEManage 32-bit [>>](./LibSEManage_32bit.md) |

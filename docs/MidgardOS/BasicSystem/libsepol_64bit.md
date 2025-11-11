# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libcunit_32bit.md) LibCunit 32-bit | [HOME](../README.md) | LibSEPol 32-bit [>>](./libsepol_32bit.md) |

## libsepol 64-bit

Name: libsepol 64-bit<br />
Summary: A library for manipulating SELinux binary policies<br />
License: LGPL v2.1<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/libsepol-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/libsepol-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 3.2 MiB<br />

## Configuration

This package has no traditional configuration script.

## Compilation and Installation

To compile libsepol 64-bit, run the following commands:

```bash
make
```

The test suite cannot be run at this time due to a circular dependency with `checkpolicy`.

Finally, to install libsepol 64-bit into the build tree, run the following commands:

```bash
PREFIX=/usr LIBDIR=/usr/lib64 SHLIBDIR=/usr/lib64 make install
```

Other libraries and programs of the suite of tools for SELinux require linking against the static library for LibSEPol, as such it is not removed.

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | chkcon, sepol_check_access, sepol_compute_av, sepol_compute_member, sepol_compute_relabel, and sepol_validate_transition |
| Installed Libraries | libsepol.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./libcunit_32bit.md) LibCunit 32-bit | [HOME](../README.md) | LibSEPol 32-bit [>>](./libsepol_32bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSEManage_32bit.md) LibSEManage 32-bit | [HOME](../README.md) | GNU LibGPG-Error 64-bit [>>](./GNULibGPG-Error_64bit.md) |

## CheckPolicy

Name: CheckPolicy<br />
Summary: SELinux policy compiler<br />
License: GPL v2 or later<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/checkpolicy-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/checkpolicy-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.5 MiB<br />

## Configuration

The package does not have a traditional configuration script.

## Compilation and Installation

To compile CheckPolicy, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make test
```

Finally, to install CheckPolicy into the build tree, run the following command:

```bash
PREFIX=/usr LIBDIR=/usr/lib64 SHLIBDIR=/usr/lib64 make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | checkmodule and checkpolicy |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSEManage_32bit.md) LibSEManage 32-bit | [HOME](../README.md) | GNU LibGPG-Error 64-bit [>>](./GNULibGPG-Error_64bit.md) |

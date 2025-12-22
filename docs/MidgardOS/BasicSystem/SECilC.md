# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./XmlTo.md) XmlTo | [HOME](../README.md) | LibEConf 64-bit [>>](./LibEConf_64bit.md) |

## SECilC

Name: SECilC<br />
Summary: SELinux Common Intermediate Language compiler<br />
License: BSD 2-clause<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/secilc-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/secilc-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 96 KiB<br />

## Configuration

This package does not have a traditional configuration script.

## Compilation and Installation

To compile SECilC, run the following command:

```bash
export XML_CATALOG_FILES="/etc/xml/catalog"
make
```

Next, run the test suite:

```bash
make test && echo $?
```

If `0` is output at the end, the test ran successfully.

Finally, to install SECilC into the build tree, run the following command:

```bash
make install
unset XML_CATALOG_FILES
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | secil2conf, secil2tree, and secilc |

| Navigation |||
| --- | --- | ---: |
| [<<](./XmlTo.md) XmlTo | [HOME](../README.md) | LibEConf 64-bit [>>](./LibEConf_64bit.md) |

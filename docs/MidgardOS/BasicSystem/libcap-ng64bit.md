# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libxcrypt32bit.md) libxcrypt 32-bit | [HOME](../README.md) | libcap-ng 32-bit [>>](./libcap-ng32bit.md) |

## libcap-ng 64-bit

Name: libcap-ng 64-bit<br />
Summary: An alternate POSIX capabilities library<br />
License: GPL v2/LGPL v2.1<br />
Version: 0.8.5<br />
URL: [https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.tar.gz](https://people.redhat.com/sgrubb/libcap-ng/libcap-ng-0.8.5.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 384 KiB<br />

## Configuration

To configure libxcrypt 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static
```

## Compilation and Installation

To compile libxcrypt 64-bit, run the following commands:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install libxcrypt 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libcap-ng.la
rm -fv /usr/lib64/libdrop_ambient.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | captest, filecap, netcap, and pscap |
| Installed Libraries | libcap-ng.so and libdrop_ambient.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./libxcrypt32bit.md) libxcrypt 32-bit | [HOME](../README.md) | libcap-ng 32-bit [>>](./libcap-ng32bit.md) |

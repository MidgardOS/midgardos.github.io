# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Audit_32bit.md) Linux Audit Framework 32-bit | [HOME](../README.md) | Shadow Utils 64-bit - pass 1: Without PAM [>>](./Shadow_64bit-pass1.md) |

## RPCSvc-Proto

Name: RPCSvc-Proto<br />
Summary: The RPCSVC headers and generation tool<br />
License: GPL v2.0 or later<br />
Version: 1.4.4<br />
URL: [https://github.com/thkukuk/rpcsvc-proto/releases/download/v1.4.4/rpcsvc-proto-1.4.4.tar.xz](https://github.com/thkukuk/rpcsvc-proto/releases/download/v1.4.4/rpcsvc-proto-1.4.4.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 520 KiB <br />

## Configuration

To configure RPCSvc-Proto for install into the build root, run the following command:

```bash
./configure --prefix=/usr
```

## Compilation and Installation

To compile RPCSvc-Proto, run the following command:

```bash
make
```

Now, run the test suite:

```bash
make check
```

Finally, to install RPCSvc-Proto into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | rpcgen |
| Installed Files | /usr/include/rpcsvc |

| Navigation |||
| --- | --- | ---: |
| [<<](./Audit_32bit.md) Linux Audit Framework 32-bit | [HOME](../README.md) | Shadow Utils 64-bit - pass 1: Without PAM [>>](./Shadow_64bit-pass1.md) |

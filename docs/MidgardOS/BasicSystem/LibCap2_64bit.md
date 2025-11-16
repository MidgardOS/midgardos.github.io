# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ACL_32bit.md) ACL 32-bit | [HOME](../README.md) | LibCap2 32-bit [>>](./LibCap2_32bit.md) |

## LibCap2 64-bit

Name: LibCap2 64-bit<br />
Summary: Userspace tools for POSIX 1003.1e privilege capabilities<br />
License: BSD-3-Clause or GPL v2.0<br />
Version: 2.76<br />
URL: [https://www.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.76.tar.xz](https://www.kernel.org/pub/linux/libs/security/linux-privs/libcap2/libcap-2.76.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 3.1 MiB<br />

## Configuration

The package does not have a traditional configuration script.

## Compilation and Installation

To compile LibCap2 64-bit, run the following commands:

```bash
sed -i '/install -m.*STA/d' libcap/Makefile
make prefix=/usr lib=lib64
```

Next, run the test suite:

```bash
make test
```

Finally, to install LibCap2 64-bit into the build tree, run the following command:

```bash
make prefix=/usr lib=lib64 install
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | capsh, getcap, getpcaps, and setcap |
| Installed Libraries | libcap.so and libpsx.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./ACL_32bit.md) ACL 32-bit | [HOME](../README.md) | LibCap2 32-bit [>>](./LibCap2_32bit.md) |

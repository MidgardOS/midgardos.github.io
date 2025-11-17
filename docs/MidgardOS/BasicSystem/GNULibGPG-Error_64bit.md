# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./CheckPolicy.md) CheckPolicy | [HOME](../README.md) | GNU LibGPG-Error 32-bit [>>](./GNULibGPG-Error_32bit.md) |

## GNU LibGPG-Error 64-bit

Name: GNU LibGPG-Error 64-bit<br />
Summary: A library that defines the various error states for GnuPG<br />
License: GPL v2.0 or later/LGPL v2.1 or later<br />
Version: 1.56<br />
URL: [https://gnupg.org/ftp/gcrypt/libgpg-error/libgpg-error-1.56.tar.bz2](https://gnupg.org/ftp/gcrypt/libgpg-error/libgpg-error-1.56.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 1.9 MiB<br />

## Configuration

To configure GNU LibGPG-Error 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
            --docdir=/usr/share/doc/libgpg-error-1.56   \
            --enable-install-gpg-error-config           \
            --disable-static
```

## Compilation and Installation

To compile GNU LibGPG-Error 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU LibGPG-Error 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libgpg-error.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | gpg-error, gpg-error-config, gpgrt-config, and yat2m |
| Installed Libraries | libgpg-error.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./CheckPolicy.md) CheckPolicy | [HOME](../README.md) | GNU LibGPG-Error 32-bit [>>](./GNULibGPG-Error_32bit.md) |

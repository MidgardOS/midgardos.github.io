# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTexinfo.md) GNU Texinfo | [HOME](../README.md) | Util-Linux 32-bit [>>](./UtilLinux32bit.md) |

## Util-Linux 64-bit

Name: Util-Linux 64-bit<br />
Summary: Miscellaneous command-line utilities - 64-bit build<br />
License: GPL v2 or later<br />
Version: 2.41.1<br />
URL: [https://www.kernel.org/pub/linux/utils/util-linux/v2.41/util-linux-2.41.1.tar.xz](https://www.kernel.org/pub/linux/utils/util-linux/v2.41/util-linux-2.41.1.tar.xz)<br />

## Configuration

To configure Util Linux for install into the build root, run the following command:

```bash
mkdir -pv /var/lib/hwclock
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --runstatedir=/run      \
            --disable-chfn-chsh     \
            --disable-login         \
            --disable-nologin       \
            --disable-su            \
            --disable-setpriv       \
            --disable-runuser       \
            --disable-pylibmount    \
            --disable-static        \
            --disable-liblastlog2   \
            --without-python        \
            ADJTIME_PATH=/var/lib/hwclock/adjtime \
            --docdir=/usr/share/doc/util-linux-2.41.1
```

## Compilation and Installation

To compile Util-Linux 64-bit, run the following command:

```bash
make
```

Finally, to install Util-Linux 64-bit into the build tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTexinfo.md) GNU Texinfo | [HOME](../README.md) | Util-Linux 32-bit [>>](./UtilLinux32bit.md) |

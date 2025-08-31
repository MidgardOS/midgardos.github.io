# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./UtilLinux64bit.md) Util-Linux 64-bit | [HOME](../README.md) | Cleaning up and Archiving Temporary System [>>](./CleanupAndArchival.md) |

## Util-Linux 32-bit

Name: Util-Linux 32-bit<br />
Summary: Miscellaneous command-line utilities - 32-bit build<br />
License: GPL v2 or later<br />
Version: 2.41.1<br />
URL: [https://www.kernel.org/pub/linux/utils/util-linux/v2.41/util-linux-2.41.1.tar.xz](https://www.kernel.org/pub/linux/utils/util-linux/v2.41/util-linux-2.41.1.tar.xz)<br />

## Configuration

To configure Util Linux 32-bir for install into the build root, run the following command:

```bash
CC="gcc -m32" \
./configure --host=i686-midgardos-linux-gnu           \
            --libdir=/usr/lib                         \
            --libexecdir=/usr/lib                     \
            --runstatedir=/run                        \
            --docdir=/usr/share/doc/util-linux-2.41.1 \
            --disable-chfn-chsh                       \
            --disable-login                           \
            --disable-nologin                         \
            --disable-su                              \
            --disable-setpriv                         \
            --disable-runuser                         \
            --disable-pylibmount                      \
            --disable-static                          \
            --disable-liblastlog2                     \
            --without-python                          \
            ADJTIME_PATH=/var/lib/hwclock/adjtime
```

## Compilation and Installation

To compile Util-Linux 32-bit, run the following command:

```bash
make
```

Finally, to install Util-Linux 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$(pwd)/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./UtilLinux64bit.md) Util-Linux 64-bit | [HOME](../README.md) | Cleaning up and Archiving Temporary System [>>](./CleanupAndArchival.md) |

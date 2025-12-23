# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Kmod_64bit.md) Kmod 64-bit | [HOME](../README.md) | GNU Coreutils [>>](./GNUCoreutils.md) |

## Kmod 32-bit

Name: Kmod 32-bit<br />
Summary: Tools and libraries for loading Kernel modules<br />
License: LGPL v2.1 or later<br />
Version: 34.2<br />
URL: [https://www.kernel.org/pub/linux/utils/kernel/kmod/kmod-34.2.tar.xz](https://www.kernel.org/pub/linux/utils/kernel/kmod/kmod-34.2.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 508 KiB<br />

## Configuration

To configure Kmod 32-bit for install into the build root, run the following command:

```bash
rm -rf build
mkdir -p build
cd       build

PKG_CONFIG_PATH="/usr/lib/pkgconfig"                    \
CC="gcc -m32 -march=i686"                               \
CXX="g++ -m32 -march=i686"                              \
meson setup --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
            --buildtype=release                         \
            -Dbashcompletiondir=/etc/bash_completion.d  \
            -Dfishcompletiondir=no                      \
            -Dzshcompletiondir=/etc/zsh_completion.d    \
            -Dmoduledir=/usr/lib/modules                \
            -Ddlopen=all                                \
            -Ddocs=false                                \
            ..
```

## Compilation and Installation

To compile Kmod 32-bit, run the following command:

```bash
ninja
```

Finally, to install Kmod 32-bit into the build tree, run the following command:

```bash
DESTDIR=$PWD/DESTDIR ninja install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of Kmod 32-bit for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./Kmod_64bit.md) Kmod 64-bit | [HOME](../README.md) | GNU Coreutils [>>](./GNUCoreutils.md) |

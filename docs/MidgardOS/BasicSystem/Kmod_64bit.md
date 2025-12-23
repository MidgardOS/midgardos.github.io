# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Linux-PAM_32bit.md) Linux PAM 32-bit | [HOME](../README.md) | Kmod 32-bit [>>](./Kmod_32bit.md) |

## Kmod 64-bit

Name: Kmod 64-bit<br />
Summary: Tools and libraries for loading Kernel modules<br />
License: LGPL v2.1 or later<br />
Version: 34.2<br />
URL: [https://www.kernel.org/pub/linux/utils/kernel/kmod/kmod-34.2.tar.xz](https://www.kernel.org/pub/linux/utils/kernel/kmod/kmod-34.2.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 508 KiB<br />

## Configuration

To configure Kmod 64-bit for install into the build root, run the following command:

```bash
mkdir -p build
cd       build

meson setup --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
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

To compile Kmod 64-bit, run the following command:

```bash
ninja
```

There are no tests for this package.

Finally, to install Kmod 64-bit into the build tree, run the following command:

```bash
ninja install
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | kmod, depmod, insmod, lsmod, modinfo, modprobe, and rmmod |
| Installed Libraries | libkmod.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./Linux-PAM_32bit.md) Linux PAM 32-bit | [HOME](../README.md) | Kmod 32-bit [>>](./Kmod_32bit.md) |

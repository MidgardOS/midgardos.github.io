# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Vim.md) Vi Improved | [HOME](../README.md) | Python module MarkupSafe [>>](./Python313-MarkupSafe.md) |

## GNU Nano

Name: GNU Nano<br />
Summary: A simple terminal text editor<br />
License: GPL v3.0 or later<br />
Version: 8.7<br />
URL: [https://ftp.gnu.org/gnu/nano/nano-8.7.tar.xz](https://ftp.gnu.org/gnu/nano/nano-8.7.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 3.7 MiB<br />

## Configuration

To configure GNU Nano for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/nano-8.7    \
            --sysconfdir=/etc                   \
            --disable-rpath                     \
            --enable-threads=posix              \
            --enable-utf8                       \
            --enable-year2038
```

## Compilation and Installation

To compile GNU Nano, run the following command:

```bash
make
```

This package does not come with a test suite.

Finally, to install GNU Nano into the build tree, run the following command:

```bash
make install
install -v -m 644 -o root -g root ../system_files/etc/nanorc /etc/
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | nano and rnano |
| Installed Files | /usr/share/nano |

| Navigation |||
| --- | --- | ---: |
| [<<](./Vim.md) Vi Improved | [HOME](../README.md) | Python module MarkupSafe [>>](./Python313-MarkupSafe.md) |

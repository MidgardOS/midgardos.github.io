# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses32bit.md) NCurses 32-bit | [HOME](../README.md) | GNU Coreutils [>>](./GNUCoreutils.md) |

## GNU Bash

Name: GNU Bash<br />
Summary: The GNU Bourne-Again Shell<br />
License: GPL v3 or later<br />
Version: 5.3<br />
URL: [https://ftp.gnu.org/gnu/bash/bash-5.3.tar.gz](https://ftp.gnu.org/gnu/bash/bash-5.3.tar.gz)<br />

## Configuration

To configure GNU Bash for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --build=${BRFS_HOST}    \
            --host=$BRFS_TARGET     \
            --without-bash-malloc
```

## Compilation and Installation

To compile GNU Bash, run the following command:

```bash
make
```

Finally, to install GNU Bash into the cross-tools tree, run the following command:

```bash
make DESTDIR=${BRFS} install
```

Temporarily install a compatibility symbolic link for `/bin/sh` to point to `/bin/bash`. This will be replaced later when `zsh` is installed.

```bash
ln -sv bash $BRFS/bin/sh
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses32bit.md) NCurses 32-bit | [HOME](../README.md) | GNU Coreutils [>>](./GNUCoreutils.md) |

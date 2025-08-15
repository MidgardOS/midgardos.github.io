| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBash.md) | [HOME](../README.md) | [>>](./GNUDiffutils.md) |

# GNU Coreutils

Name: GNU Coreutils<br />
Summary: Basic file, shell, and text utilities needed for working with any Linux system<br />
License: GPL v3 or later<br />
Version: 9.7<br />
URL: [https://ftp.gnu.org/gnu/coreutils/coreutils-9.7.tar.xz](https://ftp.gnu.org/gnu/coreutils/coreutils-9.7.tar.xz)<br />

## Configuration

To configure GNU Coreutils for install into the build root, run the following command:

```bash
./configure --prefix=/usr                     \
            --host=$BRFS_TARGET               \
            --build=$BRFS_HOST                \
            --enable-no-install-program=kill,uptime,hostname
```

The `--enable-no-install-program` is to allow them to come from other packages:

- `hostname` will be provided by the `hostname` package
- `kill` will be provided by the `utils-linux` package
- `uptime` will be provided by the `procps-ng`

## Compilation and Installation

To compile GNU Coreutils, run the following command:

```bash
make
```

Finally, to install GNU Coreutils into the build tree, run the following command:

```bash
make DESTDIR=${BRFS} install
mkdir -pv $BRFS/usr/share/man/man8
mv -v $BRFS/usr/share/man/man1/chroot.1 $BRFS/usr/share/man/man8/chroot.8
sed -i 's/"1"/"8"/'                     $BRFS/usr/share/man/man8/chroot.8
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBash.md) | [HOME](../README.md) | [>>](./GNUDiffutils.md) |

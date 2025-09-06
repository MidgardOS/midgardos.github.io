# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./DejaGNU.md) DejaGNU | [HOME](../README.md) | GNU Binutils [>>](./GNUBinutils.md) |

# PkgConf

Name: PkgConf<br />
Summary: A software compilation and linker metadata tool and library<br />
License: ISC<br />
Version: 2.5.1<br />
URL: [https://distfiles.ariadne.space/pkgconf/pkgconf-2.5.1.tar.xz](https://distfiles.ariadne.space/pkgconf/pkgconf-2.5.1.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 5.0 MiB<br />

**NOTE: At this time, we don't use the download from GitHub for this package as it requires AutoConf and AutoMake, which have not been installed yet**

## Configuration

To configure PkgConf for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/usr                         \
            --libdir=/usr/lib64                   \
            --libexecdir=/usr/lib64               \
            --docdir=/usr/share/doc/pkgconf-2.5.1 \
            --disable-static                      \
            --with-pkg-config-dir=/usr/lib64/pkgconfig:/usr/lib/pkgconfig:/usr/share/pkgconfig
```

## Compilation and Installation

To compile PkgConf, run the following command:

```bash
make
```

Finally, to install PkgConf into the cross-tools tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libpkgconf.la
ln -sv pkgconf   /usr/bin/pkg-config
ln -sv pkgconf.1 /usr/share/man/man1/pkg-config.1
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | pkgconf, pkg-config (link to pkgconf), and bomtool |
| Installed Libraries | libpkgconf.so |
| Installed Directories | /usr/share/doc/pkgconf-2.5.1 |

| Navigation |||
| --- | --- | ---: |
| [<<](./DejaGNU.md) DejaGNU | [HOME](../README.md) | GNU Binutils [>>](./GNUBinutils.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibEConf_64bit.md) LibEConf 64-bit | [HOME](../README.md) | Linux PAM 64-bit [>>](./Linux-PAM_64bit.md) |

## LibEConf 32-bit

Name: LibEConf 32-bit<br />
Summary: Enhanced configuration file parser library<br />
License: MIT<br />
Version: 0.8.3<br />
URL: [https://github.com/openSUSE/libeconf/archive/refs/tags/v0.8.3.tar.gz](https://github.com/openSUSE/libeconf/archive/refs/tags/v0.8.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.2 MiB<br />

## Configuration

To configure LibEConf 32-bit for install into the build root, run the following commands:

```bash
rm -rf x86_64-pc-linux-gnu
CONFIG_SHELL="${CONFIG_SHELL:-/usr/bin/bash}"; export CONFIG_SHELL
CFLAGS="${CFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CFLAGS
CXXFLAGS="${CXXFLAGS:--O2 -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CXXFLAGS
FFLAGS="${FFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FFLAGS
FCFLAGS="${FCFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FCFLAGS
LDFLAGS="${LDFLAGS:-}"; export LDFLAGS

/usr/bin/meson setup --buildtype=plain --prefix=/usr --libdir=/usr/lib --libexecdir=/usr/lib --bindir=/usr/bin --sbindir=/usr/sbin --includedir=/usr/include --datadir=/usr/share --mandir=/usr/share/man --infodir=/usr/share/info --localedir=/usr/share/locale --sysconfdir=/etc --localstatedir=/var --sharedstatedir=/var/lib --wrap-mode=nodownload --auto-features=enabled -Ddefault_library=both . i686-pc-linux-gnu
```

## Compilation and Installation

To compile LibEConf 32-bit, run the following command:

```bash
/usr/bin/meson compile -C i686-pc-linux-gnu -j4 --verbose
```

Finally, to install LibEConf 32-bit into the build tree, run the following commands:

```bash
DESTDIR=$PWD/DESTDIR \
/usr/bin/meson install -C i686-pc-linux-gnu --no-rebuild
rm -fv DESTDIR/usr/lib/libeconf.a
cp -Rv DESTDIR/usr/lib/* /usr/lib/

unset CONFIG_SHELL
unset CFLAGS
unset CXXFLAGS
unset FFLAGS
unset FCFLAGS
unset LDFLAGS
```

## Contents

See the contents section of the 64-bit build of LibEConf for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibEConf_64bit.md) LibEConf 64-bit | [HOME](../README.md) | Linux PAM 64-bit [>>](./Linux-PAM_64bit.md) |

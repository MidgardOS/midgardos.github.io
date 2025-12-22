# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./SECilC.md) SELinux Cil Compiler | [HOME](../README.md) | LibEConf 32-bit [>>](./LibEConf_32bit.md) |

## LibEConf 64-bit

Name: LibEConf 64-bit<br />
Summary: Enhanced configuration file parser library<br />
License: MIT<br />
Version: 0.8.3<br />
URL: [https://github.com/openSUSE/libeconf/archive/refs/tags/v0.8.3.tar.gz](https://github.com/openSUSE/libeconf/archive/refs/tags/v0.8.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.2 MiB<br />

## Configuration

To configure LibEConf 64-bit for install into the build root, run the following commands:

```bash
CONFIG_SHELL="${CONFIG_SHELL:-/usr/bin/bash}"; export CONFIG_SHELL
CFLAGS="${CFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CFLAGS
CXXFLAGS="${CXXFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CXXFLAGS
FFLAGS="${FFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FFLAGS
FCFLAGS="${FCFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FCFLAGS
LDFLAGS="${LDFLAGS:-}"; export LDFLAGS

/usr/bin/meson setup --buildtype=plain --prefix=/usr --libdir=/usr/lib64 --libexecdir=/usr/lib64 --bindir=/usr/bin --sbindir=/usr/sbin --includedir=/usr/include --datadir=/usr/share --mandir=/usr/share/man --infodir=/usr/share/info --localedir=/usr/share/locale --sysconfdir=/etc --localstatedir=/var --sharedstatedir=/var/lib --wrap-mode=nodownload --auto-features=enabled -Ddefault_library=both . x86_64-pc-linux-gnu
```

## Compilation and Installation

To compile LibEConf 64-bit, run the following command:

```bash
/usr/bin/meson compile -C x86_64-pc-linux-gnu -j4 --verbose
```

Next, run the test suite:

```bash
/usr/bin/meson test -C x86_64-pc-linux-gnu --verbose
```

Finally, to install LibEConf 64-bit into the build tree, run the following commands:

```bash
/usr/bin/meson install -C x86_64-pc-linux-gnu --no-rebuild
rm -fv /usr/lib64/libeconf.a

unset CONFIG_SHELL
unset CFLAGS
unset CXXFLAGS
unset FFLAGS
unset FCFLAGS
unset LDFLAGS
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | econftool |
| Installed Libraries | libeconf.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./SECilC.md) SELinux Cil Compiler | [HOME](../README.md) | LibEConf 32-bit [>>](./LibEConf_32bit.md) |

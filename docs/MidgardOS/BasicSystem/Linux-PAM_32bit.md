# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Linux-PAM_64bit.md) Linux PAM 64-bit | [HOME](../README.md) | SCDoc [>>](./SCDoc.md) |

## Linux PAM 32-bit

Name: Linux PAM 32-bit<br />
Summary: The Linux implementation of Pluggable Authentication Modules<br />
License: GPL v2.0 or later or BSD 3-Clause<br />
Version: 1.7.1<br />
URL: [https://github.com/linux-pam/linux-pam/releases/download/v1.7.1/Linux-PAM-1.7.1.tar.xz](https://github.com/linux-pam/linux-pam/releases/download/v1.7.1/Linux-PAM-1.7.1.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 5.9 MiB<br />

## Preparation

The patches required for Linux PAM are already applied if the unpacked sources were not removed in the last step.

## Configuration

**NOTE: This package will need rebuilt after SystemD is installed.**

To configure Linux PAM 32-bit for install into the build root, run the following commands:

```bash
unset CONFIG_SHELL
unset CFLAGS
unset CXXFLAGS
unset FFLAGS
unset FCFLAGS
unset LDFLAGS
unset PKG_CONFIG_LIBDIR
rm -rf x86_64-pc-linux-gnu && rm -rf i686-pc-linux-gnu
CONFIG_SHELL="${CONFIG_SHELL:-/usr/bin/bash}"; export CONFIG_SHELL
CFLAGS="${CFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CFLAGS
CXXFLAGS="${CXXFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CXXFLAGS
FFLAGS="${FFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FFLAGS
FCFLAGS="${FCFLAGS:--O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FCFLAGS
LDFLAGS="${LDFLAGS:-}"; export LDFLAGS
PKG_CONFIG_LIBDIR=/usr/lib/pkgconfig; export PKG_CONFIG_LIBDIR

/usr/bin/meson setup --buildtype=plain --prefix=/usr --libdir=/usr/lib --libexecdir=/usr/lib --bindir=/usr/bin --sbindir=/usr/sbin --includedir=/usr/include --datadir=/usr/share --mandir=/usr/share/man --infodir=/usr/share/info --localedir=/usr/share/locale --sysconfdir=/etc --localstatedir=/var --sharedstatedir=/var/lib --wrap-mode=nodownload --auto-features=enabled \
    -Dvendordir=/etc \
    -Ddocdir=/usr/share/doc/pam-1.7.1 \
    -Dhtmldir=/usr/share/doc/pam-1.7.1/html \
    -Dpdfdir=/usr/share/doc/pam-1.7.1/pdf \
    -Dsecuredir=/usr/lib/security \
    -Dpam_unix-try-getspnam=true \
    -Dlogind=disabled \
    -Delogind=disabled \
    -Dnis=disabled \
    -Dpwaccess=disabled . i686-pc-linux-gnu
```

## Compilation and Installation

To compile Linux PAM 32-bit, run the following command:

```bash
/usr/bin/meson compile -C i686-pc-linux-gnu -j4 --verbose
```

Finally, to install Linux PAM 32-bit into the build tree, run the following commands:

```bash
DESTDIR=$PWD/DESTDIR /usr/bin/meson install -C i686-pc-linux-gnu --no-rebuild
rm -rfv DESTDIR/usr/lib/systemd
cp -Rv DESTDIR/usr/lib/* /usr/lib/

unset CONFIG_SHELL
unset CFLAGS
unset CXXFLAGS
unset FFLAGS
unset FCFLAGS
unset LDFLAGS
unset PKG_CONFIG_LIBDIR
```

## Contents

See the contents section of the 64-bit build of Linux PAM for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./Linux-PAM_64bit.md) Linux PAM 64-bit | [HOME](../README.md) | SCDoc [>>](./SCDoc.md) |

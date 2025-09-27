# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Automake.md) GNU Automake | [HOME](../README.md) | Linux Audit Framework 32-bit [>>](./audit32bit.md) |

## Linux Audit Framework 64-bit

Name: Linux Audit Framework 64-bit<br />
Summary: The Linux auditing framework<br />
License: GPL v2.0 and LGPL v2.1<br />
Version: 4.1.2<br />
URL: [https://github.com/linux-audit/audit-userspace/archive/refs/tags/v4.1.2.tar.gz](https://github.com/linux-audit/audit-userspace/archive/refs/tags/v4.1.2.tar.gz)<br />

Average Build Time: approximately 1 SBU<br />
Used Install Space: 5.3 MiB <br />

## Configuration

To configure the Linux Audit Framework 64-bit package for compilation, run the following commands:

```bash
autoreconf -fv --install
sed -i 's/ ids / /' audisp/plugins/Makefile.am
sed -i 's/ ids / /' audisp/plugins/Makefile.in
export CFLAGS="-O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables -fno-strict-aliasing"
export CXXFLAGS="$CFLAGS"
export LDFLAGS="-Wl,-z,relro,-z,now"
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --sysconfdir=/etc       \
            --enable-gssapi-krb5=no \
            --with-libwrap          \
            --with-libcap-ng=yes    \
            --with-golang=no        \
            --with-python=no        \
            --with-python3=no       \
            --disable-zos-remote
```

Note that this does not disable the static libraries. This is required during build due to how the Makefile for this package is written.

## Compilation and Installation

To compile the Linux Audit Framework 64-bit, run the following command:

```bash
CFLAGS=$CFLAGS make -O -j1 V=1 VERBOSE=1 
```

Next, run the test suite:

```bash
make check
```

Finally, to install the Linux Audit Framework 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libaudit.la
rm -fv /usr/lib64/libauparse.la
rm -fv /usr/lib64/libauplugin.la
unset CFLAGS
unset CXXFLAGS
unset LDFLAGS
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | audisp-af_unix, audisp-filter, audisp-remote, audisp-syslog, auditctl, augenrules, aulast, aulastlog, aureport, ausearch, and ausyscall |
| Installed Libraries | libaudit.so, libauparse.so, and libauplugin.so |
| Installed Services | auditd |
| Installed Files | /etc/libaudit.conf, /etc/audit/audisp-filter.conf, /etc/audit/audisp-remote.conf, audit-stop.rules, and auditd.conf |

| Navigation |||
| --- | --- | ---: |
| [<<](./Automake.md) GNU Automake | [HOME](../README.md) | Linux Audit Framework 32-bit [>>](./audit32bit.md) |

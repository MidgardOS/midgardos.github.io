# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Audit_64bit.md) Linux Audit Framework 64-bit | [HOME](../README.md) | RPCSvc-Proto [>>](./RPCSvc-Proto.md) |

## Linux Audit Framework 32-bit

Name: Linux Audit Framework 32-bit<br />
Summary: The Linux auditing framework<br />
License: GPL v2.0 and LGPL v2.1<br />
Version: 4.1.2<br />
URL: [https://github.com/linux-audit/audit-userspace/archive/refs/tags/v4.1.2.tar.gz](https://github.com/linux-audit/audit-userspace/archive/refs/tags/v4.1.2.tar.gz)<br />

Average Build Time: approximately 1 SBU<br />
Used Install Space: 5.3 MiB <br />

## Configuration

To configure the Linux Audit Framework 32-bit package for compilation, run the following commands:

```bash
make distclean
export CFLAGS="-O2 -g -m32 -march=i686 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables -fno-strict-aliasing"
export CXXFLAGS="$CFLAGS"
export LDFLAGS="-Wl,-z,relro,-z,now"
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --sysconfdir=/etc        \
            --enable-gssapi-krb5=no  \
            --with-libwrap           \
            --with-libcap-ng=yes     \
            --with-golang=no         \
            --with-python=no         \
            --with-python3=no        \
            --disable-zos-remote
```

## Compilation and Installation

To compile the Linux Audit Framework 32-bit, run the following command:

```bash
CFLAGS=$CFLAGS make -O -j1 V=1 VERBOSE=1
```

Finally, to install the Linux Audit Framework 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/*.la
rm -rfv DESTDIR/usr/lib/initscripts
rm -rfv DESTDIR/usr/lib/systemd
rm -rfv DESTDIR/usr/lib/tmpfiles.d
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of the Linux Audit Framework for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./Audit_64bit.md) Linux Audit Framework 64-bit | [HOME](../README.md) | RPCSvc-Proto [>>](./RPCSvc-Proto.md) |

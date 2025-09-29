# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./shadow64bit-pass1.md) Shadow Utils 64-bit | [HOME](../README.md) | GNU Compiler Collection [>>](./GNUGCC.md) |

## Shadow Utils 32-bit

Name: Shadow Utils 32-bit<br />
Summary: Managing user and group passwords securely<br />
License: BSD 3-Clause<br />
Version: 4.18.0<br />
URL: [https://github.com/shadow-maint/shadow/releases/download/4.18.0/shadow-4.18.0.tar.xz](https://github.com/shadow-maint/shadow/releases/download/4.18.0/shadow-4.18.0.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 12 MiB<br />

## Configuration

To configure Shadow Utils 32-bit for install into the build root, run the following commands:

```bash
make distclean
export CFLAGS="-O2 -g -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables -fpie"
export LDFLAGS="-pie"
./configure --host=i686-midgardos-linux-gnu \
            --sysconfdir=/etc               \
            --libdir=/usr/lib               \
            --with-{b,yes}crypt             \
            --without-libbsd                \
            --with-acl                      \
            --with-attr                     \
            --without-btrfs                 \
            --without-nscd                  \
            --with-audit                    \
            --without-sssd                  \
            --with-group-name-max-length=32
```

## Compilation and Installation

To compile Shadow Utils 32-bit, run the following command:

```bash
make
```

Finally, to install Shadow Utils 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR exec_prefix=/usr install
rm -fv DESTDIR/usr/lib/libsubid.a
rm -fv DESTDIR/usr/lib/libsubid.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
unset CFLAGS
unset LDFLAGS
```

## Contents

See the contents section of the 64-bit build of Shadow Utils for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./shadow64bit-pass1.md) Shadow Utils 64-bit | [HOME](../README.md) | GNU Compiler Collection [>>](./GNUGCC.md) |

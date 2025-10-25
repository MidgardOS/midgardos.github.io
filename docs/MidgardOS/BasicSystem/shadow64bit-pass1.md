# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./rpcsvc-proto.md) RPCSVC Proto | [HOME](../README.md) | Shadow Utils 32-bit [>>](./shadow32bit.md) |

## Shadow Utils 64-bit

Name: Shadow Utils 64-bit<br />
Summary: Managing user and group passwords securely<br />
License: BSD 3-Clause<br />
Version: 4.18.0<br />
URL: [https://github.com/shadow-maint/shadow/releases/download/4.18.0/shadow-4.18.0.tar.xz](https://github.com/shadow-maint/shadow/releases/download/4.18.0/shadow-4.18.0.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 12 MiB<br />

## Configuration

First, disable the `groups` and `login` programs and the man pages that are already from Coreutils, Util-Linux, or from the Man-Pages package.

```bash
sed -i 's/groups$(EXEEXT) //' src/Makefile.in
find man -name Makefile.in -exec sed -i 's/groups\.1 / /'   {} \;
find man -name Makefile.in -exec sed -i 's/getspnam\.3 / /' {} \;
find man -name Makefile.in -exec sed -i 's/passwd\.5 / /'   {} \;
```

To use a more secure hashing format than the old crypt algorithm that limits passwords to eight characters along with cleaning up the PATH variable since `/bin` and `/sbin` are symlinks into `/usr`, run the following `sed` command:

```bash
sed -e 's:#ENCRYPT_METHOD DES:ENCRYPT_METHOD YESCRYPT:' \
    -e 's:/var/mail:/var/spool/mail:'                   \
    -e '/PATH=/{s@/sbin:@@;s@/bin:@@}'                  \
    -i etc/login.defs
```

To configure Shadow Utils 64-bit for install into the build root, run the following commands:

```bash
touch /usr/bin/passwd
export CFLAGS="-O2 -g -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables -fpie"
export LDFLAGS="-pie"
./configure --sysconfdir=/etc               \
            --libdir=/usr/lib64             \
            --with-{b,yes}crypt             \
            --without-libbsd                \
            --with-acl                      \
            --with-attr                     \
            --without-btrfs                 \
            --with-audit                    \
            --without-sssd                  \
            --with-group-name-max-length=32
```

Note, we're disabling System Security Services daemon support for now. This will be added back in a later build.

## Compilation and Installation

To compile Shadow Utils 64-bit, run the following command:

```bash
make
```

Unfortunately, the test suite must not be run under the root account, and since `su` and `login` are not installed, it is not possible to run the tests until Util-Linux and PAM are installed.

Finally, to install Shadow Utils 64-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR exec_prefix=/usr install
make DESTDIR=$PWD/DESTDIR -C man install-man
rm -fv DESTDIR/usr/bin/login
rm -fv DESTDIR/usr/bin/su
rm -fv DESTDIR/usr/sbin/nologin
rm -fv DESTDIR/usr/share/man/man1/login.1
rm -fv DESTDIR/usr/share/man/man1/su.1
rm -fv DESTDIR/usr/share/man/man5/suauth.5
rm -fv DESTDIR/usr/share/man/man8/nologin.8
rm -fv DESTDIR/usr/lib64/libsubid.a
rm -fv DESTDIR/usr/lib64/libsubid.la
cp -Rv DESTDIR/usr/* /usr/
cp -Rv DESTDIR/etc/* /etc/
unset CFLAGS
unset LDFLAGS
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | chage, chfn, chgpasswd, chpasswd, chsh, expiry, faillog, getsubids, gpasswd, groupadd, groupdel, groupmems, groupmod, grpck, grpconv, logoutd, newgidmap, newgrp, newuidmap, newusers, passwd, pwck, pwconv, pwunconv, sg, useradd, userdel, usermod, vigr, vipw |
| Installed Libraries | libsubid.so |
| Installed Files | /etc/limits, /etc/login.access, /etc/login.defs |

| Navigation |||
| --- | --- | ---: |
| [<<](./rpcsvc-proto.md) RPCSVC Proto | [HOME](../README.md) | Shadow Utils 32-bit [>>](./shadow32bit.md) |

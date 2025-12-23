# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Kmod_32bit.md) Kmod 32-bit | [HOME](../README.md) | Shadow Utils 32-bit [>>](./Shadow_32bit-pass1.md) |

## Shadow Utils 64-bit - With Linux PAM Support

Name: Shadow Utils 64-bit<br />
Summary: Managing user and group passwords securely<br />
License: BSD 3-Clause<br />
Version: 4.18.0<br />
URL: [https://github.com/shadow-maint/shadow/releases/download/4.18.0/shadow-4.18.0.tar.xz](https://github.com/shadow-maint/shadow/releases/download/4.18.0/shadow-4.18.0.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 12 MiB<br />

## Notes

The Shadow package hooks into a number of login-related components. This means that there will be additional rebuilds required in the future. Namely:

- After install of BtrFS tools and libraries
- After install of SystemD and its libraries
- After System Security Services Daemon and its libraries is installed
- After LibBSD is installed

The above packages will be installed later, and will require rebuild of this package at that time.

## Configuration

First, disable the `groups` program and the man pages that are already from Coreutils, Util-Linux, or from the Man-Pages package.

```bash
sed -i 's/groups$(EXEEXT) //' src/Makefile.in
find man -name Makefile.in -exec sed -i 's/groups\.1 / /'   {} \;
find man -name Makefile.in -exec sed -i 's/getspnam\.3 / /' {} \;
find man -name Makefile.in -exec sed -i 's/passwd\.5 / /'   {} \;
```

Next, set sane defaults and use a more secure hashing format than the old crypt algorithm that limits passwords to eight characters along with cleaning up the PATH variable since `/bin` and `/sbin` are symlinks into `/usr`, run the following `sed` command:

```bash
sed -e 's:#ENCRYPT_METHOD DES:ENCRYPT_METHOD YESCRYPT:'			\
    -e 's:/var/mail:/var/spool/mail:'					\
    -e '/PATH=/{s@/sbin:@@;s@/bin:@@}'					\
    -e 's:MOTD_FILE	/etc/motd:#MOTD_FILE	/etc/motd:'		\
    -e 's:PASS_MIN_LEN	5:PASS_MIN_LEN	8:'				\
    -e 's:SYS_UID_MIN		  101:SYS_UID_MIN		  100:'	\
    -e 's:SYS_GID_MIN		  101:SYS_GID_MIN		  100:'	\
    -e 's:LOGIN_RETRIES		5:LOGIN_RETRIES		3:'		\
    -e 's:PASS_CHANGE_TRIES	5:PASS_CHANGE_TRIES	3:'		\
    -i etc/login.defs
```

To configure Shadow Utils 64-bit for install into the build root, run the following commands:

```bash
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

Note, we're disabling System Security Services daemon, BtrFS and libbsd support for now. This will be added back in a later build.

## Compilation and Installation

To compile Shadow Utils 64-bit, run the following command:

```bash
make
```

Unfortunately, the test suite must **not** be run under the root account, and since `su` and `login` are not installed, it is not possible to run the tests until Util-Linux and PAM are installed.

Finally, to install Shadow Utils 64-bit into the build tree, run the following commands:

**NOTE: Temporarily install `su`, it will be replaced with the versions from `util-linux`**

```bash
unalias cp
make DESTDIR=$PWD/DESTDIR exec_prefix=/usr install
make DESTDIR=$PWD/DESTDIR -C man install-man
rm -fv DESTDIR/usr/bin/login
rm -fv DESTDIR/usr/sbin/nologin
rm -fv DESTDIR/usr/share/man/man1/login.1
rm -fv DESTDIR/usr/share/man/man5/suauth.5
rm -fv DESTDIR/usr/share/man/man8/nologin.8
rm -fv DESTDIR/usr/lib64/libsubid.a
rm -fv DESTDIR/usr/lib64/libsubid.la
cp -Rfv DESTDIR/usr/* /usr/
cp -Rfv etc/login.defs /etc/
unset CFLAGS
unset LDFLAGS
alias cp="cp -i"
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | chage, chfn, chgpasswd, chpasswd, chsh, expiry, faillog, getsubids, gpasswd, groupadd, groupdel, groupmems, groupmod, grpck, grpconv, logoutd, newgidmap, newgrp, newuidmap, newusers, passwd, pwck, pwconv, pwunconv, sg, useradd, userdel, usermod, vigr, vipw |
| Temporary Installed Programs | su |
| Installed Libraries | libsubid.so |
| Installed Files | /etc/login.defs |

| Navigation |||
| --- | --- | ---: |
| [<<](./Kmod_32bit.md) Kmod 32-bit | [HOME](../README.md) | Shadow Utils 32-bit [>>](./Shadow_32bit-pass1.md) |

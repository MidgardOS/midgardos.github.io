# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibEConf_32bit.md) LibEConf 32-bit | [HOME](../README.md) | Linux PAM 32-bit [>>](./Linux-PAM_32bit.md) |

## Linux PAM 64-bit

Name: Linux PAM 64-bit<br />
Summary: The Linux implementation of Pluggable Authentication Modules<br />
License: GPL v2.0 or later or BSD 3-Clause<br />
Version: 1.7.1<br />
URL: [https://github.com/linux-pam/linux-pam/releases/download/v1.7.1/Linux-PAM-1.7.1.tar.xz](https://github.com/linux-pam/linux-pam/releases/download/v1.7.1/Linux-PAM-1.7.1.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 5.9 MiB<br />

## Preparation

There are a couple patches available to fix some bugs and add new features to Linux PAM. To apply these changes, run the following commands:

```bash
patch -Np1 -i ../patches/linux-pam/pam-limit-nproc.patch
patch -Np1 -i ../patches/linux-pam/post-v1.7.1.patch
```

## Configuration

**NOTE: This package will need rebuilt after SystemD is installed.**

To configure Linux PAM 64-bit for install into the build root, run the following command:s

```bash
CONFIG_SHELL="${CONFIG_SHELL:-/usr/bin/bash}"; export CONFIG_SHELL
CFLAGS="${CFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CFLAGS
CXXFLAGS="${CXXFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables}"; export CXXFLAGS
FFLAGS="${FFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FFLAGS
FCFLAGS="${FCFLAGS:--O2 -g -m64 -fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables }"; export FCFLAGS
LDFLAGS="${LDFLAGS:-}"; export LDFLAGS

/usr/bin/meson setup --buildtype=plain --prefix=/usr --libdir=/usr/lib64 --libexecdir=/usr/lib64 --bindir=/usr/bin --sbindir=/usr/sbin --includedir=/usr/include --datadir=/usr/share --mandir=/usr/share/man --infodir=/usr/share/info --localedir=/usr/share/locale --sysconfdir=/etc --localstatedir=/var --sharedstatedir=/var/lib --wrap-mode=nodownload --auto-features=enabled \
    -Dvendordir=/etc \
    -Ddocdir=/usr/share/doc/pam-1.7.1 \
    -Dhtmldir=/usr/share/doc/pam-1.7.1/html \
    -Dpdfdir=/usr/share/doc/pam-1.7.1/pdf \
    -Dsecuredir=/usr/lib64/security \
    -Dpam_unix-try-getspnam=true \
    -Dlogind=disabled \
    -Delogind=disabled \
    -Dpwaccess=disabled . x86_64-pc-linux-gnu
```

## Compilation and Installation

To compile Linux PAM 64-bit, run the following command:

```bash
/usr/bin/meson compile -C x86_64-pc-linux-gnu -j4 --verbose
```

Next, run the test suite:

```bash
if [[ ! -d /etc/pam.d ]]; then
    install -v -m755 -d -o root -g root /etc/pam.d
fi
if [[ ! -e /etc/pam.d/other ]]; then
    cat > /etc/pam.d/other << "EOF"
auth     required       pam_deny.so
account  required       pam_deny.so
password required       pam_deny.so
session  required       pam_deny.so
EOF
fi
/usr/bin/meson test -C x86_64-pc-linux-gnu --verbose
```

Finally, to install Linux PAM 64-bit into the build tree, run the following commands:

```bash
/usr/bin/meson install -C x86_64-pc-linux-gnu --no-rebuild
chmod -v 4755 /usr/sbin/unix_chkpwd
install -d -v -m755 -o root -g root /usr/share/doc/pam-1.7.1/modules
cp -fpv x86_64-pc-linux-gnu/modules/pam_*/pam_*.txt /usr/share/doc/pam-1.7.1/modules/
install -d -v -m755 -o root -g root /usr/lib/tmpfiles.d
install -d -v -m755 -o root -g root /etc/security/{limits.d,namespace.d,pam_env.conf.d}

unset CONFIG_SHELL
unset CFLAGS
unset CXXFLAGS
unset FFLAGS
unset FCFLAGS
unset LDFLAGS
```

## Configuration

Linux PAM is a very complex and integral component to system health, and does require proper configuration to ensure both system security and proper operation.

Some of the configuration used in this section will not be final, since other modules and features will be added to the system which will change the default configuration.

Please note that this section is used for _both_ the 64-bit and 32-bit build of PAM.

The default configuration is delivered as part of the `system_files` tree to simplify installation of the configuration. To install the required configuration files, run the commands below:

```bash
rm -fv /etc/pam.d/other
cp -Rv ../system_files/etc/pam.d/* /etc/pam.d/
cp -Rv ../system_files/usr/lib/tmpfiles.d/pam.conf /usr/lib/tmpfiles.d/
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | faillock, mkhomedir_helper, pam_namespace_helper, pam_timestamp_check, pwhistory_helper, unix_chkpwd, and unix_update |
| Installed Libraries | libpam.so, libpam_misc.so, and libpamc.so |
| Installed Plugins | pam_access.so, pam_canonicalize_user.so, pam_debug.so, pam_deny.so, pam_echo.so, pam_env.so, pam_exec.so, pam_faildelay.so, pam_faillock.so, pam_filter.so, pam_ftp.so, pam_group.so, pam_issue.so, pam_keyinit.so, pam_limits.so, pam_listfile.so, pam_localuser.so, pam_loginuid.so, pam_mail.so, pam_mkhomedir.so, pam_motd.so, pam_namespace.so, pam_nologin.so, pam_permit.so, pam_pwhistory.so, pam_rhosts.so, pam_rootok.so, pam_securetty.so, pam_selinux.so, pam_sepermit.so, pam_setquota.so, pam_shells.so, pam_stress.so, pam_succeed_if.so, pam_time.so, pam_timestamp.so, pam_tty_audit.so, pam_umask.so, pam_unix.so, pam_userdb.so, pam_usertype.so, pam_warn.so, pam_wheel.so, and pam_xauth.so |
| Installed Services | pam_namespace.service |
| Installed Files | /etc/environment, /etc/security/access.conf, /etc/security/faillock.conf, /etc/security/group.conf, /etc/security/limits.conf, /etc/security/namespace.conf, /etc/security/pam_env.conf, /etc/security/pwhistory.conf, /etc/security/sepermit.conf, /etc/security/time.conf, /etc/pam.d/common-account, /etc/pam.d/common-auth, /etc/pam.d/common-password, /etc/pam.d/common-session, /etc/pam.d/common-session-nologin, /etc/pam.d/other, /etc/pam.d/postlogin-account, /etc/pam.d/postlogin-auth, /etc/pam.d/postlogin-password, and /etc/pam.d/postlogin-session |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibEConf_32bit.md) LibEConf 32-bit | [HOME](../README.md) | Linux PAM 32-bit [>>](./Linux-PAM_32bit.md) |

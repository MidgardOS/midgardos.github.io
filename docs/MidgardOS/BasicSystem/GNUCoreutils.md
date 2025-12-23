# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Kmod_32bit.md) Kmod 32-bit | [HOME](../README.md) | GNU Diffutils [>>](./GNUDiffutils.md) |

## GNU Coreutils

Name: GNU Coreutils<br />
Summary: Basic file, shell, and text utilities needed for working with any Linux system<br />
License: GPL v3 or later<br />
Version: 9.9<br />
URL: [https://ftp.gnu.org/gnu/coreutils/coreutils-9.9.tar.xz](https://ftp.gnu.org/gnu/coreutils/coreutils-9.9.tar.xz)<br />

Average Build Time: 1.2 SBU<br />
Used Install Space: iB<br />

## Preparation

To fix a bug in POSIX conformance with character boundaries with single and multi byte locales, apply the following patch:

```bash
patch -Np0 -i ../patches/coreutils/coreutils-9.9-i18n-1.patch
```

## Configuration

To configure GNU Coreutils for install into the build root, run the following command:

```bash
autoreconf -fv
automake -af
FORCE_UNSAFE_CONFIGURE=1 \
./configure --prefix=/usr               \
            --libdir=/usr/lib64         \
            --libexecdir=/usr/lib64     \
            --enable-no-install-program=kill,uptime,hostname
```

## Compilation and Installation

To compile GNU Coreutils, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make NON_ROOT_USERNAME=tester check-root
cat > /etc/pam.d/su << "EOF"
#%PAM-1.0
auth            sufficient      pam_rootok.so
auth            substack        common-auth
auth            include         postlogin-auth
account         sufficient      pam_succeed_if.so uid = 0 use_uid quiet
account         include         common-account
password        include         common-password
session         include         common-session
session         include         postlogin-session
session         optional        pam_xauth.so
EOF
useradd -c "Test User" -u 1000 -U -m tester
groupadd -g 102 dummy -U tester
unalias cp
cp -R ../coreutils-9.9 /tmp/
alias cp="cp -i"
chown -R tester /tmp/coreutils-9.9/
pushd /tmp/coreutils-9.9
    su tester -c "PATH=$PATH make -k RUN_EXPENSIVE_TESTS=yes check" < /dev/null
popd
rm -rf /tmp/coreutils-9.9
groupdel -f dummy
userdel -rf tester
```

There are two tests in both the privileged and unprivileged tests that are known to fail while inside a chroot. It is safe to ignore these failures for now.

Note that the `/etc/pam.d/su` configuration is a good temporary configuration for the `su` binary. It will be replaced when Util-Linux is installed.

Finally, to install GNU Coreutils into the build tree, run the following command:

```bash
make install
mv -v /usr/bin/chroot /usr/sbin
mv -v /usr/share/man/man1/chroot.1 /usr/share/man/man8/chroot.8
sed -i 's/"1"/"8"/' /usr/share/man/man8/chroot.8
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs |  |
| Installed Libraries |  |
| Installed Plugins |  |
| Installed Services |  |
| Installed Files |  |

| Navigation |||
| --- | --- | ---: |
| [<<](./Kmod_32bit.md) Kmod 32-bit | [HOME](../README.md) | GNU Diffutils [>>](./GNUDiffutils.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GAWK.md) GNU AWK | [HOME](../README.md) | GNU Roff [>>](./Groff.md) |

## GNU Findutils

Name: GNU Findutils<br />
Summary: CLI tools for searching for files based on various criteria<br />
License: GPL v3 or later<br />
Version: 4.10.0<br />
URL: [https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz](https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz)<br />

Average Build Time: 0.7 SBU<br />
Used Install Space: 3.7 MiB<br />

## Configuration

To configure GNU Findutils for install into the build root, run the following command:

```bash
./configure --prefix=/usr                     \
            --libdir=/usr/lib64               \
            --libexecdir=/usr/lib64/findutils \
            --localstatedir=/var/lib/locate   \
            --docdir=/usr/share/doc/findutils-4.10.0
```

## Compilation and Installation

To compile GNU Findutils, run the following command:

```bash
make
```

Next, run the test suite:

```bash
if [[ ! -f /etc/pam.d/su ]]; then
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
fi
useradd -c "Test User" -u 1000 -U -m tester
groupadd -g 102 dummy -U tester
unalias cp
cp -R ../findutils-4.10.0 /tmp/
alias cp="cp -i"
chown -R tester /tmp/findutils-4.10.0/
pushd /tmp/findutils-4.10.0
    su tester -c "PATH=$PATH && autoreconf -fiv && make check"
popd
rm -rf /tmp/findutils-4.10.0
groupdel -f dummy
userdel -rf tester
```

Finally, to install GNU Findutils into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | find, locate, updatedb, and xargs |
| Installed Plugins | frcode |

| Navigation |||
| --- | --- | ---: |
| [<<](./GAWK.md) GNU AWK | [HOME](../README.md) | GNU Roff [>>](./Groff.md) |

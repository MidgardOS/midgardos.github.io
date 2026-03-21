# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUDiffutils.md) GNU Diffutils | [HOME](../README.md) | GNU Findutils [>>](./GNUFindutils.md) |

## GNU AWK

Name: GNU AWK<br />
Summary: GNU Awk is a CLI tool for processing text data<br />
License: GPL v3 or later<br />
Version: 5.4.0<br />
URL: [https://ftp.gnu.org/gnu/gawk/gawk-5.4.0.tar.xz](https://ftp.gnu.org/gnu/gawk/gawk-5.4.0.tar.xz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 8.3 MiB<br />

## Configuration

To configure GNU AWK for install into the build root, run the following commands:

```bash
sed -i 's/extras//' Makefile.in
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --docdir=/usr/share/doc/gawk-5.4.0
```

## Compilation and Installation

To compile GNU AWK, run the following command:

```bash
make
```

Next, run the test suite:

```bash
useradd -c "Test User" -u 1000 -U -m tester
groupadd -g 102 dummy -U tester
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
unalias cp
cp -Rv ../gawk-5.4.0 /tmp/
alias cp="cp -i"
chown -Rv tester /tmp/gawk-5.4.0/
pushd /tmp/gawk-5.4.0
    su tester -c "PATH=$PATH && autoreconf -fiv && make check"
popd
rm -rfv /tmp/gawk-5.4.0
groupdel -f dummy
userdel -rf tester
```

There are three failing tests due to multibyte parsing issues. These will be resolved in a later build.

Finally, to install GNU AWK into the build tree, run the following command:

```bash
rm -f /usr/bin/gawk-5.4.0
make install
ln -sv gawk.1 /usr/share/man/man1/awk.1
install -v -D -m644 -o root -g root doc/{awkforai.txt,*.{eps,pdf,jpg}} -t /usr/share/doc/gawk-5.4.0
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | awk, gawk, gawk-5.3.2, and gawkbug |
| Installed Libraries | filefuncs.so, fnmatch.so, fork.so, inplace.so, intdiv.so, ordchr.so, readdir.so, readfile.so, revoutput.so, revtwoway.so, rwarray.so, and time.so |
| Installed Plugins | grcat and pwcat |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUDiffutils.md) GNU Diffutils | [HOME](../README.md) | GNU Findutils [>>](./GNUFindutils.md) |

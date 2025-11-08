# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./hostname.md) Hostname | [HOME](../README.md) | GNU Less [>>](./GNULess.md) |

## GNU InetUtils

Name: GNU InetUtils<br />
Summary: Various network utilities<br />
License: GPL v3 or later<br />
Version: 2.6<br />
URL: [https://ftp.gnu.org/gnu/inetutils/inetutils-2.6.tar.xz](https://ftp.gnu.org/gnu/inetutils/inetutils-2.6.tar.xz)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: <br />

## Configuration

To configure GNU InetUtils for install into the build root, run the following commands:

```bash
sed -i 's/def HAVE_TERMCAP_TGETENT/ 1/' telnet/telnet.c
./configure --prefix=/usr                           \
            --bindir=/usr/bin                       \
            --libdir=/usr/lib64                     \
            --libexecdir=/usr/lib64                 \
            --docdir=/usr/share/doc/inetutils-2.6   \
            --disable-rexecd                        \
            --disable-rlogind                       \
            --disable-rshd                          \
            --disable-logger                        \
            --disable-syslogd                       \
            --disable-telnetd                       \
            --disable-rcp                           \
            --disable-rexec                         \
            --disable-rlogin                        \
            --disable-rsh                           \
            --disable-uucpd                         \
            --disable-ifconfig                      \
            --disable-hostname                      \
            --disable-dnsdomainname                 \
            --disable-traceroute                    \
            --disable-whois
```

## Compilation and Installation

To compile GNU InetUtils, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU InetUtils into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
install -dvm 755 -o root -g root DESTDIR/usr/sbin
mv -v DESTDIR/usr/lib64/* DESTDIR/usr/sbin/
install -dvm 755 -o root -g root /usr/lib/systemd/system
cp -Rv DESTDIR/usr/bin/* /usr/bin/
cp -Rv DESTDIR/usr/sbin/* /usr/sbin/
cp -Rv DESTDIR/usr/share/* /usr/share/
```

Note: the systemd unit files are not installed for the services that have been installed for now. This will be completed in a later revision of this package.

## Contents

| Contents | |
| --- | --- |
| Installed Programs | ftp ping ping6 talk telnet tftp |
| Installed Services | ftpd inetd talkd tftpd |

| Navigation |||
| --- | --- | ---: |
| [<<](./hostname.md) Hostname | [HOME](../README.md) | GNU Less [>>](./GNULess.md) |

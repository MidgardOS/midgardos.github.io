# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libnsl232bit.md) libnsl2 32-bit | [HOME](../README.md) | TCP Wrappers 32-bit [>>](./tcp_wrappers32bit.md) |

## TCP Wrappers 64-bit

Name: TCP Wrappers 64-bit<br />
Summary: Access and authorization wrapper for TCP daemons<br />
License: BSD<br />
Version: 7.6<br />
URL: [https://github.com/tcp-wrappers/code/archive/refs/tags/tcp_wrappers_7.6-ipv6.4.tar.gz](https://github.com/tcp-wrappers/code/archive/refs/tags/tcp_wrappers_7.6-ipv6.4.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 312 KiB<br />

## Patching

The TCP Wrappers sources need a number of patches to work correctly on modern systems. Please run the following commands to patch the sources.

These patches were borrowed from the Fedora Linux project.

**Note: the patches should be downloaded to their own `patches/tcp_wrappers` directory under the `/MidgardOS/sources` directory.**

```bash
cd /sources/code-tcp_wrappers_7.6-ipv6.4
patch -p1 <../patches/tcp_wrappers/tcpw7.2-config.patch
patch -p1 <../patches/tcp_wrappers/tcpw7.2-setenv.patch
patch -p1 <../patches/tcp_wrappers/tcpw7.6-netgroup.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-bug11881.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-bug17795.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-bug17847.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-fixgethostbyname.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-docu.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-man.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers.usagi-ipv6.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-shared.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-sig.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-ldflags.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-fix_sig-bug141110.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-162412.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-220015.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-siglongjmp.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-sigchld.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-196326.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers_7.6-249430.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-inetdconf.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-bug698464.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-xgets.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-initgroups.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-warnings.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-uchart_fix.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-altformat.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-aclexec.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-inetcf-c99.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-gcc15-errors.patch
patch -p1 <../patches/tcp_wrappers/tcp_wrappers-7.6-gcc15-warnings.patch
```

## Configuration

The TCP Wrappers sources do not have a traditional configuration script.

## Compilation and Installation

To compile TCP Wrappers 64-bit, run the following command:

```bash
make REAL_DAEMON_DIR=/usr/sbin MAJOR=0 MINOR=7 REL=6 linux
```

Finally, to install TCP Wrappers 64-bit into the build tree, run the following commands:

```bash
BUILD_ROOT=DESTDIR
rm -rfv ${BUILD_ROOT}
mkdir -pv ${BUILD_ROOT}/usr/include
mkdir -pv ${BUILD_ROOT}/usr/lib64
mkdir -pv ${BUILD_ROOT}/usr/share/man/man{3,5,8}
mkdir -pv ${BUILD_ROOT}/usr/sbin
install -pv -m644 hosts_access.3 ${BUILD_ROOT}/usr/share/man/man3
install -pv -m644 hosts_access.5 hosts_options.5 ${BUILD_ROOT}/usr/share/man/man5
install -pv -m644 tcpd.8 tcpdmatch.8 safe_finger.8 try-from.8 ${BUILD_ROOT}/usr/share/man/man8
ln -sfv hosts_access.5 ${BUILD_ROOT}/usr/share/man/man5/hosts.allow.5
ln -sfv hosts_access.5 ${BUILD_ROOT}/usr/share/man/man5/hosts.deny.5
cp -Rv libwrap.so* ${BUILD_ROOT}/usr/lib64/
install -pv -m644 tcpd.h ${BUILD_ROOT}/usr/include/
install -v -m755 safe_finger ${BUILD_ROOT}/usr/sbin/
install -v -m755 tcpd ${BUILD_ROOT}/usr/sbin/
install -v -m755 try-from ${BUILD_ROOT}/usr/sbin/
install -v -m755 tcpdmatch ${BUILD_ROOT}/usr/sbin/
cp -Rv DESTDIR/usr/* /usr/
```

Next, install the hosts.allow file in `/etc`:

```bash:
cat <<EOF > /etc/hosts.allow
# /etc/hosts.allow
# Make sure package tcpd is installed on your system for this to work.
# See 'man tcpd' and 'man 5 hosts_access' for a detailed description
# of /etc/hosts.allow and /etc/hosts.deny.
#
# short overview about daemons and servers that are built with
# tcp_wrappers support:
# 
# package name  | daemon path           |       token
# ----------------------------------------------------------------------------
# ssh, openssh  | /usr/sbin/sshd        |  sshd, sshd-fwd-x11, sshd-fwd-<port>
# quota         | /usr/sbin/rpc.rquotad |  rquotad
# tftpd         | /usr/sbin/in.tftpd    |  in.tftpd
# portmap       | /sbin/portmap         |  portmap
#                       The portmapper does not verify against hostnames
#                       to prevent hangs. It only checks non-local addresses.
# 
# (kernel nfs server)
# nfs-utils     | /usr/sbin/rpc.mountd  |  mountd
# nfs-utils     | /sbin/rpc.statd       |  statd
#
# (unfsd, userspace nfs server)
# nfs-server    | /usr/sbin/rpc.mountd  |  rpc.mountd
# nfs-server    | /usr/sbin/rpc.ugidd   |  rpc.ugidd
#
# (printing services)
# lprng         | /usr/sbin/lpd         |  lpd
# cups          | /usr/sbin/cupsd       |  cupsd
#                       The cupsd server daemon reports to the cups
#                       error logs, not to the syslog(3) facility.
#
# (Uniterrupted Power Supply Software)
# apcupsd       | /sbin/apcupsd         |  apcupsd
# apcupsd       | /sbin/apcnisd         |  apcnisd
# 
# All of the other network servers such as samba, apache or X, have their own
# access control scheme that should be used instead.
#
# In addition to the services above, the services that are started on request 
# by inetd or xinetd use tcpd to "wrap" the network connection. tcpd uses
# the last component of the server pathname as a token to match a service in
# /etc/hosts.{allow,deny}. See the file /etc/inetd.conf for the token names.
# The following examples work when uncommented:
#
# Example 1: Fire up a mail to the admin if a connection to the printer daemon
# has been made from host foo.bar.com, but simply deny all others:
# lpd : foo.bar.com : spawn /bin/echo "%h printer access" | #                               mail -s "tcp_wrappers on %H" root
# 
# Example 2: grant access from local net, reject with message from elsewhere.
# in.telnetd : ALL EXCEPT LOCAL : ALLOW
# in.telnetd : ALL : #    twist /bin/echo -e "\n\raccess from %h declined.\n\rGo away.";sleep 2
#
# Example 3: run a different instance of rsyncd if the connection comes 
#            from network 172.20.0.0/24, but regular for others:
# rsyncd : 172.20.0.0/255.255.255.0 : twist /usr/local/sbin/my_rsyncd-script
# rsyncd : ALL : ALLOW

EOF
```

Finally, install the hosts.deny file in `/etc/`:

```bash
cat <<EOF > /etc/hosts.deny
# /etc/hosts.deny
# Make sure package tcpd is installed on your system for this to work.
# See 'man tcpd' and 'man 5 hosts_access' as well as /etc/hosts.allow
# for a detailed description.

http-rman : ALL EXCEPT LOCAL

EOF
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | safe_finger, tcpd, try-from, and tcpdmatch |
| Installed Libraries | libwrap.so |
| Installed Files | /etc/hosts.allow and /etc/hosts.deny |

| Navigation |||
| --- | --- | ---: |
| [<<](./libnsl232bit.md) libnsl2 32-bit | [HOME](../README.md) | TCP Wrappers 32-bit [>>](./tcp_wrappers32bit.md) |

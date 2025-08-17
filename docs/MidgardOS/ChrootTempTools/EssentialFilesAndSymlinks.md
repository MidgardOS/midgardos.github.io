# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./EnteringChroot.md) Entering the Chroot | [HOME](../README.md) | GNU Gettext [>>](./GNUGettext.md) |

## Extra Directories

There are a few more directories that are required to meet the Linux Filesystem Hierarchy Standard for LSB conformance. Create them now. Note that the install command below is correct since there is no passwd or group file to define the root user.

```bash
install -d -v -m755 -o 0 -g 0 /etc/{default,opt,sysconfig}
install -d -v -m755 -o 0 -g 0 /lib/firmware
install -d -v -m755 -o 0 -g 0 /media/{floppy,cdrom}
install -d -v -m755 -o 0 -g 0 /usr/lib/locale
install -d -v -m755 -o 0 -g 0 /usr/local/share/{color,dict,dot,info,locale,man,misc,terminfo,zoneinfo}
install -d -v -m755 -o 0 -g 0 /usr/local/share/man/man{1..8}
install -d -v -m755 -o 0 -g 0 /usr/share/{color,dict,dot,info,locale,man,misc,terminfo,zoneinfo}
install -d -v -m755 -o 0 -g 0 /usr/share/man/man{1..8}
install -d -v -m755 -o 0 -g 0 /var/lib/{color,misc,locate}
```

## Essential Files and Symbolic Links

For the build root to work optimally, there are a few files that should be created now. Historically, Linux used a file in `/etc` for keeping tabs on mounted volumes. This is now handled by the kernel directly and is exported to user-space using `/proc/mounts`. To meet the needs of some older utilities that still reference `/etc/mtab`, create a symlink for it now:
```bash
ln -sv /proc/self/mounts /etc/mtab
```

Next, create a basic `/etc/hosts` file that will be referenced in some application test suites. and in Perl's configuration files:
```bash
cat > /etc/hosts << EOF
127.0.0.1  localhost
::1        localhost
fe00::0    ipv6-localnet
ff00::0    ipv6-mcastprefix
ff02::1    ipv6-allnodes
ff02::2    ipv6-allrouters
ff02::3    ipv6-allhosts
EOF
```

When MidgardOS is normally installed via RPMs, the system users and groups are delivered via RPM packages that populate `/etc/passwd` and `/etc/group`. As those are not available at this time, these will be populated via the following commands:
```bash
cat > /etc/passwd << "EOF"
root:x:0:0:root:/root:/bin/bash
bin:x:1:1:bin:/dev/null:/usr/bin/false
daemon:x:2:2:Daemon User:/dev/null:/usr/bin/false
adm:x:3:4:adm:/var/adm:/usr/bin/false
sync:x:5:0:sync:/sbin:/bin/sync
messagebus:x:18:18:D-Bus Message Daemon User:/run/dbus:/usr/bin/false
systemd-journal-gateway:x:73:73:systemd Journal Gateway:/:/usr/bin/false
systemd-journal-remote:x:74:74:systemd Journal Remote:/:/usr/bin/false
systemd-journal-upload:x:75:75:systemd Journal Upload:/:/usr/bin/false
systemd-network:x:76:76:systemd Network Management:/:/usr/bin/false
systemd-resolve:x:77:77:systemd Resolver:/:/usr/bin/false
systemd-timesync:x:78:78:systemd Time Synchronization:/:/usr/bin/false
systemd-coredump:x:79:79:systemd Core Dumper:/:/usr/bin/false
uuidd:x:80:80:UUID Generation Daemon User:/dev/null:/usr/bin/false
systemd-oom:x:81:81:systemd Out Of Memory Daemon:/:/usr/bin/false
nobody:x:65534:65534:Unprivileged User:/dev/null:/usr/bin/false
EOF

cat > /etc/group << "EOF"
root:x:0:
bin:x:1:daemon
sys:x:2:
kmem:x:3:
tape:x:4:
tty:x:5:
daemon:x:6:
floppy:x:7:
disk:x:8:
lp:x:9:
dialout:x:10:
audio:x:11:
video:x:12:
utmp:x:13:
cdrom:x:15:
adm:x:16:
messagebus:x:18:
systemd-journal:x:23:
input:x:24:
mail:x:34:
kvm:x:61:
systemd-journal-gateway:x:73:
systemd-journal-remote:x:74:
systemd-journal-upload:x:75:
systemd-network:x:76:
systemd-resolve:x:77:
systemd-timesync:x:78:
systemd-coredump:x:79:
uuidd:x:80:
systemd-oom:x:81:
wheel:x:97:
users:x:999:
nogroup:x:65534:
EOF
```

Note that the `root` account password will be set later.

To remove the "I have no name!" prompt, reload the shell by running the following command:
```bash
exec /usr/bin/bash --login
```

Finally, to ensure that `login`, `agetty`, and `init` commands can properly use certain log files, they need initialized, and their proper permissions set:
```bash
touch /var/log/{btmp,lastlog,faillog,wtmp}
chgrp -v utmp /var/log/lastlog
chmod -v 664  /var/log/lastlog
chmod -v 600  /var/log/btmp
```

| Navigation |||
| --- | --- | ---: |
| [<<](./EnteringChroot.md) Entering the Chroot | [HOME](../README.md) | GNU Gettext [>>](./GNUGettext.md) |

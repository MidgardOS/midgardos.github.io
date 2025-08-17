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
```

| Navigation |||
| --- | --- | ---: |
| [<<](./EnteringChroot.md) Entering the Chroot | [HOME](../README.md) | GNU Gettext [>>](./GNUGettext.md) |

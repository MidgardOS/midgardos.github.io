# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./EnteringChroot.md) Entering the Chroot | [HOME](../README.md) | Validating the C and C++ Compilers [>>](./ValidatingCompilers.md) |

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
install -d -v -m755 -o 0 -g 0 /var/lib/{color,hwclock,misc,locate}
```

## Essential Files and Symbolic Links

For the build root to work optimally, there are a few files that should be created now. Historically, Linux used a file in `/etc` for keeping tabs on mounted volumes. This is now handled by the kernel directly and is exported to user-space using `/proc/mounts`. To meet the needs of some older utilities that still reference `/etc/mtab`, create a symlink for it now:

```bash
ln -sv /proc/self/mounts /etc/mtab
```

Next, create a basic `/etc/hosts` file that will be referenced in some application test suites. and in Perl's configuration files:
```bash
install -v -m644 -o root -g root ../system_files/etc/hosts /etc
```

When MidgardOS is normally installed via RPMs, the system users and groups are delivered via RPM packages that populate `/etc/passwd` and `/etc/group`. As those are not available at this time, these will be populated via the following commands:

```bash
install -v -m644 -o root -g root ../system_files/etc/passwd /etc/
install -v -m644 -o root -g root ../system_files/etc/group /etc/
install -v -m644 -o root -g root ../system_files/etc/netgroup /etc/
install -v -m644 -o root -g root ../system_files/etc/ethers /etc/
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
| [<<](./EnteringChroot.md) Entering the Chroot | [HOME](../README.md) | Validating the C and C++ Compilers [>>](./ValidatingCompilers.md) |

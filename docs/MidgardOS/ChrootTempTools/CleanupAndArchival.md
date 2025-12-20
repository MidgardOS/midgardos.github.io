# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./UtilLinux32bit.md) Util-Linux 32-bit | [HOME](../README.md) | Building the Base System Tools [>>](../BasicSystem/Overview.md) |

## Clean-up

To ensure that documentation is from the final builds before being packaged into RPMs, remove them now:

```bash
rm -rf /usr/share/{info,man,doc}/*
```

Next, since modern Linux systems don't use `libltdl` to wrap the calls from `/lib/ld.so`, and to prevent build errors from other tools, remove all libtool archive files:

```bash
find /usr/{lib,lib64} -name \*.la -delete
```

Now, we no longer require the `/tools` tree, so remove it:

```bash
rm -rf /tools
```

## Archiving the Build Root

At this point, all of the core and essential libraries and applications have been installed, and the build root is in a good state for building the Base System and beyond. It is a good idea to back up and archive the build root at this point to allow restoring from backup if any future package causes failures.

The backup must be done outside the chroot environment. To leave the chroot environment, run the following command:

```bash
exit
```

Before creating the archive, unmount the kernel Virtual Filesystems:

```bash
BRFS=/MidgardOS
for mp in "dev/shm" "dev/pts" "sys" "proc" "run" "dev"; do
  if mountpoint -q "$BRFS/$mp"; then
    umount -v "$BRFS/$mp"
  fi
done
```

Now that the kernel Virtual Filesystems are unmounted, create the archive of the build root with the following commands:

```bash
BRFS=/MidgardOS
cd $BRFS && tar -cJpf $HOME/MidgardOS-temp-tools-2025.0.tar.xz .
```

| Navigation |||
| --- | --- | ---: |
| [<<](./UtilLinux32bit.md) Util-Linux 32-bit | [HOME](../README.md) | Building the Base System Tools [>>](../BasicSystem/Overview.md) |

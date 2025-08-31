# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ChangingOwnerships.md) Changing Ownerships | [HOME](../README.md) | Entering the Chroot [>>](./EnteringChroot.md) |

## Preparing the Virtual Kernel Filesystems

One of the ways applications use to communicate with the Linux kernel is through various Virtual File Systems exposed into user-space. These filesystems expose device access, process information, and system data, among many more. To ensure that they are able to do so from within the chroot environment is to mount them inside the environment.

First, as the `root` user, ensure that the mountpoints required exist:
```bash
BRFS=/MidgardOS
for DIR in "dev" "proc" "sys" "run"; do
    install -d -v -m 755 -o root -g root $BRFS/$DIR
done
```

## Mounting the Virtual Kernel Filesystems

Now that the mountpoints have been created, run the following commands to mount the Virtual Kernel Filesystems into the build root tree paths:

```bash
BRFS=/MidgardOS
mount -v -t devtmpfs devtmpfs $BRFS/dev/
mount -v -t devpts devpts -o gid=5,mode=620 $BRFS/dev/pts/
mount -v -t proc proc $BRFS/proc/
mount -v -t sysfs sysfs $BRFS/sys/
mount -v -t tmpfs tmpfs $BRFS/run/
if [ -h $BRFS/dev/shm ]; then
  install -v -d -m 1777 ${BRFS}$(realpath /dev/shm)
else
  mount -v -t tmpfs -o mode=1777,nosuid,nodev,inode64 tmpfs $BRFS/dev/shm/
fi
```

| Navigation |||
| --- | --- | ---: |
| [<<](./ChangingOwnerships.md) Changing Ownerships | [HOME](../README.md) | Entering the Chroot [>>](./EnteringChroot.md) |

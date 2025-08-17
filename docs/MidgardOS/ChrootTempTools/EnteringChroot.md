# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./MountingVirtualKernelFilesystems.md) Mounting Virtual Kernel Filesystems | [HOME](../README.md) | Creating Essential Files and Symlinks [>>](./EssentialFilesAndSymlinks.md) |

## Entering the Chroot

When entering the chroot environment, it is essential that all required virtual filesystems are mounted before entry, and that the entry script is run as the `root` user:

```bash
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root. Exiting"
    exit 1
fi

BRFS=/MidgardOS
if ! mountpoint -q $BRFS/dev; then
    echo "$BRFS/dev is not mounted. Mounting the devtmpfs filesystem"
    mount -vt devtmpfs devtmpfs $BRFS/dev
fi
if ! mountpoint -q $BRFS/dev/pts; then
    echo "$BRFS/dev/pts is not mounted. Mounting devpts filesystem"
    mount -vt devpts devpts -o gid=5,mode=620 $BRFS/dev/pts
fi
if ! mountpoint -q $BRFS/proc; then
    echo "$BRFS/proc is not mounted. Mounting proc filesystem"
    mount -vt proc proc $BRFS/proc
fi
if ! mountpoint -q $BRFS/sys; then
    echo "$BRFS/sys is not mounted. Mounting sysfs filesystem"
    mount -vt sysfs sysfs $BRFS/sys
fi
if ! mountpoint -q $BRFS/run; then
    echo "$BRFS/run is not mounted. Mounting tmpfs filesystem"
    mount -vt tmpfs tmpfs $BRFS/run
fi
if [ -h $BRFS/dev/shm ] && ! mountpoint -q $BRFS/dev/shm; then
    echo "$BRFS/dev/shm is a symlink. Ensuring target directory exists"
    install -v -d -m 1777 ${BRFS}$(realpath /dev/shm)
elif ! mountpoint -q $BRFS/dev/shm; then
    echo "$BRFS/dev/shm is not mounted. Mounting tmpfs filesystem"
    mount -vt tmpfs -o mode=1777,nosuid,nodev,inode64 tmpfs $BRFS/dev/shm
fi

chroot "$BRFS" /usr/bin/env -i         \
    HOME=/root                         \
    TERM="$TERM"                       \
    PS1='(buildroot chroot) \u:\w\$ '  \
    PATH=/usr/bin:/usr/sbin            \
    MAKEFLAGS="-j$(nproc)"             \
    TESTSUITEFLAGS="-j$(nproc)"        \
    /bin/bash --login
```

Once this series of commands has been run, you should successfully be inside the chroot'd build environment. To know for sure, you should see the following in your terminal:
```
# ./enter_chroot.sh
(buildroot chroot) I have no name!:/#
```

| Navigation |||
| --- | --- | ---: |
| [<<](./MountingVirtualKernelFilesystems.md) Mounting Virtual Kernel Filesystems | [HOME](../README.md) | Creating Essential Files and Symlinks [>>](./EssentialFilesAndSymlinks.md) |

# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ChangingOwnerships.md) Changing Ownerships | [HOME](../README.md) | Entering the Chroot [>>](./EnteringChroot.md) |

## Preparing the Virtual Kernel Filesystems

One of the ways applications use to communicate with the Linux kernel is through various virtual filesystems exposed into user-space. These filesystems expose device access, process information, and system data, among many others. To ensure that they are able to do so from within the chroot environment, we must mount them inside it.

## Mounting the Virtual Kernel Filesystems

The mountpoints were created earlier when the directories and various symbolic links were created. Now you can run the following commands as the `root` user to create a script and mount the Virtual Kernel Filesystems into the build root tree paths:

```bash
cat <<'EOF' > /root/mountvfs.sh
#!/bin/bash

set -euo pipefail

function check_mounted {
    local mountpoint=$1

    retval=0
    mountpoint $mountpoint >/dev/null || retval=$?

    echo $retval
}

BRFS=/MidgardOS
if [[ "$(check_mounted $BRFS/dev)" != 32 ]] ; then
    mount -v -t devtmpfs devtmpfs $BRFS/dev/
else
    echo "$BRFS/dev already mounted. Skipping"
fi
if [[ "$(check_mounted $BRFS/dev/pts)" != 32 ]] ; then
    mount -v -t devpts devpts -o gid=5,mode=620 $BRFS/dev/pts/
else
    echo "$BRFS/dev/pts already mounted. Skipping"
fi
if [[ "$(check_mounted $BRFS/proc)" != 32 ]]; then
    mount -v -t proc proc $BRFS/proc/
else
    echo "$BRFS/proc already mounted. Skipping"
fi
if [[ "$(check_mounted $BRFS/sys)" != 32 ]]; then
    mount -v -t sysfs sysfs $BRFS/sys/
else
    echo "$BRFS/sys already mounted. Skipping"
fi
if [[ "$(check_mounted $BRFS/run)" != 32 ]]; then
    mount -v -t tmpfs tmpfs $BRFS/run/
else
    echo "$BRFS/run already mounted. Skipping"
fi
if [[ "$(check_mounted $BRFS/dev/shm)" != 32 ]]; then
    if [[ -h $BRFS/dev/shm ]]; then
        install -v -d -m 1777 ${BRFS}$(realpath /dev/shm)
    else
        mount -v -t tmpfs -o mode=1777,nosuid,nodev,inode64 tmpfs $BRFS/dev/shm/
    fi
else
    echo "$BRFS/dev/shm already mounted. Skipping"
fi
EOF
chmod 755 /root/mountvfs.sh
./mountvfs.sh
```

This script is useful for preparing to re-enter the build root later.

| Navigation |||
| --- | --- | ---: |
| [<<](./ChangingOwnerships.md) Changing Ownerships | [HOME](../README.md) | Entering the Chroot [>>](./EnteringChroot.md) |

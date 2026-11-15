# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./MountingVirtualKernelFilesystems.md) Mounting Virtual Kernel Filesystems | [HOME](../README.md) | Creating Essential Files and Symlinks [>>](./EssentialFilesAndSymlinks.md) |

## Entering the Chroot

When entering the chroot environment, it is essential that all required virtual filesystems are mounted before entry, and that the entry script is run as the `root` user. To do so, create a shell script with the following contents:

```bash
cat << 'EOF' > /root/enter_chroot.sh
#!/bin/bash

if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root. Exiting"
    exit 1
fi

BRFS=/MidgardOS
/root/mountvfs.sh

chroot "$BRFS" /usr/bin/env -i         \
    HOME=/root                         \
    TERM="$TERM"                       \
    PS1='(buildroot chroot) \u:\w\$ '  \
    PATH=/usr/bin:/usr/sbin            \
    MAKEFLAGS="-j$(nproc)"             \
    TESTSUITEFLAGS="-j$(nproc)"        \
    /bin/bash --login
EOF
chmod 755 enter_chroot.sh
./enter_chroot.sh
```

Once this series of commands has been run, you should successfully be inside the chroot'd build environment. To know for sure, you should see the following in your terminal:
```
# ./enter_chroot.sh
(buildroot chroot) I have no name!:/#
```

| Navigation |||
| --- | --- | ---: |
| [<<](./MountingVirtualKernelFilesystems.md) Mounting Virtual Kernel Filesystems | [HOME](../README.md) | Creating Essential Files and Symlinks [>>](./EssentialFilesAndSymlinks.md) |

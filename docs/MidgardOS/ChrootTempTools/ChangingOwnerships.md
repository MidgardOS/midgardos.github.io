# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ChrootBuildingTempTools.md) Rationale | [HOME](../README.md) | Mounting Virtual Kernel Filesystems [>>](./MountingVirtualKernelFilesystems.md) |

## Rationale

At this point, many of the files and directories under the build root are owned by the `builder` user, not `root`. The `builder` user, while it will be created inside the build tree for later use, should not have group or user ownership of core directories and files inside a normal functioning system.

To correct this, run the following commands to remove any POSIX ACLs on directories, and set the group and user ownerships to the `root` user and group:

```bash
BRFS=/MidgardOS
pushd $BRFS
sudo setfacl -bPR ./
sudo chown --from builder -hRv root:root ./
popd
```

| Navigation |||
| --- | --- | ---: |
| [<<](./ChrootBuildingTempTools.md) Rationale | [HOME](../README.md) | Mounting Virtual Kernel Filesystems [>>](./MountingVirtualKernelFilesystems.md) |

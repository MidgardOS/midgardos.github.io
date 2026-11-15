# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection - pass 2 | [HOME](../README.md) | Changing Ownership [>>](./ChangingOwnerships.md) |

## Rationale

At this point of the build, all of the circular dependencies of the core operating system tools have been dealt with. To better isolate the build root from the host operating system, we'll enter a "Changed Root", which is called a "chroot" environment, where the only element from the host operating system visible from within is the Linux kernel that it runs.

To make this work, the ownerships on files and directories inside the build root will be modified to match how a normal Linux system is set, and some virtual kernel filesystems used to communicate information from the kernel to user-space will be mounted under the build root tree.

All commands from here on out will be run as the `root` user, so caution should be taken to avoid breaking the environment that has been built so far.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection - pass 2 | [HOME](../README.md) | Changing Ownership [>>](./ChangingOwnerships.md) |

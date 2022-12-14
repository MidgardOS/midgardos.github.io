| Navigation |||
| --- | --- | ---: |
| [<<](./CreatePartition.md) | [HOME](./README.md) | [>>](./FormatAndMount.md) |

MidgardOS is a traditional Linux distribution, however it follows some modern aspects that collapse many directories in the `/` level of the filesystem tree into `/usr`. This step will add the mount point directories on the host system, and the basic directory structure for the core OS along with the tools and cross-tools directory trees. Additionally, we'll format the new partitions we created in the last step with the requisite filesystems and add them to the host system's `/etc/fstab` file for mount at boot of the virtual machine.

## Format the New EFI Partition

MidgardOS is designed to run on UEFI based hardware. This requires an EFI boot volume to allow the system to boot the operating system. Additionally, MidgardOS uses systemd-boot as the boot loader. This means that the normal EFI partition should be mounted to the MidgardOS `/boot` directory.

To create the EFI boot filesystem on the 512MiB EFI partition that was created in the last step, run the following command:

```sh
/sbin/mkfs.vfat -F 32 -n EFI --verbose /dev/sdb1
```

## Format the MidgardOS Root Filesystem

MidgardOS by default uses the SGI XFS filesystem for the root disk format. This filesystem has excellent performance for small and medium sized files, and has reasonably good journalling performance and capabilities to protect against data loss under powerloss occurances. To format the new root partition, run the following command in the virtual machine as a privileged user:

```bash
/sbin/mkfs.xfs -m bigtime=1,crc=1,finobt=1,inobtcount=1 -i align=1 -L MidgardOS /dev/sdb2
```

## Create the Mount Points for the New Filesystems

To mount the newly formatted volumes, we first need to create the directories that they will be mounted to. Part of the directory creations cannot be done until the root file system is mounted.

First, create the mount point for the new root file system:

```bash
install -v -d -m 755 -o root -g root /MidgardOS
```

Now, mount the new root filesystem:

```bash
mount -v -t xfs -o rw,grpid,prjquota,attr2,inode64 -L MidgardOS /MidgardOS
```

Now that the new root filesystem is mounted, create the mount point for the EFI partition:

```bash
install -v -d -m 755 -o root -g root /MidgardOS/boot
```

Now, mount the EFI partition to the new directory:

```bash
FAT_MOUNT_DIRECTIVES="fmask=0022,dmask=0022,codepage=437,iocharset=utf8,discard,showexec,sys_immutable,rodir,shortname=mixed"
mount -t vfat \
      -o rw,$FAT_MOUNT_DIRECTIVES \
      -L EFI /MidgardOS/boot
```

## Creating the Basic Directory Tree



| Navigation |||
| --- | --- | ---: |
| [<<](./CreatePartition.md) | [HOME](./README.md) | [>>](./FormatAndMount.md) |
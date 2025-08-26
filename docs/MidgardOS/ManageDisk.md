# Section 1 - Preparation

| Navigation |||
| --- | --- | ---: |
| [<<](./CreatePartition.md) Partition Creation | [HOME](./README.md) | Directories and Symlinks [>>](./DirectoriesAndSymlinks.md) |

MidgardOS is a traditional Linux distribution, however it follows some modern aspects that collapse many directories in the `/` level of the filesystem tree into `/usr`. This step will add the mount point directories on the host system, Additionally, we'll format the new partitions we created in the last step with the requisite filesystems and add them to the host system's `/etc/fstab` file for mount at boot of the virtual machine.

## Format the New EFI Partition

MidgardOS is designed to run on UEFI based hardware. This requires an EFI boot volume to allow the system to boot the operating system. Additionally, MidgardOS uses systemd-boot as the boot loader. This means that the normal EFI partition should be mounted to the MidgardOS `/boot` directory.

To create the EFI boot filesystem on the 512MiB EFI partition that was created in the last step, run the following command:

```sh
/sbin/mkfs.vfat -F 32 -n EFI -v /dev/nvme0n1p1
```

## Format the MidgardOS Root Filesystem

MidgardOS by default uses the SGI XFS filesystem for the root disk format. This filesystem has excellent performance for small and medium sized files, and has reasonably good journalling performance and capabilities to protect against data loss under powerloss occurances. To format the new root partition, run the following command in the virtual machine as a privileged user:

```bash
/sbin/mkfs.xfs -m bigtime=1,crc=1,finobt=1,inobtcount=1 -L MidgardOS /dev/nvme0n1p2
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
mount -v \
      -t vfat \
      -o rw,$FAT_MOUNT_DIRECTIVES \
      -L EFI /MidgardOS/boot
```

## Make the Volumes Available When the Host Reboots

To make it easier to access the new volumes and ensure they are available after rebooting the host operating system, we'll add them either to the host's `/etc/fstab` or as a systemd mount unit.

If you're more comfortable with using the host's `/etc/fstab`, run the following command to add the appropriate lines to it:

```bash
FAT_PARAMS="fmask=0022,dmask=0022,codepage=437,iocharset=utf8,discard,showexec,sys_immutable,rodir,shortname=mixed"
EFI_MOUNT_PARAMS="rw,${FAT_PARAMS},x-systemd.requires-mount-for=/MidgardOS"
XFS_MOUNT_PARAMS="rw,grpid,attr2,inode64,usrquota,prjquota"
echo -e "LABEL=MidgardOS\t/MidgardOS\txfs\t${XFS_MOUNT_PARAMS}\t1\t2" >> /etc/fstab
echo -e "LABEL=EFI\t/MidgardOS/boot\tvfat\t${EFI_MOUNT_PARAMS}\t1\t0" >> /etc/fstab
```

If, however, you want to use the mounting facilities in systemd, run the following commands to create the needed files in `/etc/systemd/system`:

```
cat << EOF > /etc/systemd/system/MidgardOS.rootfs.mount
[Unit]
Description=MidgardOS root filesystem volume
RequiresMountsFor=/

[Mount]
What=LABEL=MidgardOS
Where=/MidgardOS
Type=xfs
Options=rw,grpid,attr2,inode64,usrquota,prjquota
TimeoutSec=30

[Install]
WantedBy=local-fs.target
EOF

cat << EOF > /etc/systemd/system/MidgardOS-boot.mount
[Unit]
Description=MidgardOS EFI filesystem volume
RequiresMountsFor=/MidgardOS

[Mount]
What=LABEL=EFI
Where=/MidgardOS/boot
Type=vfat
Options=rw,fmask=0022,dmask=0022,codepage=437,iocharset=utf8,discard,showexec,sys_immutable,rodir,shortname=mixed
TimeoutSec=30

[Install]
WantedBy=local-fs.target
EOF

systemctl daemon-reload
systemctl enable MidgardOS.mount
systemctl enable MidgardOS-boot.mount
```

| Navigation |||
| --- | --- | ---: |
| [<<](./CreatePartition.md) Partition Creation | [HOME](./README.md) | Directories and Symlinks [>>](./DirectoriesAndSymlinks.md) |
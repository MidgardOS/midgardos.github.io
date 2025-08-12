| Navigation |||
| --- | --- | ---: |
| [<<](./README.md) | [HOME](./README.md) | [>>](./HostRequirements.md) |

# Pre-requisites for Building the MidgardOS Core Build Root

The build process for creating a new MidgardOS core operating system build root is not a simple task. It is assumed that the person building the build root is well versed in Linux systems administration and has a solid understanding of developing software for UNIX and unix-like systems.

As such, it is recommended to only start building a MidgardOS build root if you are comfortable at a Linux Bash or Zsh shell and are willing to work through discovering and fixing issues as the build progresses.

To avoid causing potential issues with your host computer, it is strongly recommended to build the OS inside a virtual machine, attaching extra disks, etc. as you go.

The current environment that is tested with building a MidgardOS core operating system build root is built using a VirtualBox VM running openSUSE Leap 15.6 or Fedora Linux 42 with the following virtual hardware configuration:

- 4 cores
- 8192 MiB RAM
- Video
  - 64MiB Video RAM
  - 1 screen
  - Controller: VMSVGA
  - Disable 3D Acceleration
- UEFI configuration, without Secure Boot
- 2 VMDK thin-provisioned disks on an NVMe controller
  - OS disk: 64GiB, two partitions
    - Partition 1: 512MiB FAT32 (EFI partition): Mounted at `/boot/efi`
    - Partition 2: Remaining disk, XFS (system root partition): Mounted at `/`
  - MidgardOS target disk: 64GiB, two partitions
    - Partition 1: 512MiB FAT32 (EFI partition): Mounted at `/MidgardOS/boot`
    - Partition 2: Remaining disk, XFS (target system root partition): Mounted at `/MidgardOS`
- Disable Audio Device
- USB 3.0 xHCI device
- Disable all Serial devices
- 1 Ethernet device
  - Bridged connection
    - Type: Intel Pro 1000 MT Server
    - Promiscuous Mode: Deny
    - Connected: true
- Disable Shared Folders

After creating the virtual machine, install the OS in 'Server only' text mode. After boot, set up the following configuration elements:

1. MidgardOS will use the unified CGroup v2 heirarchy, the host OS' system needs to support this inetead of the legacy hybrid hierarchy. To support this, add the `systemd.unified_cgroup_hierarchy=1` flag on the host's kernel command line if required.

| Navigation |||
| --- | --- | ---: |
| [<<](./README.md) | [HOME](./README.md) | [>>](./HostRequirements.md) |

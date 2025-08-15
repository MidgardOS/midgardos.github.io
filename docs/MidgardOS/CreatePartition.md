| Navigation |||
| --- | --- | ---: |
| [<<](./PrepOverview.md) | [HOME](./README.md) | [>>](./ManageDisk.md) |

# Section 1 - Preparation

## Rationale

While an isolated directory that is chrooted into can be used for systems that only do builds, to allow better testing and self-test validation of MidgardOS, we opt for using a normal partition to enable booting into the operating system later.

## Creating the Partitions

On the second disk of the virtual machine, if not done so at VM creation, add an additional 64 GiB disk (what will become the **nvme0n1** or **nvme0n2** disk device node depending on the order it was created in the VM) and boot up the VM. Now, as a privileged user, run `cfdisk /dev/$DEVICE_NODE`. It will ask to create a new partition table, with a few different types of partitioning schemes. Choose GPT for the partitioning type.

From here, create a 512 MiB EFI partition, and a 63.5 GiB Linux filesystem partition. Don't worry about formatting and mounting the disks, this will be covered in a later part of this documentation.

| Navigation |||
| --- | --- | ---: |
| [<<](./PrepOverview.md) | [HOME](./README.md) | [>>](./ManageDisk.md) |
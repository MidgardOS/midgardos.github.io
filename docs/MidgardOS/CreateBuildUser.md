| Navigation |||
| --- | --- | ---: |
| [<<](./FormatAndMount.md) | [HOME](./README.md) | [>>](./SetEnvironmentVars.md) |

# Creating the Build User

During the creation of the build environment, a non-privileged user other than root should be used for most tasks. When something calls for elevated privileges, the `sudo` command will be used to elevate rights to the superuser account. To simplify and standardize the build process for MidgardOS, run the following command to create the non-privileged `builder` user:

Note: if you are hosting the build from Fedora Linux, you might run into an issue where `useradd` will fail to properly create the user due to an issue with generating the `/etc/subuid` or `/etc/subgid` entry. To work around this, move the `/etc/subuid` and `/etc/subgid` files before creating the users, and put them back after creating the user.

```bash
groupadd -g 1001 -f build
useradd -u 1001 -s /bin/bash -g build -d /home/builder -c "MidgardOS build user" -M -N builder
install -v -d -m 755 -o builder -g build /home/builder
```

## Grant Full Permission Rights to the `tools` and `cross-tools` Trees

To allow the `builder` user to have full rights to the `sources`, `tools`, and `cross-tools` directory trees, run the following commands:

```bash
setfacl -m u:builder:rwx /MidgardOS/sources
setfacl -m u:builder:rwx /MidgardOS/tools
setfacl -m u:builder:rwx /MidgardOS/{etc,usr,usr/{bin,include,include/scsi,lib,lib64,sbin,share},var/lib}
```

## Add Local `sources` Symbolic Link in the `builder` User's Home Directory

To make downloading sources and patches easier during initial builds, it is recommended to create a symbolic link from `/MidgardOS/sources` in the `builder` user home directory, like so:

```bash
ln -sv /MidgardOS/sources ~builder/sources
```

| Navigation |||
| --- | --- | ---: |
| [<<](./FormatAndMount.md) | [HOME](./README.md) | [>>](./SetEnvironmentVars.md) |

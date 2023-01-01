| Navigation |||
| --- | --- | ---: |
| [<<](./CreateBuildUser.md) | [HOME](./README.md) | [>>](./IgnoringPreFinalSWTests.md) |

# Setting Environment Variables and Workspace

To allow builds to use the tools that are built using the steps in this guide and to ensure that errant host behaviours are mitigated as best as possible, there are a few things that need configured for the `builder` user. First, after logging in as root to the build VM, change to the `builder` user account, like so:

```bash
su - builder
```

## Create a New `.bash_profile`

To ensure that there are no extra environment variables that could cause issues with the builds, run the following command to create a new `.bash_profile` for the `builder` account:

```bash
cat > ~/.bash_profile << "EOF"
exec env -i HOME=${HOME} TERM=${TERM} PS1='\u:\w\$ ' /bin/bash
EOF
```

## Create a New `.bashrc`

The new `bash` instance that is exec'd in from the `.bash_profile` is a non-login shell, so an appropriate `.bashrc` needs put into place:

```bash
cat > ~/.bashrc << "EOF"
set +h
umask 022
BRFS=/MidgardOS
LC_ALL=POSIX
PATH=/cross-tools/bin:/bin:/usr/bin
export BRFS LC_ALL PATH
unset CFLAGS CXXFLAGS PKG_CONFIG_PATH
EOF
```

## Build Environment Variables

While building the cross compilation and basic tools, there are a few extra environment variables that will need set to ease running configuration and build commands. The first of these is the target build host triplet. This is used to tell the compiler to build the binaries for a specific architecture. As the target host and the initial build host have potentially the same value, which will block the compiler from building cross-compiled target tools, the target build triplet needs modifed by forcing it to be a "cross" target. To do so, set the `BRFS_HOST` environment variable:

```bash
export BRFS_HOST=$(echo ${MACHTYPE} | sed -e 's/-[^-]*/-cross/')
```

Next, the triplet for the target architecture needs set, much like for the target host:

```bash
export BRFS_TARGET="x86_64-unknown-linux-gnu"
```

Now, to support building the 32-bit libraries, set the same for the 32-bit architecture:

```bash
export BRFS_TARGET32="i686-pc-linux-gnu"
```

Finally, the environment variables for the target-specific linker flags needs set:

```bash
export BUILD32="-m32"
export BUILD64="-m64"
```

To make these persist during the build of the cross-tools and basic tools trees, run the following command:

```bash
cat >> ~/.bashrc << EOF
export BRFS_HOST="${BRFS_HOST}"
export BRFS_TARGET="${BRFS_TARGET}"
export BRFS_TARGET32="${BRFS_TARGET32}"
export BUILD32="${BUILD32}"
export BUILD64="${BUILD64}"
EOF
```

Now that the `.bashrc` and `.bash_profile` have been created, source the `.bash_profile` to validate that the new environment is configured correctly:

```bash
source ~/.bash_profile
```

| Navigation |||
| --- | --- | ---: |
| [<<](./CreateBuildUser.md) | [HOME](./README.md) | [>>](./IgnoringPreFinalSWTests.md) |

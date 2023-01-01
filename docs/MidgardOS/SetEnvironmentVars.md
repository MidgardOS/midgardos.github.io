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

Now that the `.bashrc` and `.bash_profile` have been created, source the `.bash_profile` to validate that the new environment is configured correctly:

```bash
source ~/.bash_profile
```

| Navigation |||
| --- | --- | ---: |
| [<<](./CreateBuildUser.md) | [HOME](./README.md) | [>>](./IgnoringPreFinalSWTests.md) |

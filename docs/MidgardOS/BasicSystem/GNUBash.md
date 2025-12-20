# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGrep.md) GNU Grep | [HOME](../README.md) | GNU GDBM 64-bit [>>](./GNUGDBM_64bit.md) |

## GNU Bash

Name: GNU Bash<br />
Summary: The GNU Bourne-Again Shell<br />
License: GPL v3<br />
Version: 5.3<br />
URL: [https://ftp.gnu.org/gnu/bash/bash-5.3.tar.gz](https://ftp.gnu.org/gnu/bash/bash-5.3.tar.gz)<br />

Average Build Time: 1.5 SBU<br />
Used Install Space: 56 MiB<br />

## Configuration

To configure GNU Bash for install into the build root, run the following command:

```bash
./configure --prefix=/usr                                           \
            --libdir=/usr/lib64                                     \
            --libexecdir=/usr/lib64                                 \
            --with-curses                                           \
            --with-gnu-ld                                           \
            --without-gnu-malloc                                    \
	        --without-bash-malloc                                   \
	        --enable-mem-scramble                                   \
            --enable-threads=posix		                            \
	        --enable-job-control		                            \
	        --enable-net-redirections	                            \
        	--enable-alias			                                \
	        --enable-readline		                                \
	        --enable-history		                                \
	        --enable-bang-history		                            \
	        --enable-directory-stack	                            \
	        --enable-process-substitution	                        \
	        --enable-prompt-string-decoding	                        \
	        --enable-select			                                \
	        --enable-help-builtin		                            \
	        --enable-separate-helpfiles	                            \
	        --enable-array-variables	                            \
            --enable-alt-array-implementation                       \
            --enable-brace-expansion	                            \
    	    --enable-command-timing		                            \
	        --enable-disabled-builtins	                            \
    	    --enable-glob-asciiranges-default                       \
        	--enable-translatable-strings	                        \
	        --disable-strict-posix-default	                        \
    	    --enable-multibyte		                                \
    	    --enable-separate-helpfiles=/usr/share/bash/helpfiles   \
            --with-installed-readline                               \
            --docdir=/usr/share/doc/bash-5.3
```

## Compilation and Installation

To compile GNU Bash, run the following command:

```bash
make
```

The tests are not designed to work correctly running uner the root user. The tests will be run later when it is rebuilt as an RPM.

Finally, to install GNU Bash into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
pushd DESTDIR
    ln -sv usr/bin/bash usr/bin/rbash
    ln -sv usr/bin/bash usr/bin/sh
    install -v -d -m755 -o root -g root etc
    install -v -d -m755 -o root -g root etc/skel
    install -v -d -m755 -o root -g root etc/skel/.bashrc.d
    install -v -d -m755 -o root -g root etc/profile.d
popd
cat <<EOF > DESTDIR/etc/skel/.bashrc
# .bashrc
# Source global definitions
if [ -f /etc/bashrc ]; then
    . /etc/bashrc
fi

# User specific environment
if ! [[ "$PATH" =~ "$HOME/.local/bin:$HOME/bin:" ]]; then
    PATH="$HOME/.local/bin:$HOME/bin:$PATH"
fi
export PATH

# Comment the following line if you like systemctl's auto-paging feature:
export SYSTEMD_PAGER=

# User specific aliases and functions
if [[ -d ~/.bashrc.d ]]; then
    for rc in ~/.bashrc.d/*; do
        if [[ -f "$rc" ]]; then
            . "$rc"
        fi
    done
fi
unset rc

# check for other alias or environment files
if [[ -f $HOME/.aliases ]]; then
    . $HOME/.bash.aliases
elif [[ -f $HOME/.bash.environ ]]; then
    . $HOME/.bash.environ
fi
EOF
cat <<EOF > DESTDIR/etc/skel/.bash.aliases
# Place your bash aliases here
EOF
cat <<EOF > DESTDIR/etc/skel/.bash.environ
# Place your personal environment variables here
EOF
cat <<EOF > DESTDIR/etc/skel/.bash_logout
# Cleanup functions and tasks that should be run on logout should be added here
EOF
cat <<EOF > DESTDIR/etc/skel/.bash_profile
# .bash_profile
# Get the aliases and functions
if [ -f ~/.bashrc ]; then
        . ~/.bashrc
fi

# User specific environment and startup programs
EOF
install -v -m 644 -o root -g root /sources/system_files/etc/bashrc DESTDIR/etc/bashrc
install -v -m 644 -o root -g root /sources/system_files/etc/profile DESTDIR/etc/profile
cp -Rv DESTDIR/etc/* /etc/
cp -Rv DESTDIR/usr/* /usr/
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | bash, bashbug, rbash, sh |
| Installed Files | /etc/bashrc, /etc/profile, /etc/skel/.bashrc, /etc/skel/.bash.aliases, /etc/skel/.bash.environ, /etc/skel/.bash_logout, /etc/skel/.bash_profile |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGrep.md) GNU Grep | [HOME](../README.md) | GNU GDBM 64-bit [>>](./GNUGDBM_64bit.md) |

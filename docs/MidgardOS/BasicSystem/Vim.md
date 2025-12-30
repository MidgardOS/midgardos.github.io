# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTexinfo.md) GNU Texinfo | [HOME](../README.md) | GNU Nano [>>](./GNUNano.md) |

## Vi Improved

Name: Vi Improved<br />
Summary: An advanced powerful terminal text editor<br />
License: $PKG_LICENSE<br />
Version: 9.1.2031<br />
URL: []($PKG_URL)<br />

Average Build Time: 3.7 SBU<br />
Used Install Space: 48 MiB<br />

## Notes

Vi Improved has a number of optional dependencies that are not available in MidgardOS at this time. These include:

- Lua
- X Input Manager
- Ruby
- Java
- X11
- Gtk2 and Gtk3
- LibCanberra
- LibSodium
- GPM
- Wayland
- Motif

## Configuration

To configure Vi Improved for install into the build root, run the following command:

```bash
echo '#define SYS_VIMRC_FILE "/etc/vimrc"' >> src/feature.h
./configure --prefix=/usr                   \
            --libdir=/usr/lib64             \
            --libexecdir=/usr/lib64         \
            --docdir=/usr/share/doc/vim-9.1 \
            --enable-year2038
```

## Compilation and Installation

To compile Vi Improved, run the following command:

```bash
make
```

Next, run the test suite:

```bash
sed '/test_plugin_glvs/d' -i src/testdir/Make_all.mak
useradd -c "Test User" -u 1000 -U -m tester
unalias cp
cp -R ../vim-9.1.2031 /tmp/
alias cp="cp -i"
chown -R tester /tmp/vim-9.1.2031/
cd /tmp/vim-9.1.2031
su tester -c "PATH=$PATH autoreconf -fvi && TERM=xterm-256color LANG=en_US.UTF-8 make -j1 test" &> vim-test.log
cd -
rm -rf /tmp/vim-9.1.2031
userdel -rf tester
```

One test out of the more than 7300 tests is failing due to a potential syntax issue in the test. For now, it is safe to proceed.

Finally, to install Vi Improved into the build tree, run the following command:

```bash
make install
ln -sv vim /usr/bin/vi
for L in  /usr/share/man/{,*/}man1/vim.1; do
    ln -sv vim.1 $(dirname $L)/vi.1
done
ln -sv ../vim/vim91/doc /usr/share/doc/vim-9.1
```

## Configuration

Run the following commands to initially configure Vi Improved for use:

```bash
cat > /etc/vimrc << "EOF"
" Begin /etc/vimrc

" Ensure defaults are set before customizing settings, not after
source $VIMRUNTIME/defaults.vim
let skip_defaults_vim=1

set nocompatible
set backspace=2
set mouse=
syntax on
if (&term == "xterm") || (&term == "putty")
  set background=dark
endif

" End /etc/vimrc
EOF
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | ex, rview, rvim, vi, view, vim, vimdiff, vimtutor, and xxd |
| Installed Files | /etc/vimrc, /usr/share/vim |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTexinfo.md) GNU Texinfo | [HOME](../README.md) | GNU Nano [>>](./GNUNano.md) |

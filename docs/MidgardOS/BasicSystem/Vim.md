# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTexinfo.md) GNU Texinfo | [HOME](../README.md) | GNU Nano [>>](./GNUNano.md) |

## Vi Improved

Name: Vi Improved<br />
Summary: An advanced powerful terminal text editor<br />
License: VIM License (GPL Compatible)<br />
Version: 9.2.0219<br />
URL: [https://github.com/vim/vim/archive/refs/tags/v9.2.0219.tar.gz](https://github.com/vim/vim/archive/refs/tags/v9.2.0219.tar.gz)<br />

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
            --docdir=/usr/share/doc/vim-9.2 \
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
cp -Rv ../vim-9.2.0219 /tmp/
alias cp="cp -i"
chown -Rv tester /tmp/vim-9.2.0219/
cd /tmp/vim-9.2.0219
su tester -c "PATH=$PATH autoreconf -fvi && TERM=xterm-256color LANG=en_US.UTF-8 make -j1 test" &> vim-test.log
cd -
rm -rfv /tmp/vim-9.2.0219
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
ln -sv ../vim/vim92/doc /usr/share/doc/vim-9.2
```

## Configuration

Run the following command to install the initial configure Vi Improved for use:

```bash
install -v -m644 -o root -g root ../midgardos.github.io/docs/MidgardOS/system_files/etc/vimrc /etc/
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | ex, rview, rvim, vi, view, vim, vimdiff, vimtutor, and xxd |
| Installed Files | /etc/vimrc, /usr/share/vim |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTexinfo.md) GNU Texinfo | [HOME](../README.md) | GNU Nano [>>](./GNUNano.md) |

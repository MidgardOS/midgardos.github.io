# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection | [HOME](../README.md) | NCurses 32-bit [>>](./NCurses_32bit.md) |

## NCurses 64-bit

Name: NCurses 64-bit<br />
Summary: Terminal-independent library for console interfaces<br />
License: MIT<br />
Version: 6.5-20250823<br />
URL: [https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250823.tgz](https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250823.tgz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 46 MiB<br />

## Configuration

To configure NCurses 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr            \
            --libdir=/usr/lib64      \
            --libexecdir=/usr/lib64  \
            --mandir=/usr/share/man  \
            --with-shared            \
            --without-debug          \
            --without-normal         \
            --with-cxx-shared        \
            --enable-pc-files        \
            --with-pkg-config-libdir=/usr/lib64/pkgconfig
```

## Compilation and Installation

To compile NCurses 64-bit, run the following command:

```bash
make
```

The test suite for NCurses cannot be run before it is installed. It will be tested later when being rebuilt as an RPM.

Finally, to install NCurses 64-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/dest install
sed -e 's/^#if.*XOPEN.*$/#if 1/' -i dest/usr/include/curses.h
cp --remove-destination -afv dest/* /
for lib in ncurses form panel menu ; do
    ln -sfv lib${lib}w.so /usr/lib64/lib${lib}.so
    ln -sfv ${lib}w.pc    /usr/lib64/pkgconfig/${lib}.pc
done
ln -sfv libncursesw.so /usr/lib64/libcurses.so
cp -v -R doc -T /usr/share/doc/ncurses-6.5-20250809
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | captoinfo (link to tic), clear, infocmp, infotocap (link to tic), ncursesw6-config, reset (link to tset), tabs, tic, toe, tput, and tset |
| Installed Libraries | libcurses.so (symlink), libform.so (symlink), libformw.so, libmenu.so (symlink), libmenuw.so, libncurses.so (symlink), libncursesw.so, libncurses++w.so, libpanel.so (symlink), and libpanelw.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection | [HOME](../README.md) | NCurses 32-bit [>>](./NCurses_32bit.md) |

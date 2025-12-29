# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./IPRoute2.md) IPRoute2 | [HOME](../README.md) | LibPipeline 64-bit [>>](./LibPipeline_64bit.md) |

## Kbd

Name: Kbd<br />
Summary: Utilities, key tables, and console fonts<br />
License: $PKG_LICENSE<br />
Version: 2.9.0<br />
URL: [https://github.com/legionus/kbd/releases/download/v2.9.0/kbd-2.9.0.tar.xz](https://github.com/legionus/kbd/releases/download/v2.9.0/kbd-2.9.0.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 5.3 MiB<br />

## Notes

This package's tests need Valgrind. Once Valgrind is installed, rebuild this package.

## Preparation

To correct handling of the keyboard's DELETE and BACKSPACE keys in keymaps, apply the following patch:

```bash
patch -Np1 -i ../patches/kbd-2.9.0-backspace-1.patch
```

## Configuration

To configure Kbd for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/kbd-2.9.0   \
            --sysconfdir=/etc                   \
            --enable-optional-progs             \
            --enable-libkeymap                  \
            --enable-libkfont                   \
            --enable-nls                        \
            --disable-static
```

## Compilation and Installation

To compile Kbd, run the following commands:

```bash
make
cd ../src/kbd
make
cd -
```

This package's test suite fails with a broken final test for libkbdfile. This is safe to skip for now.

Finally, to install Kbd into the build tree, run the following command:

```bash
make install
cd ../src/kbd
make install
cd -
install -v -d -m 755 -o root -g root /etc/pam.d
install -v -m 644 -o root -g root ../system_files/etc/pam.d/vlock /etc/pam.d/
install -v -d -m 755 -o root -g root /usr/sbin
install -v -m 755 -o root -g root ../system_files/usr/sbin/kbdsettings /usr/sbin/
install -v -d -m 755 -o root -g root /usr/lib/systemd/system
install -v -m 644 -o root -g root ../system_files/usr/lib/systemd/system/kbdsettings.service /usr/lib/systemd/system/
rm -fv /usr/lib64/libkbdfile.la
rm -fv /usr/lib64/libkeymap.la
rm -fv /usr/lib64/libkfont.la
unalias cp
cp -R -v docs/doc -T /usr/share/doc/kbd-2.9.0
alias cp="cp -i"
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | chvt, clrunimap, deallocvt, dumpkeys, fgconsole, getkeycodes, getunimap, kbd_mode, kbdinfo, kbdrate, loadkeys, loadunimap, mapscrn, openvt, outpsfheader, psfaddtable, psfgettable, psfstriptable, psfxtable, resizecons, screendump, setfont, setkeycodes, setleds, setlogcons, setmetamode, setpalette, setvesablank, setvtrgb, showconsolefont, showkey, spawn_console, spawn_login, unicode_start, unicode_stop, and vlock |
| Installed Libraries | libkbdfile.so, libkeymap.so, and libkfont.so |
| Installed Files | /usr/share/consolefonts, /usr/share/consoletrans, /usr/share/keymaps, and /usr/share/unimaps |

| Navigation |||
| --- | --- | ---: |
| [<<](./IPRoute2.md) IPRoute2 | [HOME](../README.md) | LibPipeline 64-bit [>>](./LibPipeline_64bit.md) |

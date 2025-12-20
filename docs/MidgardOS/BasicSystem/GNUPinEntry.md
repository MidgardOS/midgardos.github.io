# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPrivacyGuard.md) GNU Privacy Guard | [HOME](../README.md) | LibXML2 64-bit [>>](./LibXML2_64bit.md) |

## GNU PinEntry

Name: GNU PinEntry<br />
Summary: Simple PIN or passphrase entry dialogs<br />
License: GPL v2 or later<br />
Version: 1.3.2<br />
URL: [https://www.gnupg.org/ftp/gcrypt/pinentry/pinentry-1.3.2.tar.bz2](https://www.gnupg.org/ftp/gcrypt/pinentry/pinentry-1.3.2.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 512 KiB<br />

## Configuration

To configure GNU PinEntry for install into the build root, run the following command:

```bash
./configure --prefix=/usr               \
            --libdir=/usr/lib64         \
            --libexecdir=/usr/lib64     \
            --enable-pinentry-curses    \
	        --enable-pinentry-tty       \
	        --disable-libsecret         \
	        --disable-pinentry-qt       \
	        --disable-pinentry-gtk2     \
	        --disable-pinentry-gnome3   \
	        --disable-pinentry-emacs    \
	        --disable-inside-emacs      \
	        --disable-pinentry-fltk     \
	        --disable-pinentry-efl
```

NOTE: This version of the GNU PinEntry tool is configured only for CLI dialogs via NCurses or the raw TTY interfaces. No other toolkits are configured for use at this time.

## Compilation and Installation

To compile GNU PinEntry, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU PinEntry into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | pinentry, pinentry-curses, and pinentry-tty |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPrivacyGuard.md) GNU Privacy Guard | [HOME](../README.md) | LibXML2 64-bit [>>](./LibXML2_64bit.md) |

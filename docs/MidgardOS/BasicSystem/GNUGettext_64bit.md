# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./PSMisc.md) PSMisc | [HOME](../README.md) | GNU Gettext 32-bit [>>](./GNUGettext_32bit.md) |

## GNU Gettext 64-bit

Name: GNU Gettext 64-bit<br />
Summary: Utilities for i18n and l10n tasks<br />
License: GPL v3<br />
Version: 0.26<br />
URL: [https://ftp.gnu.org/gnu/gettext/gettext-0.26.tar.xz](https://ftp.gnu.org/gnu/gettext/gettext-0.26.tar.xz)<br />

Average Build Time: 2.6 SBU<br />
Used Install Space: 395 MiB<br />

## Configuration

To configure GNU Gettext 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --docdir=/usr/share/doc/gettext-0.26
```

## Compilation and Installation

To compile GNU Gettext 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Gettext 64-bit into the build tree, run the following commands:

```bash
make install
chmod -v 0755 /usr/lib64/preloadable_libintl.so
rm -fv /usr/lib64/libasprintf.la
rm -fv /usr/lib64/libgettext*.la
rm -fv /usr/lib64/libtextstyle.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | autopoint, envsubst, gettext, gettext.sh, gettextize, msgattrib, msgcat, msgcmp, msgcomm, msgconv, msgen, msgexec, msgfilter, msgfmt, msggrep, msginit, msgmerge, msgunfmt, msguniq, ngettext, recode-sr-latin, and xgettext |
| Installed Libraries | libasprintf.so, libgettextlib.so, libgettextpo.so, libgettextsrc.so, libtextstyle.so, and preloadable_libintl.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./PSMisc.md) PSMisc | [HOME](../README.md) | GNU Gettext 32-bit [>>](./GNUGettext_32bit.md) |

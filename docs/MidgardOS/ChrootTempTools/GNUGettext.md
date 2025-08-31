# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./EssentialFilesAndSymlinks.md) Creating Essential Files and Symlinks | [HOME](../README.md) | GNU Bison [>>](./GNUBison.md) |

## GNU Gettext

Name: GNU Gettext<br />
Summary: Utilities for supporting i18n and l10n for applications<br />
License: GPL v3 or later<br />
Version: 0.26<br />
URL: [https://ftp.gnu.org/gnu/gettext/gettext-0.26.tar.xz](https://ftp.gnu.org/gnu/gettext/gettext-0.26.tar.xz)<br />

## Configuration

To configure GNU Gettext for install into the build root, run the following command:

```bash
./configure --prefix=/usr --libdir=/usr/lib64 --libexecdir=/usr/lib64 --disable-shared
```

For now, we only need a couple applications from the utilities that are normally installed.

## Compilation and Installation

To compile GNU Gettext, run the following command:

```bash
make
```

Finally, to install the minimal required tools from GNU Gettext into the build tree, run the following command:

```bash
cp -v gettext-tools/src/{msgfmt,msgmerge,xgettext} /usr/bin
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./EssentialFilesAndSymlinks.md) Creating Essential Files and Symlinks | [HOME](../README.md) | GNU Bison [>>](./GNUBison.md) |

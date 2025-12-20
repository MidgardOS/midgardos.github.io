# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUSed.md) GNU Sed | [HOME](../README.md) | GNU Gettext 64-bit [>>](./GNUGettext_64bit.md) |

## PSMisc

Name: PSMisc<br />
Summary: Tools for retrieving information about running processes<br />
License: GPL v2<br />
Version: 23.7<br />
URL: [https://sourceforge.net/projects/psmisc/files/psmisc/psmisc-23.7.tar.xz/download](https://sourceforge.net/projects/psmisc/files/psmisc/psmisc-23.7.tar.xz/download)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 6.7 MiB<br />

## Configuration

To configure PSMisc for install into the build root, run the following command:

```bash
./configure --prefix=/usr
```

## Compilation and Installation

To compile PSMisc, run the following command:

```bash
make
```

Now, run the test suite:

```bash
make check
```

Finally, to install PSMisc into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | fuser, killall, peekfd, prtstat, pslog, pstree, and pstree.x11 (link to pstree) |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUSed.md) GNU Sed | [HOME](../README.md) | GNU Gettext 64-bit [>>](./GNUGettext_64bit.md) |

# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUInetUtils.md) GNU InetUtils | [HOME](../README.md) | Perl [>>](./Perl.md) |

## GNU Less

Name: GNU Less<br />
Summary: A text viewer and pager<br />
License: GPL v3 or later<br />
Version: 685<br />
URL: [https://ftp.gnu.org/gnu/less/less-685.tar.gz](https://ftp.gnu.org/gnu/less/less-685.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.1 MiB<br />

## Configuration

To configure GNU Less for install into the build root, run the following command:

```bash
./configure --prefix=/usr --sysconfdir=/etc
```

## Compilation and Installation

To compile GNU Less, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Note, there are a couple tests for UTF8 support that fail with this release.

Finally, to install GNU Less into the build tree, run the following command:

```bash
make install
install -v -m755 -o root -g root ../system_files/usr/bin/lessopen.sh /usr/bin/lessopen.sh
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | less lessecho lesskey |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUInetUtils.md) GNU InetUtils | [HOME](../README.md) | Perl [>>](./Perl.md) |

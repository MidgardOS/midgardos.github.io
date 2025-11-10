# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./meson.md) Meson | [HOME](../README.md) | LibSEPol 64-bit [>>](./libsepol_64bit.md) |

## IP Utils

Name: IP Utils<br />
Summary: IPv4 and v6 networking utilities<br />
License: BSD 3 clause and GPL v2 or later<br />
Version: 20250605<br />
URL: [https://github.com/iputils/iputils/releases/download/20250605/iputils-20250605.tar.xz](https://github.com/iputils/iputils/releases/download/20250605/iputils-20250605.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: approximately 160 KiB<br />

## Configuration

To configure IP Utils for install into the build root, run the following command:

```bash
meson setup --prefix=/usr                               \
            --buildtype=release                         \
            -Dsystemdunitdir=/usr/lib/systemd/system    \
            -DBUILD_PING=false                          \
            -DINSTALL_SYSTEMD_UNITS=true                \
            -DSETCAP_OR_SUID_ARPING=true                \
            -DSETCAP_OR_SUID_CLOCKDIFF=true build .

```

## Compilation and Installation

To compile IP Utils, run the following commands:

```bash
cd build
export NINJAJOBS=4
ninja
```

Next, run the test suite:

```bash
meson test
```

Finally, to install IP Utils into the build tree, run the following commands:

```bash
meson install
mv -v /usr/share/iputils /usr/share/doc/
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | arping, clockdiff, and tracepath |

| Navigation |||
| --- | --- | ---: |
| [<<](./meson.md) Meson | [HOME](../README.md) | LibSEPol 64-bit [>>](./libsepol_64bit.md) |

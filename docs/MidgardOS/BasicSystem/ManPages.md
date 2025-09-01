# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Overview.md) Overview | [HOME](../README.md) | IANA Etc Files [>>](./IanaEtcFiles.md) |

## Man Pages

Name: Man Pages<br />
Summary: Documentation for system C APIs, important device node files, and significant configuration files<br />
License: BSD-2-Clause/GPLv1-through-v3/MIT/Custom<br />
Version: 6.15<br />
URL: [https://www.kernel.org/pub/linux/docs/man-pages/man-pages-6.15.tar.xz](https://www.kernel.org/pub/linux/docs/man-pages/man-pages-6.15.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 52 MiB<br />

## Preparation

The `man3/crypt*` man pages are going to be provided by the `libxcrypt` package, so remove them:

```bash
rm -f man3/crypt*
```

## Installation

To install Man Pages in the system, run the following command:

```bash
make -R GIT=false prefix=/usr install
```

## Contents

| Installed Files | Description |
| --- | --- |
| Man Pages | approximately 2400 man pages describing the C language functions, device node files, and configuration files |

| Navigation |||
| --- | --- | ---: |
| [<<](./Overview.md) Overview | [HOME](../README.md) | IANA Etc Files [>>](./IanaEtcFiles.md) |

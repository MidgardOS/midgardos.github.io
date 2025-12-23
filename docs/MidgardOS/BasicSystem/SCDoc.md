# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Linux-PAM_32bit.md) Linux PAM 32-bit | [HOME](../README.md) | Kmod 64-bit [>>](./Kmod_64bit.md) |

## SCDoc

Name: SCDoc<br />
Summary: A simple man-page generator written in C<br />
License: MIT<br />
Version: 1.11.4<br />
URL: [https://git.sr.ht/~sircmpwn/scdoc/archive/1.11.4.tar.gz](https://git.sr.ht/~sircmpwn/scdoc/archive/1.11.4.tar.gz)<br />

Average Build Time: SBU<br />
Used Install Space: <br />

## Configuration

This package does not come with a traditional configuration script.

## Compilation and Installation

To compile SCDoc, run the following command:

```bash
make PREFIX=/usr
```

Next, run the test suite:

```bash
make check
```

Finally, to install SCDoc into the build tree, run the following command:

```bash
make PREFIX=/usr install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | scdoc |

| Navigation |||
| --- | --- | ---: |
| [<<](./Linux-PAM_32bit.md) Linux PAM 32-bit | [HOME](../README.md) | Kmod 64-bit [>>](./Kmod_64bit.md) |

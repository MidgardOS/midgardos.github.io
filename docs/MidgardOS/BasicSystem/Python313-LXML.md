# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSS_32bit-pass1.md) Mozilla NSS 32-bit - pass 1 | [HOME](../README.md) | ITSTool [>>](./ITSTool.md) |

## Python module LXML

Name: Python module LXML<br />
Summary: Python bindings for libxslt and libxml2<br />
License: BSD 3-clause<br />
Version: 6.0.2<br />
URL: [https://github.com/lxml/lxml/releases/download/lxml-6.0.2/lxml-6.0.2.tar.gz](https://github.com/lxml/lxml/releases/download/lxml-6.0.2/lxml-6.0.2.tar.gz)<br />

Average Build Time: 0.8 SBU<br />
Used Install Space: 104 MiB<br />

## Compilation and Installation

To compile Python module LXML, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module LXML into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user lxml
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | lxml |
| Installed Files | /usr/lib/python3.13/site-packages/lxml and /usr/lib/python3.13/site-packages/lxml-6.0.2.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./MozillaNSS_32bit-pass1.md) Mozilla NSS 32-bit - pass 1 | [HOME](../README.md) | ITSTool [>>](./ITSTool.md) |

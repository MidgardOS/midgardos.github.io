# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNano.md) GNU Nano | [HOME](../README.md) | Python module Jinja2 [>>](./Python313-Jinja2.md) |

## Python module MarkupSafe

Name: Python module MarkupSafe<br />
Summary: A Python module to implement an XML/HTML/XHTML Markup safe string<br />
License: BSD 3-clause<br />
Version: 3.0.3<br />
URL: [https://pypi.org/packages/source/M/MarkupSafe/markupsafe-3.0.3.tar.gz](https://pypi.org/packages/source/M/MarkupSafe/markupsafe-3.0.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 500 KiB<br />

## Compilation and Installation

To compile Python module MarkupSafe, run the following command:

```bash
pip3 wheel -w dist --no-cache-dir --no-build-isolation --no-deps $PWD
```

Finally, to install Python module MarkupSafe into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist Markupsafe
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | MarkupSafe |
| Installed Files | /usr/lib64/python3.13/site-packages/markupsafe, /usr/lib64/python3.13/site-packages/markupsafe-3.0.3 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNano.md) GNU Nano | [HOME](../README.md) | Python module Jinja2 [>>](./Python313-Jinja2.md) |

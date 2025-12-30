# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-MarkupSafe.md) Python module MarkupSafe | [HOME](../README.md) | Cython [>>](./Python313-Cython.md) |

## Python module Jinja2

Name: Python module Jinja2<br />
Summary: A Python module to implements a simple template language<br />
License: BSD 3-clause<br />
Version: 3.1.6<br />
URL: [https://github.com/pallets/jinja/releases/download/3.1.6/jinja2-3.1.6.tar.gz](https://github.com/pallets/jinja/releases/download/3.1.6/jinja2-3.1.6.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.6 MiB<br />

## Compilation and Installation

To compile Python module Jinja2, run the following command:

```bash
pip3 wheel -w dist --no-cache-dir --no-build-isolation --no-deps $PWD
```

Finally, to install Python module Jinja2 into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist Jinja2
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | Jinja2 |
| Installed Files | /usr/lib/python3.13/site-packages/jinja, /usr/lib/python3.13/jinja2-3.1.6.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-MarkupSafe.md) Python module MarkupSafe | [HOME](../README.md) | Cython [>>](./Python313-Cython.md) |

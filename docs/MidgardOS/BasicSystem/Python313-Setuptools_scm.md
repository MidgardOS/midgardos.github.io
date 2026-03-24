# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Pathspec.md) Python module Pathspec | [HOME](../README.md) | Python module Pluggy [>>](./$NEXT_PAGE.md) |

## Python module Setuptools_scm

Name: Python module Setuptools_scm<br />
Summary: A Python library for extracting module versions from Git or Hg metadata<br />
License: MIT<br />
Version: 9.2.2<br />
URL: [https://github.com/pypa/setuptools-scm/archive/refs/tags/v9.2.2.tar.gz](https://github.com/pypa/setuptools-scm/archive/refs/tags/v9.2.2.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.6 MiB<br />

## Compilation and Installation

To compile Python module Setuptools_scm, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Setuptools_scm into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user setuptools_scm
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | Setuptools_scm |
| Installed Files | /usr/lib/python3.14/site-packages/setuptools_scm and /usr/lib/python3.14/site-packages/setuptools_scm-9.2.2.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Pathspec.md) Python module Pathspec | [HOME](../README.md) | Python module Pluggy [>>](./$NEXT_PAGE.md) |

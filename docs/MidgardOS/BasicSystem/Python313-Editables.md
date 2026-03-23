# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ITSTool.md) ITS Tool | [HOME](../README.md) | Python module Pathspec [>>](./Python313-Pathspec.md) |

## Python module Editables

Name: Python module Editables<br />
Summary: A Python library for creating "editable" wheels<br />
License: MIT<br />
Version: 0.5<br />
URL: [https://github.com/pfmoore/editables/archive/refs/tags/0.5.tar.gz](https://github.com/pfmoore/editables/archive/refs/tags/0.5.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 180 KiB<br />

## Compilation and Installation

To compile Python module Editables, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Editables into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user editables
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | Editables |
| Installed Files | /usr/lib/python3.14/site-packages/editables and /usr/lib/python3.14/site-packages/editables-0.5.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./ITSTool.md) ITS Tool | [HOME](../README.md) | Python module Pathspec [>>](./Python313-Pathspec.md) |

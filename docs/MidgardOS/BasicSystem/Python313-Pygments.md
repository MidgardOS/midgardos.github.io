# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Iniconfig.md) Python module Iniconfig | [HOME](../README.md) | Python module Pytest [>>](./Python313-Pytest.md) |

## Python module Pygments

Name: Python module Pygments<br />
Summary: A general syntax highlighter written in Python that supports more than 300 languages<br />
License: BSD 2-clause<br />
Version: 2.19.2<br />
URL: [https://github.com/pygments/pygments/archive/refs/tags/2.19.2.tar.gz](https://github.com/pygments/pygments/archive/refs/tags/2.19.2.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 54 MiB<br />

## Compilation and Installation

To compile Python module Pygments, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Pygments into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user Pygments
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | pygmentize |
| Installed Plugins | pygments |
| Installed Files | /usr/lib/python3.14/site-packages/pygments and /usr/lib/python3.14/site-packages/Pygments-2.19.2.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Iniconfig.md) Python module Iniconfig | [HOME](../README.md) | Python module Pytest [>>](./Python313-Pytest.md) |

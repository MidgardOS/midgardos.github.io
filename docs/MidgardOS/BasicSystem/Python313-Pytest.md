# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Iniconfig.md) Python module Iniconfig | [HOME](../README.md) | Python module Pygments [>>](./Python313-Pygments.md) |

## Python module Pytest

Name: Python module Pytest<br />
Summary: A testing framework for Python modules<br />
License: MIT<br />
Version: 9.0.2<br />
URL: [https://github.com/pytest-dev/pytest/releases/download/9.0.2/pytest-9.0.2.tar.gz](https://github.com/pytest-dev/pytest/releases/download/9.0.2/pytest-9.0.2.tar.gz)<br />

Average Build Time: 1.5 SBU<br />
Used Install Space: 45 MiB<br />

## Compilation and Installation

To compile Python module Pytest, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Pytest into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user pytest
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | pytest and py.test |
| Installed Plugins | pytest |
| Installed Files | /usr/lib/python3.14/site-packages/_pytest, /usr/lib/python3.14/site-packages/pytest, and /usr/lib/python3.14/site-packages/pytest-9.0.2.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Iniconfig.md) Python module Iniconfig | [HOME](../README.md) | Python module Pygments [>>](./Python313-Pygments.md) |

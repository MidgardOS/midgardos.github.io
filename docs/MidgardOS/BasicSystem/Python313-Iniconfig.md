# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Hatch-vcs.md) Python module Hatch-vcs | [HOME](../README.md) | Python module Pygments [>>](./Python313-Pygments.md) |

## Python module Iniconfig

Name: Python module Iniconfig<br />
Summary: A small and simple INI parser written in Python<br />
License: MIT<br />
Version: 2.3.0<br />
URL: [https://github.com/pytest-dev/iniconfig/releases/download/v2.3.0/iniconfig-2.3.0.tar.gz](https://github.com/pytest-dev/iniconfig/releases/download/v2.3.0/iniconfig-2.3.0.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 304 KiB<br />

## Compilation and Installation

To compile Python module Iniconfig, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Iniconfig into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user iniconfig
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | iniconfig |
| Installed Files | /usr/lib/python3.14/site-packages/iniconfig and /usr/lib/python3.14/site-packages/iniconfig-2.3.0.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Hatch-vcs.md) Python module Hatch-vcs | [HOME](../README.md) | Python module Pygments [>>](./Python313-Pygments.md) |

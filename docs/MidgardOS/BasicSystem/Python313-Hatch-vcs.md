# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Hatchling.md) Python module Hatchling | [HOME](../README.md) | Python module Iniconfig [>>](./Python313-Iniconfig.md) |

## Python module Hatch-vcs

Name: Python module Hatch-vcs<br />
Summary: A Hatch plugin for handling versioning based on a variety of VCS tools<br />
License: MIT<br />
Version: 0.5.0<br />
URL: [https://github.com/ofek/hatch-vcs/archive/refs/tags/v0.5.0.tar.gz](https://github.com/ofek/hatch-vcs/archive/refs/tags/v0.5.0.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 436 KiB<br />

## Compilation and Installation

To compile Python module Hatch-vcs, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Hatch-vcs into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user hatch_vcs
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | hatch_vcs |
| Installed Files | /usr/lib/python3.14/site-packages/hatch_vcs and /usr/lib/python3.14/site-packages/hatch_vcs-0.5.0.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Hatchling.md) Python module Hatchling | [HOME](../README.md) | Python module Iniconfig [>>](./Python313-Iniconfig.md) |

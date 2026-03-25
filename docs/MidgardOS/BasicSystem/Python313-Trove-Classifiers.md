# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Pluggy.md) Python module Pluggy | [HOME](../README.md) | Python module Hatchling [>>](./Python313-Hatchling.md) |

## Python module Trove-Classifiers

Name: Python module Trove-Classifiers<br />
Summary: A Python library for classifying projects and releases based on the PEP 301 specification<br />
License: Apache v2<br />
Version: 2026.1.14.14<br />
URL: [https://github.com/pypa/trove-classifiers/archive/refs/tags/2026.1.14.14.tar.gz](https://github.com/pypa/trove-classifiers/archive/refs/tags/2026.1.14.14.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 540 KiB<br />

## Preparation

First, correct an issue where the wheel contains an incorrect version when the calver module is not installed:

```bash
sed -i '/calver/s/^/#/;$iversion="2026.1.14.14"' setup.py
```

## Compilation and Installation

To compile Python module Trove-Classifiers, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps --no-cache-dir $PWD
```

Finally, to install Python module Trove-Classifiers into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist --no-user trove-classifiers
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | trove-classifiers |
| Installed Files | /usr/lib/python3.14/site-packages/trove_classifiers and /usr/lib/python3.14/site-packages/trove_classifiers-2026.1.14.14.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Pluggy.md) Python module Pluggy | [HOME](../README.md) | Python module Hatchling [>>](./Python313-Hatchling.md) |

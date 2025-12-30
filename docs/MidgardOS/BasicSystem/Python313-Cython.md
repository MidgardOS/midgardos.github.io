# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Jinja2.md) Python module Jinja2 | [HOME](../README.md) | LibSecComp 64-bit [>>](./LibSecComp_64bit.md) |

## Python module Cython

Name: Python module Cython<br />
Summary: An optimised Python transpiler for creating C extensions for Python in Python itself<br />
License: Apache v2.0<br />
Version: 3.2.3<br />
URL: [https://github.com/cython/cython/releases/download/3.2.3-1/cython-3.2.3.tar.gz](https://github.com/cython/cython/releases/download/3.2.3-1/cython-3.2.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: iB<br />

## Compilation and Installation

To compile Python module Cython, run the following command:

```bash
pip3 wheel -w dist --no-cache-dir --no-build-isolation --no-deps $PWD
```

Finally, to install Python module Cython into the build tree, run the following command:

```bash
pip3 install --no-index --find-links dist Cython
```

## Contents

| Contents | |
| --- | --- |
| Installed Plugins | Cython |
| Installed Files | /usr/lib64/python3.13/site-packages/Cython and /usr/lib64/python3.13/site-packages/cython-3.2.3.dist-info |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Jinja2.md) Python module Jinja2 | [HOME](../README.md) | LibSecComp 64-bit [>>](./LibSecComp_64bit.md) |

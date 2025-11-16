# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Setuptools.md) Python module Setuptools | [HOME](../README.md) | Meson [>>](./Meson.md) |

## Ninja

Name: Ninja<br />
Summary: A small build system focused on speed<br />
License: Apache v2<br />
Version: 1.13.1<br />
URL: [https://github.com/ninja-build/ninja/archive/refs/tags/v1.13.1.tar.gz](https://github.com/ninja-build/ninja/archive/refs/tags/v1.13.1.tar.gz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 9.1 MiB<br />

## Preparation

To prepare Ninja for install into the build root, run the following commands:

```bash
export NINJAJOBS=4
sed -i '/int Guess/a \
  int   j = 0;\
  char* jobs = getenv( "NINJAJOBS" );\
  if ( jobs != NULL ) j = atoi( jobs );\
  if ( j > 0 ) return j;\
' src/ninja.cc
```

The steps above ensure that the tool will honor the `NINJAJOBS` environment variable for the number of CPU cores to allocate jobs to. Without this, it will force all cores, plus two additional to be spawned.

## Compilation and Installation

To compile Ninja, run the following command:

```bash
python3 configure.py --bootstrap --verbose
```

The test suite requires `cmake`, which is not available at this point.

Finally, to install Ninja into the build tree, run the following command:

```bash
install -vm755 ninja /usr/bin/
install -vDm644 misc/bash-completion /usr/share/bash-completion/completions/ninja
install -vDm644 misc/zsh-completion  /usr/share/zsh/site-functions/_ninja
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | ninja |

| Navigation |||
| --- | --- | ---: |
| [<<](./Python313-Setuptools.md) Python module Setuptools | [HOME](../README.md) | Meson [>>](./Meson.md) |

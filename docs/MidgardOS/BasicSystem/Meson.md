# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Ninja.md) Ninja | [HOME](../README.md) | IP Utils [>>](./IPUtils.md) |

## Meson

Name: Meson<br />
Summary: A fast, user-friendly build system<br />
License: Apache v2<br />
Version: 1.10.2<br />
URL: [https://github.com/mesonbuild/meson/releases/download/1.10.2/meson-1.10.2.tar.gz](https://github.com/mesonbuild/meson/releases/download/1.10.2/meson-1.10.2.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 9.7 MiB<br />

## Compilation and Installation

To compile Meson, run the following command:

```bash
pip3 wheel -w dist --no-build-isolation --no-deps $PWD
```

Finally, to install Meson into the build tree, run the following commands:

```bash
pip3 install --no-index --find-links dist meson
install -vDm644 data/shell-completions/bash/meson /usr/share/bash-completion/completions/meson
install -vDm644 data/shell-completions/zsh/_meson /usr/share/zsh/site-functions/_meson
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | meson |
| Installed Plugins | mesonbuild |

| Navigation |||
| --- | --- | ---: |
| [<<](./Ninja.md) Ninja | [HOME](../README.md) | IP Utils [>>](./IPUtils.md) |

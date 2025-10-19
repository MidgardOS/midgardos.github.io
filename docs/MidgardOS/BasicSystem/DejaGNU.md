# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Expect.md) Expect | [HOME](../README.md) | PkgConf [>>](./PkgConf.md) |

## DejaGNU

Name: DejaGNU<br />
Summary: A framework for writing test suites<br />
License: GPL v3 or later<br />
Version: 1.6.3<br />
URL: [https://ftp.gnu.org/gnu/dejagnu/dejagnu-1.6.3.tar.gz](https://ftp.gnu.org/gnu/dejagnu/dejagnu-1.6.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 6.9 MiB<br />

## Configuration

To configure DejaGNU for install into the build root, run the following commands:

```bash
mkdir -v build && cd build
../configure --prefix=/usr           \
             --libdir=/usr/lib64     \
             --libexecdir=/usr/lib64
makeinfo --html --no-split -o doc/dejagnu.html ../doc/dejagnu.texi
makeinfo --plaintext       -o doc/dejagnu.txt  ../doc/dejagnu.texi
```

A warning will be emitted during the makeinfo call for generating the HTML documentation. This is expected.

## Compilation and Installation

The DejaGNU package does not have a traditional build step, however, there is a test suite that should be run:

```bash
make check
```

Finally, to install DejaGNU into the build tree, run the following command:

```bash
make install
install -v -dm755  /usr/share/doc/dejagnu-1.6.3
install -v -m644   doc/dejagnu.{html,txt} /usr/share/doc/dejagnu-1.6.3
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | dejagnu and runtest |

| Navigation |||
| --- | --- | ---: |
| [<<](./Expect.md) Expect | [HOME](../README.md) | PkgConf [>>](./PkgConf.md) |

# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPatch.md) GNU Patch | [HOME](../README.md) | GNU Texinfo [>>](./GNUTexinfo.md) |

## GNU Tar

Name: GNU Tar<br />
Summary: The GNU Tape Archive tool<br />
License: GPL v<br />
Version: 1.35<br />
URL: [https://ftp.gnu.org/gnu/tar/tar-1.35.tar.xz](https://ftp.gnu.org/gnu/tar/tar-1.35.tar.xz)<br />

Average Build Time: 0.6 SBU<br />
Used Install Space: 5 MiB<br />

## Configuration

To configure GNU Tar for install into the build root, run the following command:

```bash
FORCE_UNSAFE_CONFIGURE=1                        \
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/tar-1.35    \
            --sysconfdir=/etc                   \
            --disable-rpath
```

## Compilation and Installation

To compile GNU Tar, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Currently, there are two tests that fail. These are a test for lzip support, which needs extra tooling in the tree to work, and support for handling Linux Capabilities. Further investigation is needed as to why this test is failing. For now, it is safe to continue if these two tests are failing.

Finally, to install GNU Tar into the build tree, run the following commands:

```bash
make install
make -C doc install-html docdir=/usr/share/doc/tar-1.35
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | tar and rmt |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUPatch.md) GNU Patch | [HOME](../README.md) | GNU Texinfo [>>](./GNUTexinfo.md) |

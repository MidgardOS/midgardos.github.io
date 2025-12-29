# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibPipeline_32bit.md) LibPipeline 32-bit | [HOME](../README.md) | GNU Patch [>>](./GNUPatch.md) |

##  GNU Make

Name: GNU Make<br />
Summary: A small processing tool for orchestrating actions<br />
License: GPL v3.0 or later<br />
Version: 4.4.1<br />
URL: [https://ftp.gnu.org/gnu/make/make-4.4.1.tar.gz](https://ftp.gnu.org/gnu/make/make-4.4.1.tar.gz)<br />

Average Build Time: 0.7 SBU<br />
Used Install Space: 2.8 MiB<br />

## Configuration

To configure GNU Make for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/make-4.4.1  \
            --sysconfdir=/etc                   \
            --disable-rpath                     \
            --enable-year2038
```

## Compilation and Installation

To compile GNU Make, run the following command:

```bash
make
```

Next, run the test suite:

```bash
useradd -c "Test User" -u 1000 -U -m tester
groupadd -g 102 dummy -U tester
unalias cp
cp -R ../make-4.4.1 /tmp/
alias cp="cp -i"
chown -R tester /tmp/make-4.4.1/
cd /tmp/make-4.4.1
su tester -c "PATH=$PATH autoreconf -fiv && make check"
cd -
rm -rf /tmp/make-4.4.1
groupdel -f dummy
userdel -rf tester
```

Finally, to install GNU Make into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | make |

| Navigation |||
| --- | --- | ---: |
| [<<](./LibPipeline_32bit.md) LibPipeline 32-bit | [HOME](../README.md) | GNU Patch [>>](./GNUPatch.md) |

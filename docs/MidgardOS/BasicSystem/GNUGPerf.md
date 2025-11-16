# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGDBM_32bit.md) GNU GDBM 32-bit | [HOME](../README.md) | Expat [>>](./Expat_64bit.md) |

## GNU GPerf

Name: GNU GPerf<br />
Summary: A perfect hash function generator<br />
License: GPL v3<br />
Version: 3.3<br />
URL: [https://ftp.gnu.org/gnu/gperf/gperf-3.3.tar.gz](https://ftp.gnu.org/gnu/gperf/gperf-3.3.tar.gz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 12 MiB<br />

## Configuration

To configure GNU GPerf for install into the build root, run the following command:

```bash
./configure --prefix=/usr \
            --docdir=/usr/share/doc/gperf-3.3
```

## Compilation and Installation

To compile GNU GPerf, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU GPerf into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | gperf |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGDBM_32bit.md) GNU GDBM 32-bit | [HOME](../README.md) | Expat [>>](./Expat_64bit.md) |

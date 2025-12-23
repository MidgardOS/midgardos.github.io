# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Groff.md) GNU Roff | [HOME](../README.md) | IPRoute2 [>>](./IPRoute2.md) |

## GZip

Name: GZip<br />
Summary: Programs for compressing and decompressing files<br />
License: GPL v3 or later<br />
Version: 1.14<br />
URL: [https://ftp.gnu.org/gnu/gzip/gzip-1.14.tar.xz](https://ftp.gnu.org/gnu/gzip/gzip-1.14.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 560 KiB<br />

## Configuration

To configure GZip for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --docdir=/usr/share/doc/gzip-1.14
```

## Compilation and Installation

To compile GZip, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GZip into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | gunzip, gzexe, gzip, uncompress, zcat, zcmp, zdiff, zegrep, zfgrep, zforce, zgrep, zless, zmore, and znew |

| Navigation |||
| --- | --- | ---: |
| [<<](./Groff.md) GNU Roff | [HOME](../README.md) | IPRoute2 [>>](./IPRoute2.md) |

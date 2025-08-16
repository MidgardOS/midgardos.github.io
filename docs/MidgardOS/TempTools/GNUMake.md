# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGzip.md) | [HOME](../README.md) | [>>](./GNUPatch.md) |

## GNU Make

Name: GNU Make<br />
Summary: A powerful orchestration tool for building software and running tasks<br />
License: GPL v3 or later<br />
Version: 4.4<br />
URL: [https://ftp.gnu.org/gnu/make/make-4.4.tar.gz](https://ftp.gnu.org/gnu/make/make-4.4.tar.gz)<br />

## Configuration

To configure GNU Make for install into the build root, run the following command:

```bash
./configure --prefix=/usr       \
            --host=$BRFS_TARGET \
            --build=$BRFS_HOST
```

## Compilation and Installation

To compile GNU Make, run the following command:

```bash
make
```

Finally, to install GNU Make into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGzip.md) | [HOME](../README.md) | [>>](./GNUPatch.md) |

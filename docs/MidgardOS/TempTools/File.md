# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUDiffutils.md) | [HOME](../README.md) | [>>](./GNUFindutils.md) |

## File

Name: File<br />
Summary: A tool to determine the type of a file from heuristics<br />
License: BSD 2-Clause<br />
Version: 5.46<br />
URL: [http://ftp.astrom.com/pub/file/](ftp://ftp.astron.com/pub/file/)<br />

## Pre-Build

Unfortunately, File requires being built twice to ensure the "magic" signatures files are generated correctly. To do the first-pass build, run the following commands inside the unpacked source tree:

```bash
mkdir build
pushd build
    ../configure --disable-bzlib      \
                 --disable-libseccomp \
                 --disable-xzlib      \
                 --disable-zlib
    make
popd
```

## Configuration

To configure File for install into the build root, run the following command:

```bash
./configure             \
    --prefix=/usr       \
    --host=$BRFS_TARGET \
    --build=$BRFS_HOST
```

## Compilation and Installation

To compile File, run the following command:

```bash
make FILE_COMPILE=$(pwd)/build/src/file
```

Finally, to install File into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUDiffutils.md) | [HOME](../README.md) | [>>](./GNUFindutils.md) |

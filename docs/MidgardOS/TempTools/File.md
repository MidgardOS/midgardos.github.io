# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUDiffutils.md) GNU Diffutils | [HOME](../README.md) | GNU Findutils [>>](./GNUFindutils.md) |

## File

Name: File<br />
Summary: A tool to determine the type of a file from heuristics<br />
License: BSD 2-Clause<br />
Version: 5.48<br />
URL: [ftp://ftp.astron.com/pub/file/file-5.48.tar.gz](ftp://ftp.astron.com/pub/file/file-5.48.tar.gz)<br />

## Pre-Build

Unfortunately, File requires being built twice to ensure the "magic" signature files are generated correctly. To do the first-pass build, run the following commands inside the unpacked source tree:

```bash
autoreconf -fvi
mkdir build && cd build
../configure --libdir=/usr/lib64  \
             --disable-bzlib      \
             --disable-libseccomp \
             --disable-xzlib      \
             --disable-zlib
make
cd ../
```

## Configuration

To configure File for installation into the build root, run the following command:

```bash
./configure             \
    --prefix=/usr       \
    --libdir=/usr/lib64 \
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

More details about this package are covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUDiffutils.md) GNU Diffutils | [HOME](../README.md) | GNU Findutils [>>](./GNUFindutils.md) |

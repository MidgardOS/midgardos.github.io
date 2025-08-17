# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTar.md) GNU Tar | [HOME](../README.md) | GNU Binutils - pass 2 [>>](./GNUBinutils.md) |

## Xz

Name: Xz<br />
Summary: A Better Compression Tool<br />
License: Multiple FOSS licenses (0BSD/GPLv2+/GPLv3+)<br />
Version: 5.8.1<br />
URL: [https://github.com/tukaani-project/xz/releases/download/v5.8.1/xz-5.8.1.tar.xz](https://github.com/tukaani-project/xz/releases/download/v5.8.1/xz-5.8.1.tar.xz)<br />

## Configuration

To configure Xz for install into the build root, run the following command:

```bash
./configure --prefix=/usr                    \
            --libdir=/usr/lib64              \
            --libexecdir=/usr/lib64          \
            --host=$BRFS_TARGET              \
            --build=$BRFS_HOST               \
            --disable-static                 \
            --docdir=/usr/share/doc/xz-5.8.1
```

## Compilation and Installation

To compile Xz, run the following command:

```bash
make
```

Finally, to install Xz into the build tree, run the following command:

```bash
make DESTDIR=$BRFS install
rm -v $BRFS/usr/lib64/liblzma.la
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTar.md) GNU Tar | [HOME](../README.md) | GNU Binutils - pass 2 [>>](./GNUBinutils.md) |

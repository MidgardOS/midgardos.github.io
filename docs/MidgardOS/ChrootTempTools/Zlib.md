# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | MPDecimal [>>](./MPDecimal.md) |

## ZLib

Name: ZLib<br />
Summary: A library of compression and decompression algorithms<br />
License: ZLib License<br />
Version: 1.3.2<br />
URL: [https://zlib.net/fossils/zlib-1.3.2.tar.gz](https://zlib.net/fossils/zlib-1.3.2.tar.gz)<br />

## Configuration

To configure ZLib for installation into the build root, run the following command:

```bash
./configure --prefix=/usr
```

## Compilation and Installation

To compile ZLib, run the following command:

```bash
make
```

Finally, to install ZLib into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libz.a
```

More details about this package are covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./Perl.md) Perl | [HOME](../README.md) | MPDecimal [>>](./MPDecimal.md) |

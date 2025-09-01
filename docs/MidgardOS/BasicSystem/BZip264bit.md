# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib32bit.md) ZLib 32-bit | [HOME](../README.md) | BZip2 32-bit [>>](./BZip232bit.md) |

## BZip2 64-bit

Name: BZip2 64-bit<br />
Summary: An efficient file compression tool<br />
License: BSD-3-Clause<br />
Version: 1.0.8<br />
URL: [https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz](https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz)<br />

Average Build Time: less than 0.1 SBU
Used Install Space: 7.3 MiB

## Preparation

To allow the documentation for BZip2 to be installed, apply the following patch:

```bash
patch -Np1 -i ../bzip2-1.0.8-install_docs-1.patch
```

Next, run the following command to ensure that symlinks installed are relative:

```bash
sed -i 's@\(ln -s -f \)$(PREFIX)/bin/@\1@' Makefile
```

Now, ensure that the man pages are installed into the appropriate location:

```bash
sed -i "s@(PREFIX)/man@(PREFIX)/share/man@g" Makefile
```

## Compilation and Installation

After preparing BZip2 64-bit for build, run the following command to generate a dynamic library that the `bzip2` command will be linked against:

```bash
make -f Makefile-libbz2_so
make clean
```

Now compile BZip2 64-bit:

```bash
make
```

Finally, to install BZip2 64-bit into the build tree, run the following commands:

```bash
make PREFIX=/usr install
cp -av libbz2.so.* /usr/lib64
ln -sv libbz2.so.1.0.8 /usr/lib64/libbz2.so
  ln -sfv bzip2 $i
done
rm -fv /usr/lib64/libbz2.a
```

**NOTE: Do not delete the unpacked sources**

# Contents

| Contents | |
| --- | --- |
| Installed Programs | bunzip2 (link to bzip2), bzcat (link to bzip2), bzcmp (link to bzdiff), bzdiff, bzegrep (link to bzgrep), bzfgrep (link to bzgrep), bzgrep, bzip2, bzip2recover, bzless (link to bzmore), and bzmore |
| Installed Libraries | libbz2.so |
| Installed Directories | /usr/share/doc/bzip2-1.0.8 |

| Navigation |||
| --- | --- | ---: |
| [<<](./ZLib32bit.md) ZLib 32-bit | [HOME](../README.md) | BZip2 32-bit [>>](./BZip232bit.md) |

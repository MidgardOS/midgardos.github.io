# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./BZip2_32bit.md) BZip2 32-bit | [HOME](../README.md) | XZ 32-bit [>>](./XZ_32bit.md) |

## XZ 64-bit

Name: XZ 64-bit<br />
Summary: A Better Compression Tool<br />
License: Multiple FOSS licenses (0BSD/GPLv2+/GPLv3+)<br />
Version: 5.8.1<br />
URL: [https://github.com/tukaani-project/xz/releases/download/v5.8.1/xz-5.8.1.tar.xz](https://github.com/tukaani-project/xz/releases/download/v5.8.1/xz-5.8.1.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 24 MiB<br />

## Configuration

To configure XZ 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --docdir=/usr/share/doc/xz-5.8.1
```

## Compilation and Installation

To compile XZ 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install XZ 64-bit into the build tree, run the following command:

```bash
make install
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | lzcat (link to xz), lzcmp (link to xzdiff), lzdiff (link to xzdiff), lzegrep (link to xzgrep), lzfgrep (link to xzgrep), lzgrep (link to xzgrep), lzless (link to xzless), lzma (link to xz), lzmadec, lzmainfo, lzmore (link to xzmore), unlzma (link to xz), unxz (link to xz), xz, xzcat (link to xz), xzcmp (link to xzdiff), xzdec, xzdiff, xzegrep (link to xzgrep), xzfgrep (link to xzgrep), xzgrep, xzless, and xzmore |
| Installed Libraries | liblzma.so |
| Installed Directories | /usr/include/lzma and /usr/share/doc/xz-5.8.1 |

| Navigation |||
| --- | --- | ---: |
| [<<](./BZip2_32bit.md) BZip2 32-bit | [HOME](../README.md) | XZ 32-bit [>>](./XZ_32bit.md) |

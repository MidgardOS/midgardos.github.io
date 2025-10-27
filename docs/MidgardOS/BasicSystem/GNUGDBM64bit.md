# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBash.md) GNU Bash | [HOME](../README.md) | GNU GDBM 32-bit [>>](./GNUGDBM32bit.md) |

## GNU GDBM 64-bit

Name: GNU GDBM 64-bit<br />
Summary: The GNU Database Manager<br />
License: GPL v3<br />
Version: 1.26<br />
URL: [https://ftp.gnu.org/gnu/gdbm/gdbm-1.26.tar.gz](https://ftp.gnu.org/gnu/gdbm/gdbm-1.26.tar.gz)<br />

Average Build Time: less than 0.2 SBU<br />
Used Install Space: 13 MiB<br />

## Configuration

To configure GNU GDBM 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --enable-libgdbm-compat
```

## Compilation and Installation

To compile GNU GDBM 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU GDBM 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libgdbm.la
rm -fv /usr/lib64/libgdbm_compat.la
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | gdbm_dump, gdbm_load, and gdbmtool |
| Installed Libraries | libgdbm.so and libgdbm_compat.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBash.md) GNU Bash | [HOME](../README.md) | GNU GDBM 32-bit [>>](./GNUGDBM32bit.md) |

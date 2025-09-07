# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPC32bit.md) GNU MPC 32-bit | [HOME](../README.md) | ISL 32-bit [>>](./ISL32bit.md) |

# Integer Set Library

Name: ISL 64-bit<br />
Summary: A library for handling Integer Sets bounded by linear constraints<br />
License: MIT<br />
Version: 0.27<br />
URL: [https://libisl.sourceforge.io/isl-0.27.tar.xz](https://libisl.sourceforge.io/isl-0.27.tar.xz)<br />

Average Build Time: 0.1 SBU<br />
Used Install Space: 20 MiB<br />

## Configuration

To configure ISL 64-bit for install into our cross-compilation root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --docdir=/usr/share/doc/isl-0.27
```

## Compilation and Installation

To compile ISL 64-bit, run the following commands:

```bash
make
make html
```

Next, run the test suite:

```bash
make check ||:
```

The Python test is failing due to the test being based on python2.x syntax and is safe to ignore.

Finally, to install ISL 64-bit into the cross-tools tree, run the following commands:

```bash
make install
make install-html
install -vd /usr/share/doc/isl-0.27
install -vm644 doc/{CodingStyle,manual.pdf,SubmittingPatches,user.pod} /usr/share/doc/isl-0.27
rm -fv /usr/lib64/libisl.la
mkdir -pv /usr/share/gdb/auto-load/usr/lib64
mv -v /usr/lib64/libisl*gdb.py /usr/share/gdb/auto-load/usr/lib64
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Libraries | libisl.so |
| Installed Directories | /usr/share/doc/isl-0.27 |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUMPC32bit.md) GNU MPC 32-bit | [HOME](../README.md) | ISL 32-bit [>>](./ISL32bit.md) |

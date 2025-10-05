# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./PCRE264bit.md) PCRE2 64-bit | [HOME](../README.md) | GNU M4 [>>](./GNUM4.md) |

## PCRE2 32-bit

Name: PCRE2 32-bit<br />
Summary: Perl-compatible regular expression library<br />
License: BSD-3-Clause with PCRE2-exception<br />
Version: 10.46<br />
URL: [https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.46/pcre2-10.46.tar.gz](https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.46/pcre2-10.46.tar.gz)<br />

Average Build Time: 0.5 SBU<br />
Used Install Space: 20 MiB<br />

## Configuration

To configure PCRE2 32-bit for install into the build root, run the following command:

```bash
make distclean
CC="gcc -m32" ./configure \
            --host=i686-pc-linux-gnu            \
            --prefix=/usr                       \
            --libdir=/usr/lib                   \
            --libexecdir=/usr/lib               \
            --docdir=/usr/share/doc/pcre2-10.45 \
            --enable-unicode                    \
            --enable-jit                        \
            --enable-pcre2-16                   \
            --enable-pcre2-32                   \
            --enable-pcre2grep-libz             \
            --enable-pcre2grep-libbz2           \
            --enable-pcre2test-libreadline      \
            --disable-static
```

## Compilation and Installation

To compile PCRE2 32-bit, run the following command:

```bash
make
```

Finally, to install PCRE2 32-bit into the build tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib
rm -fv /usr/lib/libpcre2-*.la
```

## Contents

See the contents section of the 64-bit build of PCRE2 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./PCRE264bit.md) PCRE2 64-bit | [HOME](../README.md) | GNU M4 [>>](./GNUM4.md) |

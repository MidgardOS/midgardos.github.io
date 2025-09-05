# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUReadline32bit.md) GNU Readline 32-bit | [HOME](../README.md) | PCRE2 32-bit [>>](./PCRE232bit.md) |

## PCRE2 64-bit

Name: PCRE2 64-bit<br />
Summary: Perl-compatible regular expression library<br />
License: BSD-3-Clause with PCRE2-exception<br />
Version: 10.46<br />
URL: [https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.46/pcre2-10.46.tar.gz](https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.46/pcre2-10.46.tar.gz)<br />

Average Build Time: 0.5 SBU<br />
Used Install Space: 20 MiB<br />

## Configuration

To configure PCRE2 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
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

To compile PCRE2 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install PCRE2 64-bit into the build tree, run the following command:

```bash
make install
rm -fv /usr/lib64/libpcre2-*.la
```

**NOTE: Do not delete the unpacked sources**

| Contents ||
| --- | --- |
| Installed Programs | pcre2grep and pcre2test |
| Installed Libraries | libpcre2-8.so, libpcre2-16.so, libpcre2-32.so, and libpcre2-posix.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUReadline32bit.md) GNU Readline 32-bit | [HOME](../README.md) | PCRE2 32-bit [>>](./PCRE232bit.md) |

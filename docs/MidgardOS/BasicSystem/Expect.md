# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./TCL.md) TCL | [HOME](../README.md) | DejaGNU [>>](./DejaGNU.md) |

## Expect

Name: Expect<br />
Summary: An automation framework based off console strings/output<br />
License: Public Domain<br />
Version: 5.45.4<br />
URL: [https://sourceforge.net/projects/expect/files/Expect/5.45.4/expect5.45.4.tar.gz/download](https://sourceforge.net/projects/expect/files/Expect/5.45.4/expect5.45.4.tar.gz/download)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 3.9 MiB<br />

**NOTE: Expect requires a fully functioning TTY layer inside the chroot environment. Verify this before proceeding with the build**

## Preparation

Expect requires a patch to build correctly with GCC 15 and above:

```bash
patch -Np1 -i ../expect-5.45.4-gcc15-1.patch
```

## Configuration

To configure Expect for install into the build root, run the following command:

```bash
./configure --prefix=/usr             \
            --libdir=/usr/lib64       \
            --libexecdir=/usr/lib64   \
            --with-tcl=/usr/lib64/tcl \
            --enable-shared           \
            --disable-rpath           \
            --mandir=/usr/share/man   \
            --with-tclinclude=/usr/include
```

## Compilation and Installation

To compile Expect, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make test
```

Finally, to install Expect into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
pushd DESTDIR
    mkdir -pv usr/lib64/tcl
    mv usr/lib64/expect* usr/lib64/tcl/
    cp -Rv usr/bin/* /usr/bin/
    cp -Rv usr/include/* /usr/include/
    cp -Rv usr/lib64/tcl/* /usr/lib64/tcl/
    cp -Rv usr/share/man/man1/* /usr/share/man/man1/
    cp -Rv usr/share/man/man3/* /usr/share/man/man3/
popd
ln -svf tcl/expect5.45.4/libexpect5.45.4.so /usr/lib64
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | autoexpect, autopasswd, cryptdir, decryptdir, dislocate, expect, ftp-rfc, kibitz, lpunlock, mkpasswd, multixterm, passmass, rftp, rlogin-cwd, timed-read, timed-run, tknewsbiff, tkpasswd, unbuffer, weather, xkibitz, xpstat |
| Installed Libraries | libexpect5.45.4.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./TCL.md) TCL | [HOME](../README.md) | DejaGNU [>>](./DejaGNU.md) |

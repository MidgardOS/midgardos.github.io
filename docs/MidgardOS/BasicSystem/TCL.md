# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./Flex_32bit.md) Flex 32-bit | [HOME](../README.md) | Expect [>>](./Expect.md) |

## TCL

Name: TCL<br />
Summary: The Tool Command Language interpreter<br />
License: TCL License<br />
Version: 8.6.17<br />
URL: [https://downloads.sourceforge.net/tcl/tcl8.6.17-src.tar.gz](https://downloads.sourceforge.net/tcl/tcl8.6.17-src.tar.gz)<br />
Documentation URL: [https://downloads.sourceforge.net/tcl/tcl8.6.17-html.tar.gz](https://downloads.sourceforge.net/tcl/tcl8.6.17-html.tar.gz)<br />

Average Build Time: 3.0 SBU<br />
Used Install Space: 91 MiB<br />

## Configuration

To configure TCL for install into the build root, run the following commands:

```bash
SRCDIR=$(pwd)
rm -rfv pkgs/sqlite*
cd unix
export TCL_LIBRARY="/usr/lib64/tcl/tcl8.6"
export TCL_PACKAGE_PATH="/usr/lib64/tcl:/usr/lib64/tcl/tcl8.6:/usr/share/tcl"
./configure --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
            --mandir=/usr/share/man                     \
            --enable-64bit                              \
            --disable-rpath                             \
            MODULE_INSTALL_DIR="/usr/lib64/tcl/tcl8.6"  \
            PACKAGE_DIR="/usr/lib64/tcl/tcl8.6"
```

**NOTE: The sqlite3 extension is removed, since it is provided by the SQLite3 package**

## Compilation and Installation

To compile TCL, run the following commands:

```bash
make
sed -e "s|$SRCDIR/unix|/usr/lib64|" \
    -e "s|$SRCDIR|/usr/include|"  \
    -i tclConfig.sh
sed -e "s|$SRCDIR/unix/pkgs/tdbc1.1.12|/usr/lib64/tcl/tcl8.6/tdbc1.1.12|" \
    -e "s|/usr/lib64/tdbc1.1.12|/usr/lib64/tcl/tcl8.6/tdbc1.1.12|g" \
    -e "s|$SRCDIR/pkgs/tdbc1.1.12/generic|/usr/include|"     \
    -e "s|$SRCDIR/pkgs/tdbc1.1.12/library|/usr/lib64/tcl/tcl8.6|"  \
    -e "s|$SRCDIR/pkgs/tdbc1.1.12|/usr/include|"             \
    -i pkgs/tdbc1.1.12/tdbcConfig.sh
sed -e "s|$SRCDIR/unix/pkgs/itcl4.3.4|/usr/lib64/tcl/tcl8.6/itcl4.3.4|" \
    -e "s|/usr/lib64/itcl4.3.4|/usr/lib64/tcl/tcl8.6/itcl4.3.4|g" \
    -e "s|$SRCDIR/pkgs/itcl4.3.4/generic|/usr/include|"    \
    -e "s|$SRCDIR/pkgs/itcl4.3.4|/usr/include|"            \
    -i pkgs/itcl4.3.4/itclConfig.sh
unset SRCDIR
unset TCL_LIBRARY
unset TCL_PACKAGE_PATH
```

Next, run the test suite:

```bash
make test
```

Currently, there are three tests that fail. This is reasonable

Finally, to install TCL into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install install-private-headers
pushd DESTDIR
    mv usr/lib64/tclConfig.sh usr/lib64/tclooConfig.sh usr/lib64/tcl/
    mv usr/lib64/tcl8/8.6/* usr/lib64/tcl/tcl8/8.6/
    mv usr/lib64/itcl* usr/lib64/tcl/
    mv usr/lib64/tdbc* usr/lib64/tcl/
    mv usr/lib64/thread2* usr/lib64/tcl/
    chmod -v 644 usr/lib64/libtclstub8.6.a
    chmod -v 755 usr/lib64/libtcl8.6.so
    mv usr/share/man/man3/{Thread,Tcl_Thread}.3
    rm -rfv usr/lib64/tcl8
    cp -Rv usr/bin/* /usr/bin/
    cp -Rv usr/include/* /usr/include/
    cp -Rv usr/lib64/* /usr/lib64/
    cp -Rv usr/share/man/man1/* /usr/share/man/man1/
    cp -Rv usr/share/man/man3/* /usr/share/man/man3/
    cp -Rv usr/share/man/mann/* /usr/share/man/mann/
popd
install -v -d -m 755 -o root -g root /usr/share/tcl
ln -sfv tclsh8.6 /usr/bin/tclsh
cd ..
tar -xf ../tcl8.6.17-html.tar.gz --strip-components=1
mkdir -v -p /usr/share/doc/tcl-8.6
cp -v -r  ./html/* /usr/share/doc/tcl-8.6
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | tclsh (link to tclsh8.6) and tclsh8.6 |
| Installed Libraries | libtcl8.6.so and libtclstub8.6.a |
| Installed Directories | /usr/share/tcl |

| Navigation |||
| --- | --- | ---: |
| [<<](./Flex_32bit.md) Flex 32-bit | [HOME](../README.md) | Expect [>>](./Expect.md) |

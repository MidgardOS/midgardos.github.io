# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses_64bit.md) NCurses 64-bit | [HOME](../README.md) | GNU Sed [>>](./GNUSed.md) |

## NCurses 32-bit

Name: NCurses 32-bit<br />
Summary: Terminal-independent library for console interfaces<br />
License: MIT<br />
Version: 6.5-20250823<br />
URL: [https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250823.tgz](https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250823.tgz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 46 MiB<br />

## Configuration

To configure NCurses 32-bit for install into the build root, run the following commands:

```bash
make distclean
CC="gcc -m32" CXX="g++ -m32"         \
./configure --prefix=/usr            \
            --host=i686-pc-linux-gnu \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --mandir=/usr/share/man  \
            --with-shared            \
            --without-debug          \
            --without-normal         \
            --with-cxx-shared        \
            --enable-pc-files        \
            --with-pkg-config-libdir=/usr/lib/pkgconfig
```

## Compilation and Installation

To compile NCurses 32-bit, run the following command:

```bash
make
```

Finally, to install NCurses 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
mkdir -p DESTDIR/usr/lib/pkgconfig
for lib in ncurses form panel menu ; do
    rm -vf                    DESTDIR/usr/lib/lib${lib}.so
    echo "INPUT(-l${lib}w)" > DESTDIR/usr/lib/lib${lib}.so
    ln -svf ${lib}w.pc        DESTDIR/usr/lib/pkgconfig/$lib.pc
done
rm -vf                     DESTDIR/usr/lib/libcursesw.so
echo "INPUT(-lncursesw)" > DESTDIR/usr/lib/libcursesw.so
ln -sfv libncurses.so      DESTDIR/usr/lib/libcurses.so
cp -Rfv DESTDIR/usr/lib/* /usr/lib
rm -rf DESTDIR
```

## Contents

See the contents section of the 64-bit build of NCurses for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses_64bit.md) NCurses 64-bit | [HOME](../README.md) | GNU Sed [>>](./GNUSed.md) |

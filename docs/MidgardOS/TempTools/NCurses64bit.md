# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./NCursesTic.md) NCurses TIC | [HOME](../README.md) | NCurses 32-bit [>>](./NCurses32bit.md) |

## NCurses 64-bit

Name: NCurses 64-bit<br />
Summary: The Ncurses package contains libraries for terminal-independent handling of character screens - 64-bit build<br />
License: MIT<br />
Version: 6.5-20250823<br />
URL: [https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250823.tgz](https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250823.tgz)<br />

## Configuration

To configure NCurses 64-bit for install into our build root, run the following command:

```bash
./configure --prefix=/usr                \
            --host=$BRFS_TARGET          \
            --build=$BRFS_HOST           \
            --mandir=/usr/share/man      \
            --with-manpage-format=normal \
            --with-shared                \
            --without-normal             \
            --with-cxx-shared            \
            --without-debug              \
            --without-ada                \
            --disable-stripping          \
            AWK=gawk
```

## Compilation and Installation

To compile NCurses 64-bit, run the following command:

```bash
make
```

Finally, to install NCurses 64-bit into the build root tree, run the following command:

```bash
make DESTDIR=${BRFS} TIC_PATH=${BRFS}/tools/bin/tic install
ln -sv libncursesw.so $BRFS/usr/lib/libncurses.so
sed -e 's/^#if.*XOPEN.*$/#if 1/' -i $BRFS/usr/include/curses.h
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./NCursesTic.md) NCurses TIC | [HOME](../README.md) | NCurses 32-bit [>>](./NCurses32bit.md) |

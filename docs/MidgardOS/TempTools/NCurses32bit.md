# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses64bit.md) NCurses 64-bit | [HOME](../README.md) | GNU Bash [>>](./GNUBash.md) |

## NCurses 32-bit

Name: NCurses 32-bit<br />
Summary: The Ncurses package contains libraries for terminal-independent handling of character screens - 32-bit build<br />
License: MIT<br />
Version: 6.5-20250802<br />
URL: [https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250802.tgz](https://invisible-mirror.net/archives/ncurses/current/ncurses-6.5-20250802.tgz)<br />

## Configuration

To configure NCurses 32-bit for install into our build root, run the following command:

```bash
CC="$BRFS_TARGET-gcc -m32" CXX="$BRFS_TARGET-g++ -m32" \
./configure --prefix=/usr                              \
            --host=$BRFS_TARGET32                      \
            --build=$(./config.guess)                  \
            --libdir=/usr/lib                          \
            --mandir=/usr/share/man                    \
            --with-shared                              \
            --without-normal                           \
            --with-cxx-shared                          \
            --without-debug                            \
            --without-ada                              \
            --disable-stripping                        \
            AWK=gawk
```

## Compilation and Installation

To compile NCurses 32-bit, run the following command:

```bash
make
```

Finally, to install NCurses 32-bit into the build root tree, run the following command:

```bash
make DESTDIR=$PWD/DESTDIR TIC_PATH=${BRFS}/tools/bin/tic install
ln -sv libncursesw.so DESTDIR/usr/lib/libncurses.so
cp -Rv DESTDIR/usr/lib/* $BRFS/usr/lib
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./NCurses64bit.md) NCurses 64-bit | [HOME](../README.md) | GNU Bash [>>](./GNUBash.md) |

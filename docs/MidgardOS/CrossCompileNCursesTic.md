| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGNUM4.md) | [HOME](./README.md) | [>>](./CrossCompilePkgConf.md) |

# GNU NCurses TIC

Name: NCurses - Tic<br />
Summary: A terminal console control library - Terminal Interface Compiler<br />
License: MIT<br />
Version: 6.4<br />
URL: [http://www.invisible-island.net/ncurses/](http://www.invisible-island.net/ncurses)<br />

## Notes

This package will only install the `tic` binary, as it cannot be built and ran safely in a cross-compilation process as
the build of NCurses has a bit of a chicken-and-egg dependency.

## Configuration

To configure NCurses' `tic` for install into our cross-compilation root, run the following command:

```bash
AWK=gawk ./configure --prefix=/cross-tools --without-debug
```

## Compilation and Installation

To compile NCurses `tic`, run the following command:

```bash
make -C include
make -C progs tic
```

Finally, to install NCurses `tic` into the cross-tools tree, run the following command:

```bash
install -v -m755 progs/tic /cross-tools/bin
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGNUM4.md) | [HOME](./README.md) | [>>](./CrossCompilePkgConf.md) |

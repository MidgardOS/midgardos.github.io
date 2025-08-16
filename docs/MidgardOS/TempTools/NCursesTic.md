# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./M4.md) | [HOME](../README.md) | [>>](./NCurses64bit.md) |

## NCurses TIC

Name: NCurses - Tic<br />
Summary: A terminal console control library - Terminal Interface Compiler<br />
License: MIT<br />
Version: 6.5-20250802<br />
URL: [http://www.invisible-island.net/ncurses/](http://www.invisible-island.net/ncurses)<br />

## Notes

This step will only install the `tic` binary, as NCurses cannot be built and ran safely without tic already installed during the cross-compilation process as the build of NCurses has a circular dependency.

## Configuration

To configure NCurses' `tic` for install into our cross-compilation root, run the following command:

```bash
AWK=gawk ./configure --prefix=/tools --without-debug
```

## Compilation and Installation

To compile NCurses `tic`, run the following command:

```bash
make -C include
make -C progs tic
```

Finally, to install NCurses `tic` into the temporary tools tree, run the following command:

```bash
install -v -m755 progs/tic /tools/bin
```

To work around an issue with finding the (potentially) newer version of GLibC, set the ELF RPATH for the `tic` binary:

```bash
patchelf --add-rpath ${BRFS}/lib64 ${BRFS}/bin/tic
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./M4.md) | [HOME](../README.md) | [>>](./NCurses64bit.md) |

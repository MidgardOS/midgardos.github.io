# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUFindutils.md) GNU Findutils | [HOME](../README.md) | GZip [>>](./GZip.md) |

## GNU Roff

Name: GNU Roff<br />
Summary: Programs for formatting text and images<br />
License: GPL v3 or later<br />
Version: 1.23.0<br />
URL: [https://ftp.gnu.org/gnu/groff/groff-1.23.0.tar.gz](https://ftp.gnu.org/gnu/groff/groff-1.23.0.tar.gz)<br />

Average Build Time: 0.2 SBU<br />
Used Install Space: 35 MiB<br />

## Notes

GNU Roff can take advantage of a few other tools, and should be rebuilt once they are available in MidgardOS:

- X11
- URW Fonts
- The uchardat library
- Ghostscript

## Configuration

To configure GNU Roff for install into the build root, run the following command:

```bash
PAGE="letter" ./configure --prefix=/usr           \
                          --libdir=/usr/lib64     \
                          --libexecdir=/usr/lib64 \
                          --docdir=/usr/share/doc/groff-1.23.0
```

## Compilation and Installation

To compile GNU Roff, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Roff into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | addftinfo, afmtodit, chem, eqn, eqn2graph, gdiffmk, glilypond, gperl, gpinyin, grap2graph, grn, grodvi, groff, grog, grolbp, grolj4, gropdf, grops, grotty, hpftodit, indxbib, lkbib, lookbib, mmroff, neqn, nroff, pdfmom, pdfroff, pfbtops, pic, pic2graph, post-grohtml, pre-grohtml, preconv, refer, soelim, tbl, tfmtodit, and troff |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUFindutils.md) GNU Findutils | [HOME](../README.md) | GZip [>>](./GZip.md) |

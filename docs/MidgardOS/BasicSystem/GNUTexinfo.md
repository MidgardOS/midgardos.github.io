# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTar.md) GNU Tar | [HOME](../README.md) | Vi Improved [>>](./Vim.md) |

## GNU Texinfo

Name: GNU Texinfo<br />
Summary: A tool for reading, writing, and converting info pages to various formats<br />
License: <br />
Version: 7.2<br />
URL: [https://ftp.gnu.org/gnu/texinfo/texinfo-7.2.tar.xz](https://ftp.gnu.org/gnu/texinfo/texinfo-7.2.tar.xz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 16 MiB<br />

## Configuration

To configure GNU Texinfo for install into the build root, run the following commands:

```bash
sed 's/! $output_file eq/$output_file ne/' -i tp/Texinfo/Convert/*.pm
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --docdir=/usr/share/doc/texinfo-7.2 \
            --sysconfdir=/etc                   \
            --disable-rpath                     \
            --enable-year2038
```

## Compilation and Installation

To compile GNU Texinfo, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Texinfo into the build tree, run the following commands:

```bash
make install
install -v -m 755 -o root -g root ../system_files/usr/bin/geninfodir.sh /usr/bin/
make TEXMF=/usr/share/texmf install-tex
/usr/bin/geninfodir.sh
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | info, install-info, makeinfo, pdftexi2dvi, pod2texi, texi2any, texi2dvi, texi2pdf, and texindex |
| Installed Libraries | libtexinfo-convert.so, libtexinfo-convertxs.so, libtexinfo.so, libtexinfoxs.so |
| Installed Plugins | ConvertXS.so, DocumentXS.so, IndicesXS.so, MiscXS.so, Parsetexi.so, StructuringTransfoXS.so, XSParagraph.so, |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUTar.md) GNU Tar | [HOME](../README.md) | Vi Improved [>>](./Vim.md) |

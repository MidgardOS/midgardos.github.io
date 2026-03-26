# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUWhich.md) GNU Which | [HOME](../README.md) | Pk11-Kit [>>](./Pk11-Kit.md) |

## GTK-Doc

Name: GTK-Doc<br />
Summary: A tool to generate API documentation from formatted code comments<br />
License: GPL v2.0 or later<br />
Version: 1.36<br />
URL: [https://download.gnome.org/sources/gtk-doc/1.36/gtk-doc-1.36.0.tar.xz](https://download.gnome.org/sources/gtk-doc/1.36/gtk-doc-1.36.0.tar.xz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 14 MiB<br />

## Configuration

To configure GTK-Doc for install into the build root, run the following command:

```bash
mkdir -p build && cd build
meson setup ..            \
      --prefix=/usr       \
      --buildtype=release \
      -D tests=false
```

## Compilation and Installation

To compile GTK-Doc, run the following command:

```bash
ninja
```

The test suite cannot be tested before installation.

Finally, to install GTK-Doc into the build tree, run the following command:

```bash
ninja install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | gtkdocize, gtkdoc-check, gtkdoc-depscan, gtkdoc-fixxref, gtkdoc-mkdb, gtkdoc-mkhtml, gtkdoc-mkhtml2, gtkdoc-mkman, gtkdoc-mkpdf, gtkdoc-rebase, gtkdoc-scan, and gtkdoc-scangobj |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUWhich.md) GNU Which | [HOME](../README.md) | Pk11-Kit [>>](./Pk11-Kit.md) |

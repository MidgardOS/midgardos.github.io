# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./File_32bit.md) File 32-bit | [HOME](../README.md) | GNU Readline 32-bit [>>](./GNUReadline_32bit.md) |

## GNU Readline 64-bit

Name: GNU Readline 64-bit<br />
Summary: Libraries for command-line editing and history<br />
License: GPL v3 or later<br />
Version: 8.3<br />
URL: [https://ftp.gnu.org/gnu/readline/readline-8.3.tar.gz](https://ftp.gnu.org/gnu/readline/readline-8.3.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 17 MiB<br />

## Preparation

There are a couple of extra steps required to ensure that GNU Readline will install correctly:

```bash
sed -i '/MV.*old/d' Makefile.in
sed -i '/{OLDSUFF}/c:' support/shlib-install
sed -i 's/-Wl,-rpath,[^ ]*//' support/shobj-conf
```

## Configuration

To configure GNU Readline 64-bit for install into the build root, run the following command:

```bash
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --with-curses           \
            --docdir=/usr/share/doc/readline-8.3
```

## Compilation and Installation

To compile GNU Readline 64-bit, run the following command:

```bash
make SHLIB_LIBS="-lncursesw"
```

Finally, to install GNU Readline 64-bit into the build tree, run the following command:

```bash
make install
install -v -m644 doc/*.{ps,pdf,html,dvi} /usr/share/doc/readline-8.3
```

**NOTE: Do not delete the unpacked sources**

## Contents

| Contents ||
| --- | --- |
| Installed Programs | N/A |
| Installed Libraries | libhistory.so and libreadline.so |
| Installed Directories | /usr/include/readline, /usr/share/doc/readline-8.3 |

| Navigation |||
| --- | --- | ---: |
| [<<](./File_32bit.md) File 32-bit | [HOME](../README.md) | GNU Readline 32-bit [>>](./GNUReadline_32bit.md) |

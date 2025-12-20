# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBison.md) GNU Bison | [HOME](../README.md) | GNU Bash [>>](./GNUBash.md) |

## GNU Grep

Name: GNU Grep<br />
Summary: CLI tool for searching through text files<br />
License: GPL v3 or later<br />
Version: 3.12<br />
URL: [https://ftp.gnu.org/gnu/grep/grep-3.12.tar.xz](https://ftp.gnu.org/gnu/grep/grep-3.12.tar.xz)<br />

Average Build Time: 0.6 SBU<br />
Used Install Space: 48 MiB<br />

## Configuration

To configure GNU Grep for install into the build root, run the following commands:

```bash
sed -i "s/echo/#echo/" src/egrep.sh
./configure --prefix=/usr --libdir=/usr/lib64 --libexecdir=/usr/lib64
```

## Compilation and Installation

To compile GNU Grep, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Grep into the build tree, run the following command:

```bash
make install
install -v -d -m755 -o root -g root /etc/profile.d
install -v -d -m755 -o root -g root /usr/lib/grep
for F in "color-grep.sh" "color-ls.sh" "color-xzgrep.sh" "color-zgrep.sh"; do
    install -v -m755 -o root -g root ../system_files/etc/profile.d/$F /etc/profile.d/
done
install -v -m755 -o root -g root ../system_files/usr/lib/grep/grepconfig.sh /usr/lib/grep/
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | egrep, fgrep, grep |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBison.md) GNU Bison | [HOME](../README.md) | GNU Bash [>>](./GNUBash.md) |

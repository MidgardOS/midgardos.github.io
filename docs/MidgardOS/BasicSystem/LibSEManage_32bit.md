# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSEManage_64bit.md) LibSEManage 64-bit | [HOME](../README.md) | CheckPolicy [>>](./CheckPolicy.md) |

## LibSEManage 32-bit

Name: LibSEManage 32-bit<br />
Summary: SELinux policy management library<br />
License: LGPL v2.1<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/libsemanage-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/libsemanage-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 1.7 MiB<br />

## Configuration

The package does not have a traditional configuration script.

## Compilation and Installation

To compile LibSEManage 32-bit, run the following commands:

```bash
make clean
grep /usr/libexec . -rl | xargs sed -i "s|/usr/libexec|/usr/lib64|g"
make CC="gcc -m32 -march=i686"
```

Finally, to install LibSEManage 32-bit into the build tree, run the following commands:

```bash
PREFIX=/usr LIBDIR=/usr/lib LIBEXECDIR=/usr/lib SHLIBDIR=/usr/lib DESTDIR=$PWD/DESTDIR make install
rm -rfv DESTDIR/usr/lib/selinux
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

Other libraries and programs of the suite of tools for SELinux require linking against the static library for LibSEManage, as such it is not removed.

## Contents

See the contents section of the 64-bit build of LibSEManage for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./LibSEManage_64bit.md) LibSEManage 64-bit | [HOME](../README.md) | CheckPolicy [>>](./CheckPolicy.md) |

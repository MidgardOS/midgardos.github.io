# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./P11-Kit-64bit.md) P11-Kit 64-bit | [HOME](../README.md) | Make-CA [>>](./Make-CA.md) |

## P11-Kit 32-bit

Name: P11-Kit 32-bit<br />
Summary: Tools for loading and enumerating PKCS11 modules<br />
License: BSD 3-clause<br />
Version: 0.26.2<br />
URL: [https://github.com/p11-glue/p11-kit/releases/download/0.26.2/p11-kit-0.26.2.tar.xz](https://github.com/p11-glue/p11-kit/releases/download/0.26.2/p11-kit-0.26.2.tar.xz)<br />

Average Build Time: 0.5 SBU<br />
Used Install Space: 110 MiB<br />

## Configuration

To configure P11-Kit 32-bit for install into the build root, run the following command:

```bash
rm -rf p11-build
mkdir p11-build && cd  p11-build
PKG_CONFIG_PATH="/usr/lib/pkgconfig:/usr/share/pkgconfig"   \
CC="gcc -m32 -march=i686" CXX="g++ -m32 -march=i686"        \
meson setup ..                                              \
      --prefix=/usr                                         \
      --libdir=/usr/lib                                     \
      --libexecdir=/usr/lib                                 \
      --buildtype=release                                   \
      -D trust_paths=/etc/pki/anchors
```

## Compilation and Installation

To compile P11-Kit 32-bit, run the following command:

```bash
ninja
```

Finally, to install P11-Kit 32-bit into the build tree, run the following command:

```bash
DESTDIR=$PWD/DESTDIR ninja install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of the P11-Kit for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./P11-Kit-64bit.md) P11-Kit 64-bit | [HOME](../README.md) | Make-CA [>>](./Make-CA.md) |

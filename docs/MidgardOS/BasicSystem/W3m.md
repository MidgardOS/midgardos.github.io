# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GC_32bit.md) GC 32-bit | [HOME](../README.md) | XmlTo [>>](./XmlTo.md) |

## W3M

Name: W3M<br />
Summary: A simple console web-browser/pager<br />
License: ISC Public License<br />
Version: 0.5.3<br />
URL: [https://salsa.debian.org/debian/w3m/-/archive/v0.5.3+git20230121/w3m-v0.5.3+git20230121.tar.bz2](https://salsa.debian.org/debian/w3m/-/archive/v0.5.3+git20230121/w3m-v0.5.3+git20230121.tar.bz2)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.6 MiB<br />

## Preparation

W3M requires a couple patches to its code to build and run correctly. Apply them using the commands below:

```bash
patch -Np1 -i ../patches/w3m/0001-Fix-OOB-access-due-to-multiple-backspaces.patch
patch -Np1 -i ../patches/w3m/0001-Update-German-message-catalogue.patch
```

## Configuration

To configure W3M for install into the build root, run the following command:

```bash
find -name CVS -exec rm -Rf "{}" "+"
export CFLAGS="-DUSE_BUFINFO -DOPENSSL_NO_SSL_INTERN -D_GNU_SOURCE $(getconf LFS_CFLAGS) -fno-strict-aliasing `ncursesw6-config --cflags` -fPIE -std=gnu11"
export CXXFLAGS="$CFLAGS"
export LDFLAGS="`ncursesw6-config --libs` -pie"
./configure --prefix=/usr                       \
            --libdir=/usr/lib64                 \
            --libexecdir=/usr/lib64             \
            --mandir=/usr/share/man             \
            --docdir=/usr/share/doc/w3m-0.5.3   \
            --sysconfdir=/etc                   \
            --enable-ipv6                       \
		    --enable-alarm                      \
		    --enable-ansi-color                 \
		    --enable-digest-auth                \
		    --enable-external-uri-loader        \
		    --enable-gopher                     \
		    --enable-history                    \
		    --enable-image=fb                   \
		    --enable-keymap=lynx                \
		    --enable-m17n                       \
		    --enable-mouse                      \
		    --enable-nls                        \
		    --enable-nntp                       \
		    --enable-sslverify                  \
		    --enable-unicode                    \
		    --disable-w3mmailer
```

## Compilation and Installation

To compile W3M, run the following command:

```bash
make
```

There is no test suite to run for this package.

Finally, to install W3M into the build tree, run the following command:

```bash
make install install-helpfile
install -v -m 755 -o root -g root Bonus/*.cgi /usr/lib64/w3m/cgi-bin/
unset CFLAGS
unset CXXFLAGS
unset LDFLAGS
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | w3m, w3mman |

| Navigation |||
| --- | --- | ---: |
| [<<](./GC_32bit.md) GC 32-bit | [HOME](../README.md) | XmlTo [>>](./XmlTo.md) |

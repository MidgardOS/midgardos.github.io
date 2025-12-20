# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibGCrypt_64bit.md) GNU LibGCrypt 64-bit | [HOME](../README.md) | GNU LibKSBA 64-bit [>>](./GNULibKSBA_64bit.md) |

## GNU LibGCrypt 32-bit

Name: GNU LibGCrypt 32-bit<br />
Summary: A library implementing many cryptographic hash types<br />
License: GPL v2/LGPL v2.1<br />
Version: 1.11.2<br />
URL: [https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.11.2.tar.bz2](https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.11.2.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 6.5 MiB<br />

## Configuration

To configure GNU LibGCrypt 32-bit for install into the build root, run the following commands:

```bash
make distclean
export PUBKEYS="dsa elgamal rsa ecc"
export CIPHERS="arcfour blowfish cast5 des aes twofish serpent rfc2268 seed camellia idea salsa20 gost28147 chacha20 sm4 aria"
export DIGESTS="crc gostr3411-94 md4 md5 rmd160 sha1 sha256 sha512 sha3 tiger whirlpool stribog blake2 sm3"
export KDFS="s2k pkdf2 scrypt"
CC="gcc -m32" \
./configure --host=i686-pc-linux-gnu                    \
            --prefix=/usr                               \
            --libdir=/usr/lib                           \
            --libexecdir=/usr/lib                       \
            --docdir=/usr/share/doc/libgcrypt-1.11.2    \
            --enable-ciphers="$CIPHERS"                 \
            --enable-pubkey-ciphers="$PUBKEYS"          \
            --enable-digests="$DIGESTS"                 \
            --enable-kdfs="$KDFS"                       \
            --enable-noexecstack                        \
            --enable-random=getentropy                  \
            --enable-jent-support                       \
            --disable-static
unset PUBKEYS
unset CIPHERS
unset DIGESTS
unset KDFS
```

## Compilation and Installation

To compile GNU LibGCrypt 32-bit, run the following command:

```bash
make
```

Finally, to install GNU LibGCrypt 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
rm -fv DESTDIR/usr/lib/libgcrypt.la
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of GNU LibGCrypt for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULibGCrypt_64bit.md) GNU LibGCrypt 64-bit | [HOME](../README.md) | GNU LibKSBA 64-bit [>>](./GNULibKSBA_64bit.md) |

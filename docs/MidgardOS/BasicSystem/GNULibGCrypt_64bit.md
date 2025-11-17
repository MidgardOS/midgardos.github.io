# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./CheckPolicy.md) CheckPolicy | [HOME](../README.md) | GNU LibGCrypt 32-bit [>>](./GNULibGCrypt_32bit.md) |

## GNU LibGCrypt 64-bit

Name: GNU LibGCrypt 64-bit<br />
Summary: A library implementing many cryptographic hash types<br />
License: GPL v2/LGPL v2.1<br />
Version: 1.11.2<br />
URL: [https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.11.2.tar.bz2](https://gnupg.org/ftp/gcrypt/libgcrypt/libgcrypt-1.11.2.tar.bz2)<br />

Average Build Time: 0.3 SBU<br />
Used Install Space: 6.5 MiB<br />

## Configuration

To configure GNU LibGCrypt 64-bit for install into the build root, run the following commands:

```bash
export PUBKEYS="dsa elgamal rsa ecc"
export CIPHERS="arcfour blowfish cast5 des aes twofish serpent rfc2268 seed camellia idea salsa20 gost28147 chacha20 sm4 aria"
export DIGESTS="crc gostr3411-94 md4 md5 rmd160 sha1 sha256 sha512 sha3 tiger whirlpool stribog blake2 sm3"
export KDFS="s2k pkdf2 scrypt"
./configure --prefix=/usr                               \
            --libdir=/usr/lib64                         \
            --libexecdir=/usr/lib64                     \
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

To compile GNU LibGCrypt 64-bit, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU LibGCrypt 64-bit into the build tree, run the following commands:

```bash
make install
rm -fv /usr/lib64/libgcrypt.la
mkdir -pv -m 0755 /etc/gcrypt
install -m 644 ../system_files/etc/gcrypt/random.conf /etc/gcrypt/random.conf
install -m 644 ../system_files/etc/gcrypt/hwf.deny /etc/gcrypt/hwf.deny
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | dumpsexp, hmac256, libgcrypt-config, and mpicalc |
| Installed Libraries | libgcrypt.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./CheckPolicy.md) CheckPolicy | [HOME](../README.md) | GNU LibGCrypt 32-bit [>>](./GNULibGCrypt_32bit.md) |

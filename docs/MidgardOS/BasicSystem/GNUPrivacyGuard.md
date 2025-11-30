# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNPth_32bit.md) GNU nPth 32-bit | [HOME](../README.md) | GNU PinEntry [>>](./GNUPinEntry.md) |

## GNU Privacy Guard

Name: GNU Privacy Guard<br />
Summary: An OpenPGP implementation<br />
License: GPL v2 or later/LGPL v2.1 or later/LGPL v3 or later<br />
Version: 2.5.14<br />
URL: [https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.5.14.tar.bz2](https://gnupg.org/ftp/gcrypt/gnupg/gnupg-2.5.14.tar.bz2)<br />

Average Build Time: 1 SBU<br />
Used Install Space: 36 MiB<br />

## Configuration

To configure GNU Privacy Guard for install into the build root, run the following command:

```bash
./configure --prefix=/usr                                           \
            --libdir=/usr/lib64                                     \
            --libexecdir=/usr/lib64                                 \
            --docdir=/usr/share/doc/gnupg-2.5.14                    \
            --disable-static                                        \
            --enable-g13                                            \
            --disable-rpath                                         \
            --enable-large-secmem                                   \
            --with-gnu-ld                                           \
            --with-default-trust-store-file=/etc/ssl/ca-bundle.pem  \
            --enable-build-timestamp

```

**NOTE: The following features are disable for now due to missing dependencies:<br />**
**- LDAP support<br />**
**- TPM support<br />**

These features will be enabled later when this package is rebuilt as an RPM.

## Compilation and Installation

To compile GNU Privacy Guard, run the following command:

```bash
make
```

Next, run the test suite:

```bash
make check
```

Finally, to install GNU Privacy Guard into the build tree, run the following command:

```bash
make install
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | dirmngr, dirmngr-client, g13, gpg, gpg-agent, gpg-auth, gpg-authcode-sign.sh, gpg-card,  gpg-check-pattern, gpg-connect-agent, gpg-mail-tube, gpg-pair-tool, gpg-preset-passphrase, gpg-protect-tool, gpg-wks-client, gpg-wks-server, gpgconf, gpgparsemail, gpgscm, gpgsm, gpgsplit, gpgtar, gpgv, kbxutil, keyboxd, scdaemon, and watchgnupg |
| Installed Plugins | addgnupghome, applygnupgdefaults, and g13-syshelp |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUNPth_32bit.md) GNU nPth 32-bit | [HOME](../README.md) | GNU PinEntry [>>](./GNUPinEntry.md) |

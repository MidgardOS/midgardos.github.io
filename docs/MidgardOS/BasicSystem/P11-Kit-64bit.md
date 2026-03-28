# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GTK-Doc.md) GTK-Doc | [HOME](../README.md) | P11-Kit 32-bit [>>](./P11-Kit-32bit.md) |

## P11-Kit 64-bit

Name: P11-Kit 64-bit<br />
Summary: Tools for loading and enumerating PKCS11 modules<br />
License: BSD 3-clause<br />
Version: 0.26.2<br />
URL: [https://github.com/p11-glue/p11-kit/releases/download/0.26.2/p11-kit-0.26.2.tar.xz](https://github.com/p11-glue/p11-kit/releases/download/0.26.2/p11-kit-0.26.2.tar.xz)<br />

Average Build Time: 0.5 SBU<br />
Used Install Space: 110 MiB<br />

## Preparation

To prepare P11-Kit 64-bit for build, create the distribution specific anchor hook:

```bash
sed '20,$ d' -i trust/trust-extract-compat
cat >> trust/trust-extract-compat << "EOF"
# Copy existing anchor modifications to /etc/ssl/local
/usr/libexec/make-ca/copy-trust-modifications

# Update trust stores
/usr/sbin/make-ca -r
EOF
```

## Configuration

To configure P11-Kit 64-bit for install into the build root, run the following command:

```bash
mkdir p11-build && cd  p11-build
meson setup ..                  \
      --prefix=/usr             \
      --libdir=/usr/lib64       \
      --libexecdir=/usr/lib64   \
      --buildtype=release       \
      -D trust_paths=/etc/pki/anchors
```

## Compilation and Installation

To compile P11-Kit 64-bit, run the following command:

```bash
ninja
```

Next, run the test suite:

```bash
ninja test
```

Finally, to install P11-Kit 64-bit into the build tree, run the following command:

```bash
ninja install
ln -sfv /usr/lib64/p11-kit/trust-extract-compat /usr/bin/update-ca-certificates
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | p11-kit, trust, and update-ca-certificates |
| Installed Libraries | libp11-kit.so and p11-kit-proxy.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./GTK-Doc.md) GTK-Doc | [HOME](../README.md) | P11-Kit 32-bit [>>](./P11-Kit-32bit.md) |

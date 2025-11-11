# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./libsepol_32bit.md) LibSEPol 32-bit | [HOME](../README.md) | LibSELinux 32-bit [>>](./libselinux_32bit.md) |

## LibSELinux 64-bit

Name: LibSELinux 64-bit<br />
Summary: SELinux runtime and utilities<br />
License: Public Domain<br />
Version: 3.9<br />
URL: [https://github.com/SELinuxProject/selinux/releases/download/3.9/libselinux-3.9.tar.gz](https://github.com/SELinuxProject/selinux/releases/download/3.9/libselinux-3.9.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 2.7 MiB<br />

## Configuration

This package has no traditional configuration script.

## Compilation and Installation

To compile LibSELinux 64-bit, run the following command:

```bash
make
```

This package doesn't have a test suite.

Finally, to install LibSELinux 64-bit into the build tree, run the following command:

```bash
PREFIX=/usr LIBDIR=/usr/lib64 SHLIBDIR=/usr/lib64 make install
install -v -d -m 755 -o root -g root /etc/selinux
install -v -m 644 -o root -g root ../system_files/etc/selinux/config /etc/selinux/
install -v -m 755 -o root -g root ../system_files/usr/sbin/selinux-ready /usr/sbin/
```

As before with LibSEPol, the static library from LibSELinux will be linked into other libraries and programs in the SELinux suite of tools and libraries. As such it should not be removed.

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | avcstat, compute_av, compute_create, compute_member, compute_relabel, getconlist, getdefaultcon, getenforce, getfilecon, getpidcon, getpidprevcon, getpolicyload, getsebool, getseuser, matchpathcon, policyvers, sefcontext_compile, selabel_digest, selabel_get_digests_all_partial_matches, selabel_lookup_best_match, selinux_check_securetty_context, selabel_compare, selabel_lookup, selabel_partial_match, selinux_check_access, selinuxenabled, selinuxexeccon, setenforce, setfilecon, togglesebool, and validatetrans |
| Installed Libraries | libselinux.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./libsepol_32bit.md) LibSEPol 32-bit | [HOME](../README.md) | LibSELinux 32-bit [>>](./libselinux_32bit.md) |

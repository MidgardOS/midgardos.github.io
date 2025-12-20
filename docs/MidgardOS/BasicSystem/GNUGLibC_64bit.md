# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./IanaEtcFiles.md) IANA Etc Files | [HOME](../README.md) | GNU GLibC 32-bit [>>](./GNUGLibC_32bit.md) |

## GNU GLibC 64-bit

Name: GNU GLibC 64-bit<br />
Summary: The GNU C language runtime library - 64-bit<br />
License: GPL v2.0+/LGPL 2.1+<br />
Version: 2.42<br />
URL: [https://ftp.gnu.org/gnu/glibc/glibc-2.42.tar.xz](https://ftp.gnu.org/gnu/glibc/glibc-2.42.tar.xz)<br />

Average Build Time: 6 SBU<br />
Used Install Space: 2.2 GiB<br />

## Configuration

To configure GNU GLibC 64-bit for install into the build root, run the following command:

```bash
patch -Np1 -i ../glibc-2.42-fhs-1.patch
# a small fix to avoid causing Valgrind to crash
sed -e '/unistd.h/i #include <string.h>'           \
    -e '/libc_rwlock_init/c\ __libc_rwlock_define_initialized (, reset_lock); memcpy (&lock, &reset_lock, sizeof (lock));' \
    -i stdlib/abort.c
mkdir -v build && cd build
echo "rootsbindir=/usr/sbin" > configparms
../configure --prefix=/usr                                                  \
             --libdir=/usr/lib64                                            \
             --libexecdir=/usr/lib64                                        \
             --enable-bind-now                                              \
             --enable-multi-arch                                            \
             --enable-cet                                                   \
             --enable-stackguard-randomization                              \
             --enable-tunables                                              \
             --with-bugurl="https://github.com/MidgardOS/MidgardOS/Issues"  \
             --disable-werror                                               \
             libc_cv_slibdir=/usr/lib64                                     \
             --enable-kernel=5.4
```

## Compilation and Installation

To compile GNU GLibC 64-bit, run the following command:

```bash
make
touch /etc/ld.so.conf
install -d -m 755 -o root -g root /etc/ld.so.conf.d
```

Next, run the code tests:

```bash
make check
```

There are a couple tests that may fail for various reasons. The ones listed below are safe to ignore:
- io/tst-lchmod is known to fail in the chroot environment.
- misc/tst-preadvwritev2 and misc/tst-preadvwritev64v2 are known to fail if the host kernel is Linux-6.14 or later.
- Some tests, for example nss/tst-nss-files-hosts-multi and nptl/tst-thread-affinity* are known to fail due to a timeout (especially when the system is relatively slow and/or running the test suite with multiple parallel make jobs). These tests can be identified with:
```
grep "Timed out" $(find -name \*.out)
```
It's possible to re-run a single test with enlarged timeout with TIMEOUTFACTOR=<factor> make test t=<test name>. For example, TIMEOUTFACTOR=10 make test t=nss/tst-nss-files-hosts-multi will re-run nss/tst-nss-files-hosts-multi with ten times the original timeout.

Finally, to install GNU GLibC 64-bit into the build tree, run the following command:

```bash
sed '/test-installation/s@$(PERL)@echo not running@' -i ../Makefile
make install
sed '/RTLDLIST=/s@/usr@@g' -i /usr/bin/ldd
localedef -i C -f UTF-8 C.UTF-8
localedef -i cs_CZ -f UTF-8 cs_CZ.UTF-8
localedef -i de_DE -f ISO-8859-1 de_DE
localedef -i de_DE@euro -f ISO-8859-15 de_DE@euro
localedef -i de_DE -f UTF-8 de_DE.UTF-8
localedef -i el_GR -f ISO-8859-7 el_GR
localedef -i en_GB -f ISO-8859-1 en_GB
localedef -i en_GB -f UTF-8 en_GB.UTF-8
localedef -i en_HK -f ISO-8859-1 en_HK
localedef -i en_PH -f ISO-8859-1 en_PH
localedef -i en_US -f ISO-8859-1 en_US
localedef -i en_US -f UTF-8 en_US.UTF-8
localedef -i es_ES -f ISO-8859-15 es_ES@euro
localedef -i es_MX -f ISO-8859-1 es_MX
localedef -i fa_IR -f UTF-8 fa_IR
localedef -i fr_FR -f ISO-8859-1 fr_FR
localedef -i fr_FR@euro -f ISO-8859-15 fr_FR@euro
localedef -i fr_FR -f UTF-8 fr_FR.UTF-8
localedef -i is_IS -f ISO-8859-1 is_IS
localedef -i is_IS -f UTF-8 is_IS.UTF-8
```

Now that GLibC and the locales are installed, there are a few extra steps needed to configure its functionality for use on the system. First, configure the Name Service Switch so that the host resolver and account management functions can be configured to look up data correctly:

```bash
cat > /etc/nsswitch.conf << "EOF"
# In order of likelihood of use to accelerate lookup.
passwd:     files systemd
shadow:     files systemd
group:      files systemd
hosts:      files mymachines resolve [!UNAVAIL=return] files myhostname dns
services:   files
netgroup:   files
sudoers:    files
automount:  files
subid:      files
aliases:    files
ethers:     files
gshadow:    files systemd
networks:   files
protocols:  files
publickey:  files
rpc:        files
EOF
```

Once the `/etc/nsswitch.conf` is installed, timezone data needs setup. To do so, run the following commands:

```bash
tar -xf ../../tzdata2025b.tar.gz

ZONEINFO=/usr/share/zoneinfo
mkdir -pv $ZONEINFO/{posix,right}

for tz in etcetera southamerica northamerica europe africa antarctica  \
          asia australasia backward; do
    zic -L /dev/null   -d $ZONEINFO       ${tz}
    zic -L /dev/null   -d $ZONEINFO/posix ${tz}
    zic -L leapseconds -d $ZONEINFO/right ${tz}
done

cp -v zone.tab zone1970.tab iso3166.tab $ZONEINFO
zic -d $ZONEINFO -p America/New_York
unset ZONEINFO tz
tzselect
```

After running `tzselect` and answering a couple questions, it will output the name of the time zone appropriate for your system. To enable this, run the following command subsituting the time zone for the `$TZ` variable:

```bash
ln -sfv /usr/share/zoneinfo/$TZ /etc/localtime
```

Now, configure the dynamic library loader by telling it to read configurations in `/etc/ld.so.conf.d`:

```bash
cat >> /etc/ld.so.conf << "EOF"
# Add an include directory
include /etc/ld.so.conf.d/*.conf

EOF
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | gencat, getconf, getent, iconv, iconvconfig, ldconfig, ldd, lddlibc4, ld.so (symlink to ld-linux-x86-64.so.2 or ld-linux.so.2), locale, localedef, makedb, mtrace, pcprofiledump, pldd, sln, sotruss, sprof, tzselect, xtrace, zdump, and zic |
| Installed Libraries | ld-linux-x86-64.so.2, ld-linux.so.2, libBrokenLocale.{a,so}, libanl.{a,so}, libc.{a,so}, libc_nonshared.a, libc_malloc_debug.so, libdl.{a,so.2}, libg.a, libm.{a,so}, libmcheck.a, libmemusage.so, libmvec.{a,so}, libnsl.so.1, libnss_compat.so, libnss_dns.so, libnss_files.so, libnss_hesiod.so, libpcprofile.so, libpthread.{a,so.0}, libresolv.{a,so}, librt.{a,so.1}, libthread_db.so, and libutil.{a,so.1} |
| Installed Directories | /usr/include/arpa, /usr/include/bits, /usr/include/gnu, /usr/include/net, /usr/include/netash, /usr/include/netatalk, /usr/include/netax25, /usr/include/neteconet, /usr/include/netinet, /usr/include/netipx, /usr/include/netiucv, /usr/include/netpacket, /usr/include/netrom, /usr/include/netrose, /usr/include/nfs, /usr/include/protocols, /usr/include/rpc, /usr/include/sys, /usr/lib/audit, /usr/lib/gconv, /usr/lib/locale, /usr/libexec/getconf, /usr/share/i18n, /usr/share/zoneinfo, and /var/lib/nss_db |

| Navigation |||
| --- | --- | ---: |
| [<<](./IanaEtcFiles.md) IANA Etc Files | [HOME](../README.md) | GNU GLibC 32-bit [>>](./GNUGLibC_32bit.md) |

# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBison.md) GNU Bison | [HOME](../README.md) | Python 3.13 [>>](./Python313.md) |

## Perl

Name: Perl<br />
Summary: The Perl scripting language<br />
License: Perl/Artistic/GPLv2<br />
Version: 5.42.0<br />
URL: [https://www.cpan.org/src/5.0/perl-5.42.0.tar.gz](https://www.cpan.org/src/5.0/perl-5.42.0.tar.gz)<br />

## Configuration

To configure Perl for install into the build root, run the following command:

```bash
sh Configure -des                                                                        \
             -D prefix=/usr                                                              \
             -D vendorprefix=/usr                                                        \
             -D useithreads                                                              \
             -D uselongdouble                                                            \
             -D usemultiplicity                                                          \
             -D useposix                                                                 \
             -D useshrplib                                                               \
             -D use64bitint                                                              \
             -D use64bitall                                                              \
             -D privlib=/usr/lib64/perl5/5.42.0                                          \
             -D archlib=/usr/lib64/perl5/5.42.0/x86_64-linux-thread-multi                \
             -D sitelib=/usr/lib64/perl5/site_perl/5.42.0                                \
             -D sitearch=/usr/lib64/perl5/site_perl/5.42.0/x86_64-linux-thread-multi     \
             -D vendorlib=/usr/lib64/perl5/vendor_perl/5.42.0                            \
             -D vendorarch=/usr/lib64/perl5/vendor_perl/5.42.0/x86_64-linux-thread-multi
```

## Compilation and Installation

To compile Perl, run the following command:

```bash
make
```

Finally, to install Perl into the build tree, run the following command:

```bash
make install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUBison.md) GNU Bison | [HOME](../README.md) | Python 3.13 [>>](./Python313.md) |

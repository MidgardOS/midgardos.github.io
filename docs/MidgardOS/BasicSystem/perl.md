# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULess.md) GNU Less | [HOME](../README.md) | perl-XML-Parser [>>](./perl-xml-parser.md) |

## Perl

Name: Perl<br />
Summary: The Practical Extraction and Report Language<br />
License: Perl/Artistic/GPLv2<br />
Version: 5.42.0<br />
URL: [https://www.cpan.org/src/5.0/perl-5.42.0.tar.gz](https://www.cpan.org/src/5.0/perl-5.42.0.tar.gz)<br />

Average Build Time: 1.3 SBU<br />
Used Install Space: 81 MiB<br />

## Configuration

To configure Perl for install into the build root, run the following commands:

```bash
export BUILD_ZLIB=False
export BUILD_BZIP2=0
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

Next, run the test suite:

```bash
TEST_JOBS=$(nproc) make test_harness
```

There are four tests that fail currently.

Finally, to install Perl into the build tree, run the following commands:

```bash
make install
unset BUILD_ZLIB BUILD_BZIP2
```

## Contents

| Contents | |
| --- | --- |
| Installed Programs | corelist cpan enc2xs encguess h2ph h2xs instmodsh json_pp libnetcfg perl perl5.42.0 perlbug perldoc perlivp perlthanks piconv pl2pm pod2html pod2man pod2text pod2usage podchecker podselect prove ptar ptardiff ptargrep shasum splain xsubpp zipdetails |
| Installed Libraries | Numerous libraries and Perl modules |

| Navigation |||
| --- | --- | ---: |
| [<<](./GNULess.md) GNU Less | [HOME](../README.md) | perl-XML-Parser [>>](./perl-xml-parser.md) |

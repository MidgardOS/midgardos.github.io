# Section 5 - Building the Basic System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./SQLite3_64bit.md) SQLite3 64-bit | [HOME](../README.md) | Python 3.13 [>>](./Python313.md) |

## SQLite3 32-bit

Name: SQLite3 32-bit<br />
Summary: An embedded, serverless transactional database library<br />
License: SQLite3 Public Domain License<br />
Version: 3.51.0<br />
URL: [https://sqlite.org/2025/sqlite-autoconf-3510000.tar.gz](https://sqlite.org/2025/sqlite-autoconf-3510000.tar.gz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 1.8 MiB<br />

## Configuration

To configure SQLite3 32-bit for install into the build root, run the following commands:

```bash
make distclean
export OPTFLAGS="-fmessage-length=0 -D_FORTIFY_SOURCE=2 -fstack-protector -funwind-tables -fasynchronous-unwind-tables"
export CFLAGS="${OPTFLAGS} \
	-DSQLITE_ENABLE_API_ARMOR \
	-DSQLITE_ENABLE_COLUMN_METADATA \
	-DSQLITE_ENABLE_DBSTAT_VTAB \
	-DSQLITE_ENABLE_HIDDEN_COLUMNS \
	-DSQLITE_ENABLE_FTS3 \
	-DSQLITE_ENABLE_FTS4 \
	-DSQLITE_ENABLE_FTS5 \
	-DSQLITE_ENABLE_JSON1 \
	-DSQLITE_ENABLE_RBU \
	-DSQLITE_ENABLE_RTREE \
	-DSQLITE_ENABLE_UPDATE_DELETE_LIMIT \
	-DSQLITE_SOUNDEX \
	-DSQLITE_ENABLE_UNLOCK_NOTIFY \
	-DSQLITE_SECURE_DELETE \
	-DSQLITE_ENABLE_MATH_FUNCTIONS \
	-DSQLITE_STRICT_SUBTYPE=1"
CC="gcc -m32" CXX="g++ -m32"        \
./configure --host=i686-pc-linux-gnu \
            --prefix=/usr            \
            --libdir=/usr/lib        \
            --libexecdir=/usr/lib    \
            --disable-static         \
            --soname=legacy          \
            --enable-readline        \
            --enable-fts3            \
            --enable-fts4            \
            --enable-fts5            \
            --enable-update-limit    \
            --enable-rtree           \
            --enable-session
```

## Compilation and Installation

To compile SQLite3 32-bit, run the following command:

```bash
make LDFLAGS.rpath=""
```

Finally, to install SQLite3 32-bit into the build tree, run the following commands:

```bash
make DESTDIR=$PWD/DESTDIR install
cp -Rv DESTDIR/usr/lib/* /usr/lib/
```

## Contents

See the contents section of the 64-bit build of SQLite3 for details.

| Navigation |||
| --- | --- | ---: |
| [<<](./SQLite3_64bit.md) SQLite3 64-bit | [HOME](../README.md) | Python 3.13 [>>](./Python313.md) |

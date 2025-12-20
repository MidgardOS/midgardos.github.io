# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./7zip.md) 7zip CLI | [HOME](../README.md) | SQLite3 32-bit [>>](./SQLite3_32bit.md) |

## SQLite3 64-bit

Name: SQLite3 64-bit<br />
Summary: An embedded, serverless transactional database library<br />
License: SQLite3 Public Domain License<br />
Version: 3.51.0<br />
URL: [https://sqlite.org/2025/sqlite-autoconf-3510000.tar.gz](https://sqlite.org/2025/sqlite-autoconf-3510000.tar.gz)<br />

Average Build Time: 0.4 SBU<br />
Used Install Space: 43 MiB<br />

## Configuration

To configure SQLite3 64-bit for install into the build root, run the following commands:

```bash
7zz x ../sqlite-doc-3510000.zip
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
./configure --prefix=/usr           \
            --libdir=/usr/lib64     \
            --libexecdir=/usr/lib64 \
            --disable-static        \
            --soname=legacy         \
            --enable-readline       \
            --enable-fts3           \
            --enable-fts4           \
            --enable-fts5           \
            --enable-update-limit   \
            --enable-rtree          \
            --enable-session
```

## Compilation and Installation

To compile SQLite3 64-bit, run the following command:

```bash
make LDFLAGS.rpath=""
```

Finally, to install SQLite3 64-bit into the build tree, run the following commands:

```bash
make install
install -v -m755 -d /usr/share/doc/sqlite-3.51.0
cp -Rv sqlite-doc-3510000/* /usr/share/doc/sqlite-3.51.0/
unset CFLAGS
unset OPTFLAGS
```

**NOTE: Do not delete the unpacked sources after build.**

## Contents

| Contents | |
| --- | --- |
| Installed Programs | sqlite3 |
| Installed Libraries | libsqlite3.so |

| Navigation |||
| --- | --- | ---: |
| [<<](./7zip.md) 7zip CLI | [HOME](../README.md) | SQLite3 32-bit [>>](./SQLite3_32bit.md) |

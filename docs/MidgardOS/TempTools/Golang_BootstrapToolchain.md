# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection - Pass 2 | [HOME](../README.md) | GetURL [>>](./GetURL.md) |

## Golang Bootstrap Toolchain

Name: Golang Bootstrap Toolchain<br />
Summary: A binary bootstrap for Golang, a compiled, GC enabled, high-concurrency language for systems and cloud tasks<br />
License: BSD 3-clause<br />
Version: 1.26.1<br />
URL: [https://go.dev/dl/go1.26.1.linux-amd64.tar.gz](https://go.dev/dl/go1.26.1.linux-amd64.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 264 MiB<br />

## Configuration

This package is a temporary bootstrap required for building Golang later. As such, there is no configuration script to run in this phase.

## Compilation and Installation

As stated above, this package is a temporary binary toolchain. It is required since Golang itself cannot be built without Golang already being present on the system, as the Golang compiler and tools are written in Golang. The final build of Golang will occur in the Basic System Tools section.

Unlike other packages, this temporary toolchain will be installed in /usr/local/go to ensure that the temporary tools can easily be removed later.

To install the bootstrap toolchain, run the following commands:

```bash
cd /sources
tar xvf go1.26.1.tar.gz
cp -Rv go /usr/local/
chown -Rv root:root /usr/local/go
ln -sv /usr/local/go/bin/go /usr/local/bin/
ln -sv /usr/local/go/bin/gofmt /usr/local/bin/
```

## Validating Installation

To ensure that Golang is working correctly, run the following commands:

```bash
cat << EOF > /tmp/main.go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
EOF
go run /tmp/main.go
echo $?
go version
```

This should output the following:

```
hello, world
0
go version 1.26.1 linux/amd64
```

More details about this package is covered later in the core system build.


| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection - Pass 2 | [HOME](../README.md) | GetURL [>>](./GetURL.md) |

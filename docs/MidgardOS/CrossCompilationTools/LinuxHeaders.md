| Navigation |||
| --- | --- | ---: |
| [<<](./File.md) | [HOME](../README.md) | [>>](./GNUM4.md) |

# Linux System Headers

Name: linux-headers<br />
Summary: System headers meant for interfacing with the Linux Kernel<br />
License: GPL v2<br />
Version: 6.1.2<br />
URL: [https://cdn.kernel.org/pub/linux/kernel/v6.x/](https://cdn.kernel.org/pub/linux/kernel/v6.x/)

## Configuration

There is no configuration step for this package.

## Compilation and Installation

The Linux Headers are a sanitized version of the C header files for the Linux Kernel, allowing software to interface
with it's syscall API. To install them, run the following commands:

```bash
make mrproper
make ARCH=x86_64 headers_check
make ARCH=x86_64 INSTALL_HDR_PATH=/tools headers_install
```

More details about this package is covered later in the core system build.

| Navigation |||
| --- | --- | ---: |
| [<<](./File.md) | [HOME](../README.md) | [>>](./GNUM4.md) |

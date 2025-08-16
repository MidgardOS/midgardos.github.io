# Section 2 - Cross-Compilation Toolchain

| Navigation |||
| --- | --- | ---: |
| [<<](../IgnoringPreFinalSWTests.md) Ignoring Some Warnings | [HOME](../README.md) | Linux Kernel C Headers [>>](./LinuxHeaders.md) |

# Building the Cross-Compilation Toolchain

This section of the how-to walks through the process of building the cross-compilation toolchain that will be used to
build the core operating system tools.

## Common Notes

While some packages built will require some patching, this will not compare to when the core components are built as
RPM packages.

When building packages, the archive of the source code for the package should be unpacked and entered when issuing the
commands that will be documented for the build of the various components.

Also, after installation, the unpacked sources and build directories should be removed, unless explicitly instructed
otherwise to avoid misconfiguration or errors during the build later when rebuilt using RPM.

| Navigation |||
| --- | --- | ---: |
| [<<](../IgnoringPreFinalSWTests.md) Ignoring Some Warnings | [HOME](../README.md) | Linux Kernel C Headers [>>](./LinuxHeaders.md) |

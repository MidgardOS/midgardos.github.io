| Navigation |||
| --- | --- | ---: |
| [<<](./HostRequirements.md) Host Requirements | [HOME](./README.md) | Preparation Overview [>>](./PrepOverview.md) |

# Structure

This documentation is split into a few different sections:

## Section 1 - Preparation

This section covers the preparation of the disks, user creation, and configuration of various environment variables to build the components.

## Section 2 - Building the Cross-compiler Toolchain

This section covers building the cross-compiler toolchain to avoid having the new build root depend on any of the libraries or tools from the host system.

## Section 3 - Building Temporary Tools

This section covers building essential temporary tools to further isolate the new build root from the host's operating system.

## Section 4 - Entering the Chroot and Building More Temporary Tools

This section covers entering the build root via an isolated chroot, and finishing the build of some additional temporary tools.

## Section 5 - Building the Base System Tools

This section covers building the native tools for allowing the core MidgardOS system to be built.

## Section 6 - Building the Core MidgardOS System

This section covers building the actual core MidgardOS system inside its isolated environment on the host computer, including configuring the system to boot and installing the system's kernel and bootloader.

## Section 7 - Rebuilding the Core Using RPM

This section covers rebuilding the core MidgardOS system from within itself using RPM and a custom build daemon for managing dependencies. This will be performed inside a live MidgardOS install booted to the system's kernel.

If any issues come up, please post an issue in our GitHub repository for this documentation, [here](https://github.com/MidgardOS/midgardos.github.io/issues/new/choose).

| Navigation |||
| --- | --- | ---: |
| [<<](./HostRequirements.md) Host Requirements | [HOME](./README.md) | Preparation Overview [>>](./PrepOverview.md) |

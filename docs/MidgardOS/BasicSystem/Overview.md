# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](../ChrootTempTools/CleanupAndArchival.md) Clean-up and Archival | [HOME](../README.md) | Man Pages [>>](./ManPages.md) |

## Overview of Building the Basic OS

Now that the temporary tools have all been built and installed allowing the tools to be run from within an isolated chroot, the work of building the basic operating system components can begin.

While the package manager (RPM) and the dependency solver (Apt-RPM) are not yet installed, the core components to allow MidgardOS to boot can be built and installed leading to RPM and Apt-RPM to be installed into the system.

In this part of the guide, each section about a particular package will contain additional info about the package and average build times. Also note, until RPM is in use, builds will use the normal GCC `-O2` optimization flag, and avoid `-march` and `-mtune` flags to ensure build failures do not occur due to over-optimizations that can break some software during compilation.

## Static Libraries

Unless noted, most packages will be installed without static libraries during this phase to avoid extra work later.

| Navigation |||
| --- | --- | ---: |
| [<<](../ChrootTempTools/CleanupAndArchival.md) Clean-up and Archival | [HOME](../README.md) | Man Pages [>>](./ManPages.md) |

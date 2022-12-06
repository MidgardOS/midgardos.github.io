# MidgardOS

This is the website that documents the build and design of MidgardOS.

## Features

MidgardOS is a Linux-based server operating system that adheres to the following principles:

1. Reliable
1. Maintainable
1. Observational

### Reliability

The Reliability principle means that the OS will, within reason, ship multiple versions of libraries when newer ones break backwards compatibility to allow older applications to work.

This does not mean that the OS will allow multiple development versions to be installed at once, but will install the versioned libraries to allow compatibility issues to be resolved by installing the appropriate version of the library. Note that this is not retroactive to before MidgardOS is released. Note that there are a few corner cases where this may require using containerization instead due to specific library versioning issues.

### Maintainability

The Maintainability principle requires that the operating system will have automation tooling built in to allow operations and DevOps engineers to manage and maintain installs of the operating system at scale. By default, MidgardOS ships with CinC (CinC is not Chef) baked into the system, and the operating system's own configuration tools will set settings that the locally installed CinC service will apply and enforce.

### Observability

The Observability principle means that after installation, the default install will be able to be easly monitored, both for performance and operational health

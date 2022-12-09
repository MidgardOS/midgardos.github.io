# MidgardOS

This is the website that documents the build and design of MidgardOS.

### Table of Contents
1. [Principles](#principles)
1. [Features](#features)
1. [Development](#development)

## Principles

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

## Features

MidgardOS will have the following features at first release:

- RPM-based
- CinC built-in
  - OS configuration tools will leverage CinC via overrides that the chef code will read to set OS settings
- Adhere to the merged `/` -> `/usr` configurration in most Linux distributions available now
- Be designed for modern server hardware
  - This means that we will target UEFI bootloading only
- Utilitize modern system components
  - use all components that ship with systemd
    - Instead of GRUB -> systemd-boot
    - Instead of NetConf/Wicked/NetworkManager -> systemd-networkd
    - Instead of using the GLibC resolver stub -> systemd-resolved
    - Instead of chrony for system time sync -> systemd-timesyncd
    - Instead of using cron for system tasks, systemd timers will be used
    - Use `systemd-coredump` for managing application coredumps
  - To simplify the orchestration of firewall rules, use UFW instead of firewalld or raw iptables/nftables
  - FluentD in the installation pre-configured to allow exporting systemd-journald logs

## Development

[Building the Base OS](./MidgardOS/)

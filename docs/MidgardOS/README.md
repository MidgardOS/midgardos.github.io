# MidgardOS - Core OS Build Documentation

Right now, MidgardOS is manually built. Eventually, automation will be put in place to create container images that will be used to drive builds of RPMs for the OS. For now, this documentation, heavily inspired by the Cross-LFS and LFS project, will have to suffice.

## Table of Contents

### Preface

1. [Pre-requisites](./Prerequisites.md)
1. [Host System Requirements](./HostRequirements.md)
1. [Structure and Errata](./Structure.md)

### Preparing to Build the Core OS Build Root

1. Overview
1. Creating a New Partition
1. Create the Required Mount Points and Symbolic Links
1. Mounting the Newly Created Filesystem
1. Creating a Build User
1. Setting up Environment variables
1. 

## Acknowledgements

The MidgardOS development team would like to acknowledge the teams and projects that have helped shape this documentation:

- The CLFS project leadership:
  - William Harrington - Lead Developer
  - Jonathan Norman - x86, x86_64, PowerPC & UltraSPARC builds, Release Manager 2.x Series
  - Chris Staub - x86 and x86_64 builds. Leader of Quality Control
- The CLFS Team:
  - Matt Darcy - x86, X86_64, and Sparc builds
  - Manuel Canales Esparcia - Book XML
  - Justin Knierim - Website Architect
  - Ken Moffat - PowerPC and X86_64 builds. Developer of Pure 64 Hint
- Outside the Development Team for the CLFS project:
  - Jürg Billeter - Testing and assisting in the development of the Linux Headers Package
  - Richard Downing - Testing, typo, and content fixes
  - Peter Ennis - Typo and content fixes
  - Tony Morgan - Typo and content fixes
  - G. Moko - Text updates and Typos
  - Maxim Osipov - MIPS Testing
  - Doug Ronne - Various x86_64 fixes
  - Theo Schneider - Testing of the Linux Headers Package
  - Martin Ward - Recommendations for Systemd and the Boot method, among other contributions
  - William Zhou - Text updates and Typos
- Former CLFS Team Members:
  - Joe Ciccone - Lead Developer
  - Nathan Coulson - Bootscripts
  - Jim Gifford - Lead Developer
  - Jeremy Huntwork - PowerPC, x86, Sparc builds
  - Karen McGuiness - Proofreader
  - Ryan Oliver - Build Process Developer
  - Alexander E. Patrakov - Udev/Hotplug Integration
  - Jeremy Utley - Release Manager 1.x Series
  - Zack Winkles - Unstable book work
- The Linux From Scratch Project:
  - Gerard Beekmans – Creator of Linux From Scratch, on which Cross-LFS is based, which inspired the work on MidgardOS' build root development

| Navigation |||
| --- | --- | ---: |
| | [HOME](./README.md) | [>>](./Prerequisites.md) |

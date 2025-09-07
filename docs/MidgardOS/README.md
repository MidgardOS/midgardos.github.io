# MidgardOS - Core OS Build Documentation

| Navigation |||
| --- | --- | ---: |
| | [HOME](./README.md) | Prerequisites [>>](./Prerequisites.md) |

Right now, MidgardOS is manually built. Eventually, automation will be put in place to create container images that will be used to drive builds of RPMs for the OS. For now, this documentation, heavily inspired by the Cross-LFS and LFS project, will have to suffice.

## Acknowledgements

[Acknowledgements](./Acknowledgements.md)

## Table of Contents

### Preface

1. [Pre-requisites](./Prerequisites.md)
1. [Host System Requirements](./HostRequirements.md)
1. [Structure and Errata](./Structure.md)

### Preparing to Build the Core OS Build Root

1. [Overview](./PrepOverview.md)
1. [Creating the New Partitions](./CreatePartition.md)
1. [Format the New Partitions, Create Mount Points, and Mount the New Filesystems](./ManageDisk.md)
1. [Create OS Directories and Symbolic Links](./DirectoriesAndSymlinks.md)
1. [Creating a Build User](./CreateBuildUser.md)
1. [Setting up Environment variables](./SetEnvironmentVars.md)
1. [Ignoring Software Test Suites](./IgnoringPreFinalSWTests.md)

### Building the Cross-compiler Toolchain

1. [Overview](./CrossCompilationTools/Overview.md)
1. [Linux Headers](./CrossCompilationTools/LinuxHeaders.md)
1. [Cross-compiled GNU Binutils](./CrossCompilationTools/GNUBinutils.md)
1. [Cross-compiled GNU GCC](./CrossCompilationTools/GNUGCC.md)
1. [GNU GLibC 64-bit](./CrossCompilationTools/GNUGLibC64bit.md)
1. [GNU GLibC 32-bit](./CrossCompilationTools/GNUGLibC32bit.md)
1. [Cross-compiled LibStdC++ From GNU GCC](./CrossCompilationTools/LibStdC++.md)

## Temporary Tools

1. [Overview](./TempTools/Overview.md)
1. [GNU M4](./TempTools/GNUM4.md)
1. [NCurses TIC](./TempTools/NCursesTic.md)
1. [NCurses 64-bit](./TempTools/NCurses64bit.md)
1. [NCurses 32-bit](./TempTools/NCurses32bit.md)
1. [GNU Bash](./TempTools/GNUBash.md)
1. [GNU Coreutils](./TempTools/GNUCoreutils.md)
1. [GNU Diffutils](./TempTools/GNUDiffutils.md)
1. [File](./TempTools/File.md)
1. [GNU Findutils](./TempTools/GNUFindutils.md)
1. [GNU Awk](./TempTools/GAWK.md)
1. [GNU Grep](./TempTools/GNUGrep.md)
1. [GNU Gzip](./TempTools/GNUGzip.md)
1. [GNU Make](./TempTools/GNUMake.md)
1. [GNU Patch](./TempTools/GNUPatch.md)
1. [GNU Sed](./TempTools/GNUSed.md)
1. [GNU Tar](./TempTools/GNUTar.md)
1. [Xz](./TempTools/Xz.md)
1. [GNU Binutils - pass 2](./TempTools/GNUBinutils.md)
1. [GNU GCC - pass 2](./TempTools/GNUGCC.md)

## Entering the Chroot and More Temporary Tools

1. [Rationale](./ChrootTempTools/ChrootBuildingTempTools.md)
1. [Change Ownerships of Files and Directories](./ChrootTempTools/ChangeOwnerships.md)
1. [Mounting Virtual Kernel Filesystems](./ChrootTempTools/MountingVirtualKernelFilesystems.md)
1. [Entering the Chroot](./ChrootTempTools/EnteringChroot.md)
1. [Creating Essential Files and Symbolic Links](./ChrootTempTools/EssentialFilesAndSymlinks.md)
1. [Validating the C and C++ Compilers](./ChrootTempTools/ValidatingCompilers.md)
1. [Debian Hostname Utility](./ChrootTempTools/Hostname.md)
1. [GNU Gettext](./ChrootTempTools/GNUGettext.md)
1. [GNU Bison](./ChrootTempTools/GNUBison.md)
1. [Perl](./ChrootTempTools/Perl.md)
1. [Python 3.13](./ChrootTempTools/Python313.md)
1. [GNU Texinfo](./ChrootTempTools/GNUTexinfo.md)
1. [Util-Linux 64-bit](./ChrootTempTools/UtilLinux64bit.md)
1. [Util-Linux 32-bit](./ChrootTempTools/UtilLinux32bit.md)
1. [Clean-up and Archival](./ChrootTempTools/CleanupAndArchival.md)

## Building the Basic System Tools

1. [Overview](./BasicSystem/Overview.md)
1. [Man Pages](./BasicSystem/ManPages.md)
1. [IANA Etc Files](./BasicSystem/IanaEtcFiles.md)
1. [GNU GLibC 64-bit](./BasicSystem/GLibC64bit.md)
1. [GNU GLibC 32-bit](./BasicSystem/GLibC32bit.md)
1. [ZLib 64-bit](./BasicSystem/ZLib64bit.md)
1. [ZLib 32-bit](./BasicSystem/ZLib32bit.md)
1. [BZip2 64-bit](./BasicSystem/BZip264bit.md)
1. [BZip2 32-bit](./BasicSystem/BZip232bit.md)
1. [XZ 64-bit](./BasicSystem/XZ64bit.md)
1. [XZ 32-bit](./BasicSystem/XZ32bit.md)
1. [Lz4 64-bit](./BasicSystem/Lz464bit.md)
1. [Lz4 32-bit](./BasicSystem/Lz432bit.md)
1. [ZStd 64-bit](./BasicSystem/ZStd64bit.md)
1. [ZStd 32-bit](./BasicSystem/ZStd32bit.md)
1. [GNU libunistring 64-bit](./BasicSystem/libunistring64bit.md)
1. [GNU libunistring 32-bit](./BasicSystem/libunistring32bit.md)
1. [File 64-bit](./BasicSystem/File64bit.md)
1. [File 32-bit](./BasicSystem/File32bit.md)
1. [GNU Readline 64-bit](./BasicSystem/GNUReadline64bit.md)
1. [GNU Readline 32-bit](./BasicSystem/GNUReadline32bit.md)
1. [PCRE2 64-bit](./BasicSystem/PCRE264bit.md)
1. [PCRE2 32-bit](./BasicSystem/PCRE232bit.md)
1. [GNU M4](./BasicSystem/GNUM4.md)
1. [GNU Bc](./BasicSystem/GNUBc.md)
1. [Flex 64-bit](./BasicSystem/Flex64bit.md)
1. [Flex 32-bit](./BasicSystem/Flex32bit.md)
1. [TCL](./BasicSystem/TCL.md)
1. [Expect](./BasicSystem/Expect.md)
1. [DejaGNU](./BasicSystem/DejaGNU.md)
1. [PkgConf](./BasicSystem/PkgConf.md)
1. [GNU Binutils](./BasicSystem/GNUBinutils.md)
1. [GNU GMP 64-bit](./BasicSystem/GNUGMP64bit.md)
1. [GNU GMP 32-bit](./BasicSystem/GNUGMP32bit.md)
1. [GNU MPFR 64-bit](./BasicSystem/GNUMPFR64bit.md)
1. [GNU MPFR 32-bit](./BasicSystem/GNUMPFR32bit.md)
1. [GNU MPC 64-bit](./BasicSystem/GNUMPC64bit.md)
1. [GNU MPC 32-bit](./BasicSystem/GNUMPC32bit.md)
1. [ISL 64-bit](./BasicSystem/ISL64bit.md)
1. [ISL 32-bit](./BasicSystem/ISL32bit.md)
1. [GNU LibIDN2 64-bit](./BasicSystem/libidn2-64bit.md)
1. [GNU LibIDN2 32-bit](./BasicSystem/libidn2-32bit.md)

| Navigation |||
| --- | --- | ---: |
| | [HOME](./README.md) | Prerequisites [>>](./Prerequisites.md) |

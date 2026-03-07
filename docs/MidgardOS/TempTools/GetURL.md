# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection - Pass 2 | [HOME](../README.md) | Entering the Build Root [>>](../ChrootTempTools/ChrootBuildingTempTools.md) |

## GetURL

Name: GetURL<br />
Summary: A simple file download utility<br />
License: GPL v3<br />
Version: 1.0.0<br />

Average Build Time: les than 0.1 SBU<br />
Used Install Space: 8.0 MiB<br />

## Rationale

The GetURL tool is a simple download utility to allow downloading files from the Internet from inside the build root while `curl` is not available. This tool does not validate SSL/TLS certificates, since the root certificate bundle is not present yet in the environment. This tool is provided to ease the download of packages during the development process for MidgardOS, and should not be relied on after the `curl` utility is available.

## Compilation and Installation

This tool is written in Golang, so there is no configuration needed for the build. To compile GetURL, change directory into the `docs/MidgardOS/src/geturl` inside the git checkout of the documentation. and then run the following command:

```bash
go build
```

Finally, to install GetURL into the build tree, run the following command:

```bash
install -v -m 755 -o root -g root geturl /MidgardOS/usr/bin/
```

| Navigation |||
| --- | --- | ---: |
| [<<](./GNUGCC.md) GNU Compiler Collection - Pass 2 | [HOME](../README.md) | Entering the Build Root [>>](../ChrootTempTools/ChrootBuildingTempTools.md) |

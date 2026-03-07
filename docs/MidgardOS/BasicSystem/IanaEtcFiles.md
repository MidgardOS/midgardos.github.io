# Section 5 - Building the Base System Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./ManPages.md) Man Pages | [HOME](../README.md) | GNU GLibC 64-bit [>>](./GNUGLibC_64bit.md) |

## IANA Etc Files

Name: IANA Etc Files<br />
Summary: The `services` and `protocols` text databases<br />
License: MIT<br />
Version: 20260226<br />
URL: [https://github.com/Mic92/iana-etc/releases/download/20260226/iana-etc-20260226.tar.gz](https://github.com/Mic92/iana-etc/releases/download/20260226/iana-etc-20260226.tar.gz)<br />

Average Build Time: less than 0.1 SBU<br />
Used Install Space: 4.8 MiB<br />

## Installation

To install IANA Etc Files into place, run the following command:

```bash
install -v -m 644 -o root -g root services /etc/
install -v -m 644 -o root -g root protocols /etc/
```

## Contents

| Installed Files | Description |
| --- | --- |
| `/etc/protocols` | The list of protocols carried over TCP/IP |
| `/etc/services`  | The Assigned Names and Numbers for ports and service names |


| Navigation |||
| --- | --- | ---: |
| [<<](./ManPages.md) Man Pages | [HOME](../README.md) | GNU GLibC 64-bit [>>](./GNUGLibC_64bit.md) |

| Navigation |||
| --- | --- | ---: |
| [<<](./Prerequisites.md) | [HOME](./README.md) | [>>](./Structure.md) |

# Host System Requirements

To build a MidgardOS build root, you must already have a relatively recent Linux operating system installed on your target build computer. Your host system should have the following software with the minimum versions indicated. Also note that many distributions will place software headers into separate packages, often in the form of “[package-name]-devel” or “[package-name]-dev”. Be sure to install those if your distribution provides them.

| Tool or Library | Version |
| --- | --- |
| **Bash** | 4.1 or higher |
| **Binutils** | 2.30 or higher |
| **Bison** | 3.0 or higher |
| **Bzip2** | 1.0 or higher |
| **Coreutils** | 8.0 or higher |
| **Diffutils** | 3.6 or higher |
| **Findutils** | 4.8 or higher |
| **Gawk** | 4.2 or higher |
| **GCC and G++ Compilers** | 7 or higher |
| **GLibC** | 2.31 or higher |
| **Grep** | 3.1 or higher |
| **GZip** | 1.10 or higher |
| **GNU Make** | 4.2 or higher |
| **GNU Ncurses** | 5.3 or higher |
| **Patch** | 2.7 or higher |
| **GNU Sed** | 4.4 or higher |
| **GNU Tar** | 1.34 or higher |
| **Texinfo** | 6.5 or higher |
| **XZ Utils** | 5.2 or higher |

To see whether your host system has all the appropriate versions, create and run the following script. Read the output carefully for any errors, and make sure to install any packages that are reported as not found.

```bash
cat > version-check.sh << "EOF"
#!/bin/bash

# Simple script to list version numbers of critical development tools

bash --version | head -n1 | cut -d" " -f2-4
echo -n "Binutils: "; ld --version | head -n1 | cut -d" " -f3-
bison --version | head -n1
bzip2 --version 2>&1 < /dev/null | head -n1 | cut -d" " -f1,6-
echo -n "Coreutils: "; chown --version | head -n1 | cut -d")" -f2
diff --version | head -n1
find --version | head -n1
gawk --version | head -n1
gcc --version | head -n1
g++ --version | head -n1
ldd $(which ${SHELL}) | grep libc.so | cut -d ' ' -f 3 | ${SHELL} | head -n 1 | cut -d ' ' -f 1-7
grep --version | head -n1
gzip --version | head -n1
make --version | head -n1
tic -V
patch --version | head -n1
sed --version | head -n1
tar --version | head -n1
makeinfo --version | head -n1
xz --version | head -n1
echo 'int main(){}' | gcc -v -o /dev/null -x c - > dummy.log 2>&1
if ! grep -q ' error' dummy.log; then
  echo "Compilation successful" && rm dummy.log
else
  echo 1>&2  "Compilation FAILED - more development packages may need to be \
installed. If you like, you can also view dummy.log for more details."
fi
EOF

bash version-check.sh 2>errors.log &&
[ -s errors.log ] && echo -e "\nThe following packages could not be found:\n$(cat errors.log)"
```

Note that MidgardOS package builds require a system that is running x86-64 v2 or higher based hardware. To validate that your system is capable of this level of hardware support, run the following script on your host:

```awk
#!/usr/bin/awk -f

BEGIN {
    while (!/flags/) if (getline < "/proc/cpuinfo" != 1) exit 1
    if (/lm/&&/cmov/&&/cx8/&&/fpu/&&/fxsr/&&/mmx/&&/syscall/&&/sse2/) level = 1
    if (level == 1 && /cx16/&&/lahf/&&/popcnt/&&/sse4_1/&&/sse4_2/&&/ssse3/) level = 2
    if (level == 2 && /avx/&&/avx2/&&/bmi1/&&/bmi2/&&/f16c/&&/fma/&&/abm/&&/movbe/&&/xsave/) level = 3
    if (level == 3 && /avx512f/&&/avx512bw/&&/avx512cd/&&/avx512dq/&&/avx512vl/) level = 4
    if (level > 0) { print "CPU supports x86-64-v" level; exit level + 1 }
    exit 1
}
```

If it outputs x86-64-v2 or higher, your system is adequate for use:

```text
CPU supports x86-64-v2
```

| Navigation |||
| --- | --- | ---: |
| [<<](./Prerequisites.md) | [HOME](./README.md) | [>>](./Structure.md) |

# Section 4 - Entering the Chroot and More Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](./EssentialFilesAndSymlinks.md) Rationale | [HOME](../README.md) | Mounting Virtual Kernel Filesystems [>>](./GNUGettext.md) |

## Validating the C and C++ Compilers

To validate that the C and C++ compilers work as they should, run the following commands to validate that they are working as expected:

```bash
if [[ -f dummy-c.c ]]; then
  rm -f dummy-c.c
fi
if [[ -f dummy-c ]]; then
  rm -f dummy-c
fi
cat << EOF > dummy-c.c
#include <stdio.h>
int main() {
  printf("Hello world\n");
  return 0;
}
EOF
/bin/cc dummy-c.c -o dummy-c
echo "Exit code: $?"
if [[ -f dummy-cxx.cxx ]]; then
  rm -f dummy-cxx.cxx
fi
if [[ -f dummy-cxx ]]; then
  rm -f dummy-cxx
fi
cat << EOF > dummy-cxx.cxx
#include <iostream>

using namespace std;

int main() {
  cout << "Hello world" << endl;
  return 0;
}
EOF
/bin/g++ dummy-cxx.cxx -o dummy-cxx
echo "Exit code: $?"
```

If the exit codes for both tests come back as `0`, the compilers are working as expected and the remaining builds should work as they should. If, however, they do not, revert the VM and re-attempt building GCC, and verify that all the flags used are what is proscribed in this guide.

| Navigation |||
| --- | --- | ---: |
| [<<](./EssentialFilesAndSymlinks.md) Rationale | [HOME](../README.md) | Mounting Virtual Kernel Filesystems [>>](./GNUGettext.md) |


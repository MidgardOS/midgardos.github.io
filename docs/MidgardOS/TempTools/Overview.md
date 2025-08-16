# Section 3 - Temporary Tools

| Navigation |||
| --- | --- | ---: |
| [<<](../CrossCompilationTools/LibStdC++.md) | [HOME](../README.md) | [>>](./M4.md) |

## Cross Compiling Temporary Tools Overview

Now that the core cross-compilatioon tools are installed, we can build the remaining temporary tools to allow us to isolate the new build root from the host operating system. The tools in this section are built against the tools in the `/tools` tree, and will be installed in the final destination MidgardOS tree.

These tools will be used later when we `chroot` into the build environment, but still rely on the host operating system during this phase.

During this section of builds, you may see the following warning messages and can be safely ignored:

```
configure: WARNING: result yes guessed because of cross compilation
configure: WARNING: cannot check WCONTINUED if cross compiling -- defaulting to no
```

Please remember that these builds must be done under the `builder` user, and should use the environment variables that were setup earlier in this guide. Doing any of the builds in this section as the `root` user will cause the host operating system to become unusable.

| Navigation |||
| --- | --- | ---: |
| [<<](../CrossCompilationTools/LibStdC++.md) | [HOME](../README.md) | [>>](./M4.md) |

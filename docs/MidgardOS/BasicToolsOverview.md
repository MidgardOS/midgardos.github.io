| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGNUGCCP2.md) | [HOME](./README.md) | [>>](./TempToolsGMP.md) |

# Temporary Tools Overview And More Environment Variables

Now that the core cross-compilation tools are installed, we can build the remaining temporary tools to allow us to isolate the new build root from the host operating system. The tools in this section are built against the tools in the `/cross-tools` tree, and will be installed in the `/tools` tree.

During this section of the build, you may see the following warning messages and can be safely ignored:

```
configure: WARNING: result yes guessed because of cross compilation
configure: WARNING: cannot check WCONTINUED if cross compiling -- defaulting to no
```

## Extra Environment Variables

During this section of the build, some extra environment variables will be required to ensure that the builds use the newly built compilers and linkers. To set them, add the following to the builder user's `.bashrc` and then log out of the user and log back in again:

```bash
echo export CC=\""\${BRFS_TARGET}-gcc \${BUILD64}"\" >> ~/.bashrc
echo export CXX=\""\${BRFS_TARGET}-g++ \${BUILD64}"\" >> ~/.bashrc
echo export AR=\""\${BRFS_TARGET}-ar"\" >> ~/.bashrc
echo export AS=\""\${BRFS_TARGET}-as"\" >> ~/.bashrc
echo export RANLIB=\""\${BRFS_TARGET}-ranlib"\" >> ~/.bashrc
echo export LD=\""\${BRFS_TARGET}-ld"\" >> ~/.bashrc
echo export STRIP=\""\${BRFS_TARGET}-strip"\" >> ~/.bashrc
```

| Navigation |||
| --- | --- | ---: |
| [<<](./CrossCompileGNUGCCP2.md) | [HOME](./README.md) | [>>](./TempToolsGMP.md) |

| Navigation |||
| --- | --- | ---: |
| [<<](./SetEnvironmentVars.md) | [HOME](./README.md) | [>>](./CrossCompilationTools/Overview.md) |

# Section 1 - Preparation

# Ignoring Software Tests During the Cross-Complation Phase

Many software packages these days contain test suites that are meant to be used to validate and verify that the built
software is fit for purpose. Unfortunately, while building the cross-complation tools, these tests will in many cases
not function correctly, as the temporary tools used will not be in "normal" FHS required locations. Because of this, it
is strongly recommended to not run them during this stage of the build process. They will be used later, when the final
builds are run.

| Navigation |||
| --- | --- | ---: |
| [<<](./SetEnvironmentVars.md) | [HOME](./README.md) | [>>](./CrossCompilationTools/Overview.md) |

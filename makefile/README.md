https://makefiletutorial.com/#getting-started

## Makefile Syntax

```makefile
targets: prerequisites
  command
  command
  command
```

+ The targets are **filenames**, separated by spaces. Typically, there is only one per rule.
+ The commands start with a **tab** character, not spaces.
+ The prerequisites are also file names, separated by spaces. These files need to exist before the commands for the target are run. These are also called dependencies

+ The first target is the default target

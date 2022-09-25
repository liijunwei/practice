random notes while reading CSAPP...

"深入理解计算机系统" 原书第3版

[Computer Systems: A Programmer's Perspective, 3/E (CS:APP3e)](http://csapp.cs.cmu.edu/)

code: http://csapp.cs.cmu.edu/3e/code.html

+ OK make csapp.c csapp.h global available in practice/csapp/

+ [C code formatter: "centos安装clang-format"](https://www.jianshu.com/p/91265382bace)
```bash
# keyword: "yum install clang-format"
#     https://centos.pkgs.org/7/centos-sclo-rh-x86_64/llvm-toolset-7-git-clang-format-5.0.1-4.el7.x86_64.rpm.html
#     Install Howto

sudo yum install centos-release-scl-rh -y
yum search clang-format
sudo yum install llvm-toolset-7.0-git-clang-format.x86_64 -y

# macos
brew install clang-format

# usage
# https://clang.llvm.org/docs/ClangFormat.html
clang-format --help
clang-format -style=LLVM main.c
clang-format -style=LLVM -i **/*.c
```

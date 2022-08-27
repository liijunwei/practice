第二部分: 在系统上运行程序


Linking
==============

20220823


# 7.5 符号和符号表

+ "C程序员使用static属性隐藏模块内部的变量和函数声明, 就像你在java和c++中使用public/private声明一样; 在C中, 原文件扮演模块的角色. 任何带有static属性声明的全局变量或者函数都是模块私有的. 类似的, 任何不带static属性声明的全局变量和函数声明都是公共的, 可以被其他模块访问; 尽可能使用static属性来保护你的变量和函数是很好的编程习惯"

# 7.6 符号解析

## 7.6.1 链接器如何解析多重定义的全局符号

根据强弱符号的定义, linux链接器使用下面的规则来处理多重定义的符号名
+ 规则1: 不允许有多个同名的强符号
+ 规则2: 如果有一个强符号和多个弱符号同名, 那么选择强符号
+ 规则3: 如果有多个弱符号同名, 那么从这些弱符号中任意选择一个

+ ELF: Executable and Linkable Format

# 7.9 加载可执行目标文件

```bash
$ ./prog
$ ./a.out
```

+ prog/a.out 不是内置的shell命令, 所以shell会认为prog是一个可执行目标文件, 通过调用某个驻留在存储器中称为加载器(loader)的操作系统代码来运行它

+ 任何linux程序都可以通过调用`execve`函数来调用加载器
    + man execve

+ 这个将程序复制到内存并运行的过程叫做"加载"

+ 要理解家在实际是如何工作的, 你必须理解进程/虚拟内存和内存映射的概念


# 7.10 动态链接共享库

+ 静态库仍然有一些明显的缺点: 例如 静态库和所有软件一样, 需要定期维护和更新
    + rails项目容器化后, 有一种做法是先bundle, 把bundle后的镜像作为base镜像, 然后再构建应用, 这个base镜像就类似于这个静态库吧

+ 内存的一个有趣属性就是 不论系统的内存有多大, 它总是一种稀缺资源(interesting...)

+ 共享库(shared library)是一种致力于解决静态库缺陷的一个现代创新产物
    + 动态链接(dynamic linking)过程
    + 动态链接器(dynamic linker)
    + 共享目标(shared object), 在linux系统中通常用 `.so` 后缀表示(Shared Object, so)
    + 微软的操作系统大量使用了共享库, 他们称为DLL(动态链接库)

# 7.11 从应用程序中加载和链接共享库

```bash
man dlopen
```

# 7.12 位置无关代码

+ 共享库的一个主要目的就是允许多个正在运行的进程共享内存中相同的库代码, 因而节约宝贵的内存资源

+ 可以加载而无需重定位的代码称为位置无关代码(Position-Independent Code, PIC)

+ GOT: Global Offset Table, 全局偏移量表

# 7.13 库打桩机制(library interpositioning)

+ 库打桩机制允许你截获对共享库函数的调用, 取而代之执行自己的代码

## 7.13.1 编译时打桩

make ex_interpose_compile

## 7.13. 链接时打桩

make ex_interpose_link

## 7.14. 运行时打桩

make ex_interpose_runtime_int
make ex_interpose_runtime_ls


2022-08-27 15:00:50






















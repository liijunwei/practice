工具: 开发的战术
======================

## 15.2 编辑器的选择

+ 对于严肃的编辑工作, 两款编辑器完全统治了unix编程界(vi+emacs)

+ vi => "visual editor" /vee eye/, not /vie/

+ emacs => "Editing MACros" 宏编辑 /ee'maks/

+ 之前讨论编辑器的可能复杂度时, 我们注意到许多人都认为emacs太过沉重; 然而, 花时间学习它可以获得很大的生产力回报


## 15.4 make: 自动化编译

+ 关于通用生成目标应该表示什么, 以及他们该如何命名已经有了一组良好的约定; 遵从这些会让自己的makefile更容易理解和使用:(p360)

name      | convention
----------|------------------------------
all       | 生成工程中所有可以执行者
test      | 运行程序的自动测试套件
clean     | 删除`make all`时产生的所有文件; `make clean`应该将软件的编译过程重置到良好的初始状态
dist      | 制作原文件档案(通常使用`tar`程序归档), 他可以作为在另一台机器上重新编译的单元
disclean  | 删除所有文件, 除了那些使用make dist打包时指定包含的文件
realclean | 删除所有makefile构建的文件
install   | 在系统目录中安装项目工程的可执行文件和文档(通常需要root权限)以让普通用户访问
uninstall | 删除由make install安装在系统目录中的所有文件(通常需要root权限)


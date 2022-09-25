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

+ makefile tutorial: https://makefiletutorial.com/#getting-started
+ fetchmail example: https://github.com/Distrotech/fetchmail

+ 不要局限在这些通用目标上; 一旦掌握了make, 就会越来越频繁地使用makefile机制来自动化哪些依赖项目文件状态的小任务;
+ makefile就是一个方便存放完成这些小任务脚本的重要地方;
+ 而使用make可以检视如何完成这些小任务时一目了然, 并且避免了小脚本将项目工作空间弄得凌乱不堪


## 15.6 运行期调试

+ 牢记unix哲学; 将时间花费在设计质量上, 而不是低层次的细节上, 尽可能地自动化一切---包括运行期调试的细节工作

## 15.7 性能分析

./ch12-optimization.md

+ 作为通用法则, 程序90%的执行时间都耗费在10%的代码上; 性能分析软件可以帮助确定那10%抑制程序速度的区域

+ **但在unix传统中, profiler有个更为重要的功能: 它能够让人不去优化其余90%的代码; 这很棒, 并不仅仅由于这样节省了工作, 其实不去优化余下的90%代码真正有价值的效应是 降低了整体复杂度, 减少了bug**

+ 做好设计, 优先考虑什么是正确的, 然后在调整效率

+ profiler帮助完成这一切; 如果养成使用他们的好习惯, 可以改掉过早优化的坏习惯; profiler工具不仅改变工作的方式, 也改变思考的方式



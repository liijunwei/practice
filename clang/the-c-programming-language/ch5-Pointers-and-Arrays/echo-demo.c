#include <stdio.h>

/*
page 99

命令行参数
---------------
在支持C语言的环境中, 可在程序开始执行时, 将命令行参数传递给程序
调用主函数main时, 它带有两个参数: 第一个参数(argc)的值表示运行程序时命令行中参数的数目; 第二个参数(argv)是一个指向字符串数组
的指针, 其中每个字符串对应一个参数

我们通常用多级指针处理这些字符串(TODO ???)

最简单的例子是程序 echo
echo hello, world
*/

int main(int argc, char const *argv[])
{
  /* code */
  return 0;
}

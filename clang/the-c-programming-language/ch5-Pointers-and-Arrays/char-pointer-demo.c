#include <stdio.h>

/*
page 89

字符串常量 是一个字符数组
字符串常量 的最常见的用法也许是 作为函数的参数
字符串常量 可以通过一个指向其第一个元素的指针访问

除了作为函数参数外, 字符串常量还有其他用法
  假定pmessage的声明如下
  char *pmessage;

  那么, 语句:
  pmessage = "now is the time";

  将把一个指向该字符数组的指针赋值给pmessage; 该过程并没有字符串的复制, 只涉及到指针的操作

*/

int main(int argc, char const *argv[])
{
  // 下面两个定义之间有很大的区别:

  char amessage[] = "now is the time"; // 定义一个数组
  char *pmessage = "now is the time";  // 定义一个指针


  return 0;
}
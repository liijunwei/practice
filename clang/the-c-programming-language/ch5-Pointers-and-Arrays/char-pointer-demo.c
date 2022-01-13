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

  /**
   * 定义一个数组
   * amessage 仅仅是一个一维数组, 数组内的单个元素可以被修改, 但ameessage始终指向同一个存储位置
   */
  char amessage[] = "now is the time";

  /**
   * 定义一个指针
   * pmessage 是一个指针, 其初值指向一个字符串常量, 之后它可以被修改以指向其他地址
   * 但是如果试图修改字符串的内容, 结果是没有定义的
   */
  char *pmessage = "now is the time";


  return 0;
}
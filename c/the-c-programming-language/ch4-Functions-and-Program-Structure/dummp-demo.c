#include <stdio.h>

/*
page 59

函数定义的形式;
返回值类型 函数名(参数声明表)
{
  声明和语句
}

函数定义的各构成部分都可以省略, 最简单的函数如下所示
*/

void dummy(){}

int main(int argc, char const *argv[])
{
  dummy();
  printf("Dummy demo...\n");

  return 0;
}
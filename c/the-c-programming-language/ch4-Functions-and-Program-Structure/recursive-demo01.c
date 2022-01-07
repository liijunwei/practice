#include <stdio.h>

/*
page 74

把数字当做字符串打印

方法1. 把生成的数字一次存到数组里, 再逆序打印
方法2. 使用递归

*/

// 递归打印
// 该方法不能处理最大(小?)的负数
void printd(int n){
  if(n < 0){
    putchar('-');
    n = -n;
  }

  if(n / 10){
    printd(n / 10);
  }

  putchar(n % 10 + '0');
}

int main(int argc, char const *argv[])
{
  printd(998);

  return 0;
}
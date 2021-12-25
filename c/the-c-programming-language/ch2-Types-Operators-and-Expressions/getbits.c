#include <stdio.h>

/*
page 39

getbits 返回x中从右边第p位开始向右数n为的字段
这里假设最右边的一位是第0位, n与p都是合理的正值
例如: getbits(x, 4, 3) 返回x中第 4, 3, 2 三位的值
*/
unsigned getbits(unsigned x, int p, int n){
  return (x >> (p + 1 - n) & ~(~0 << n));
}

int main(int argc, char const *argv[])
{
  // 0000 1000
  printf("%d\n", getbits(100, 4, 5));
  printf("%d\n", getbits(100, 4, 4));
  printf("%d\n", getbits(100, 4, 3));
  printf("%d\n", getbits(100, 4, 2));
  printf("%d\n", getbits(100, 4, 1));
  return 0;
}
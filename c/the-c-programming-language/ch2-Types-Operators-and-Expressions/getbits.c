#include <stdio.h>

/*
page 39

getbits(x, p, n) 返回x中从右边第p位开始向右数n位的字段
这里假设最右边的一位是第0位, n与p都是合理的正值
例如: getbits(x, 4, 3) 返回x中第 4, 3, 2 三位的值

generate sample data
--------------------
toBinary 78  => 0100 1110
toBinary 89  => 0101 1001
toBinary 657 => 0010 1001 0001
*/
unsigned getbits(unsigned x, int p, int n){
  return (x >> (p + 1 - n) & ~(~0 << n));
}

int main(int argc, char const *argv[])
{
  printf("%d\n", getbits(255, 4, 3)); // 7
  printf("\n");

  printf("%d\n", getbits(78, 4, 3)); // 3
  printf("%d\n", getbits(78, 4, 2)); // 1
  printf("%d\n", getbits(78, 4, 1)); // 0
  printf("%d\n", getbits(78, 4, 0)); // 0 n为0 这个有意义吗? 不确定
  printf("\n");

  printf("%d\n", getbits(89, 4, 3)); // 6
  printf("%d\n", getbits(89, 4, 2)); // 3
  printf("%d\n", getbits(89, 4, 1)); // 1
  printf("%d\n", getbits(89, 4, 0)); // 0
  printf("\n");

  printf("%d\n", getbits(657, 6, 3)); // 1
  printf("%d\n", getbits(657, 6, 2)); // 0
  printf("%d\n", getbits(657, 6, 1)); // 0
  printf("%d\n", getbits(657, 6, 0)); // 0
  printf("\n");
  return 0;
}
#include <stdio.h>
#include <assert.h>

/*
page 40

编写一个函数 invert(x,p,n)，该函数返回对x执行下列操作后的结果值:
将x中从第p位开始的n个(二进制)位求反(即，1变成0，0变成1)，x的其余各位保持不变

*/

unsigned invert(unsigned x, int p, int n){
  int tmp = ~(~0 << n) << (p - n + 1);
  return x ^ tmp;
}

// generate sample data
// --------------------
// toBinary 78   => 01001110
// toBinary 89   => 01011001
// toBinary 657  => 001010010001
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[])
{
  assert(126 == invert(78, 5, 2)); // 01111110
  assert(62  == invert(78, 6, 3)); // 00111110
  assert(369 == invert(657, 9, 5)); // 000101110001
  assert(9675 == invert(8779, 10, 4)); // 10010111001011
  printf("PASS.\n");

  return 0;
}
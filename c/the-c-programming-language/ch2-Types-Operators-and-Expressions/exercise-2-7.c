#include <stdio.h>
#include <assert.h>

/*
page 40

编写一个函数 invert(x,p,n)，该函数返回对x执行下列操作后的结果值:
将x中从第p位开始的n个(二进制)位求反(即，1变成0，0变成1)，x的其余各位保持不变

*/

unsigned invert(unsigned x, int p, int n){

  return 100;
}

// generate sample data
// --------------------
// toBinary 78   => 01001110
// toBinary 89   => 01011001
// toBinary 657  => 00101001 0001
// toBinary 8779 => 1000 10010 01011
int main(int argc, char const *argv[])
{
  assert(100 == invert(78, 5, 2));
  assert(100 == invert(78, 5, 2));
  assert(100 == invert(8779, 9, 5));
  assert(100 == invert(8779, 9, 5));
  printf("PASS.\n");

  return 0;
}
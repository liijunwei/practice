#include <stdio.h>
#include <assert.h>

/*
page 40

编写一个函数setbits(x, p, n, y)
该函数返回对x执行下列操作后的结果值: 将x中从第p位开始的n个(二进制)位设置为y中最右边n位的值，x的其余各位保持不变

1. 把x中的n位 置为0
2. 把y中的n位 移动到p位
3. x | y
*/

unsigned setbits(unsigned x, int p, int n, int y){
  int tmpx = x & ~(~(~0 << n) << (p - n + 1));
  int tmpy = (y & ~(~0 << n)) << (p - n + 1);

  return tmpx | tmpy;
}

// generate sample data
// --------------------
// toBinary 78   => 01001110
// toBinary 89   => 01011001
// toBinary 657  => 00101001 0001
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[])
{

  assert(94  == setbits(78, 5, 2, 89));   // 01011110
  assert(126 == setbits(78, 5, 2, 8779)); // 01111110
  printf("PASS.\n");
  return 0;
}
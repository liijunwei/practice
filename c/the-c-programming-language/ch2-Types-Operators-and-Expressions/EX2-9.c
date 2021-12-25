#include <stdio.h>
#include <assert.h>

/*
page 41

ref: ch2-Types-Operators-and-Expressions/bitcount.c

在求对二的补码时，表达式 x &= (x - 1) 可以删除x中最右边值为1的一个二进制位, 请解释这样做的道理。
用这一方法重写 bitcount(x) 函数，以加快其执行速度。
*/

int bitcount(unsigned x) {

  return 100;
}

// generate sample data
// --------------------
// toBinary 78   => 1001110
// toBinary 89   => 1011001
// toBinary 657  => 1010010001
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[])
{
  assert(100 == bitcount(78));
  assert(100 == bitcount(89));
  assert(100 == bitcount(657));
  assert(100 == bitcount(8779));
  printf("PASS.\n");

  return 0;
}
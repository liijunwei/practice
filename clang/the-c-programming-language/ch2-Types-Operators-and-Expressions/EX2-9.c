#include <assert.h>
#include <stdio.h>

/*
page 41

ref: ch2-Types-Operators-and-Expressions/bitcount.c

在求对二的补码时，表达式 x &= (x - 1) 可以删除x中最右边值为1的一个二进制位,
请解释这样做的道理。 用这一方法重写 bitcount(x) 函数，以加快其执行速度。

注意: x最右边值为1的二进制位, 不一定是最右边的位

*/

int bitcount(unsigned x) {
  int b;

  for (b = 0; x != 0; x &= (x - 1)) {
    ++b;
  }

  return b;
}

// generate sample data
// --------------------
// toBinary 78   => 1001110
// toBinary 89   => 1011001
// toBinary 657  => 1010010001
// toBinary 1023 => 1111111111
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[]) {
  assert(4 == bitcount(78));
  assert(4 == bitcount(89));
  assert(4 == bitcount(657));
  assert(10 == bitcount(1023));
  assert(6 == bitcount(8779));
  printf("PASS.\n");

  return 0;
}
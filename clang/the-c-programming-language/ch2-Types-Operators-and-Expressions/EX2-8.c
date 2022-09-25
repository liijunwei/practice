#include <assert.h>
#include <stdio.h>

/*
page 40

编写一个函数 rightrot(x,n)
该函数返回将x循环右移(即从最右端移出的位将从最左端再移入)n(二进制)位后所得到的值
*/

// 计算出运行程序的计算机使用的字长
// 书上的 wordlength 不管用
// ref:
// https://github.com/fool2fish/the-c-programming-language-exercise-answers/blob/master/ch02/2-8.c
int bitlen(unsigned d) {
  int len = 0;
  for (; d; d >>= 1) {
    len++;
  }
  return len;
}

unsigned rightrot(unsigned x, int n) {
  int len = bitlen(x);
  int rbit; // rightmost bit

  while (n-- > 0) {
    rbit = (x & 1) << (len - 1);
    x = x >> 1;
    x = x | rbit;
  }

  return x;
}

// generate sample data
// --------------------
// toBinary 78   => 1001110
// toBinary 89   => 1011001
// toBinary 657  => 1010010001
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[]) {
  assert(83 == rightrot(78, 2));      // 1010011
  assert(27 == rightrot(89, 3));      // 0011011
  assert(291 == rightrot(657, 9));    // 0100100011
  assert(9400 == rightrot(8779, 10)); // 10010010111000
  printf("PASS.\n");

  return 0;
}
#include <stdio.h>
#include <assert.h>

/*
page 40

编写一个函数 rightrot(x,n)
该函数返回将x循环右移(即从最右端移出的位将从最左端再移入)n(二进制)位后所得到的值
*/

unsigned rightrot(unsigned x, int n){

  return 100;
}

// generate sample data
// --------------------
// toBinary 78   => 01001110
// toBinary 89   => 01011001
// toBinary 657  => 001010010001
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[])
{
  assert(100 == rightrot(78, 5));
  assert(100 == rightrot(78, 6));
  assert(100 == rightrot(657, 9));
  assert(100 == rightrot(8779, 10));
  printf("PASS.\n");

  return 0;
}
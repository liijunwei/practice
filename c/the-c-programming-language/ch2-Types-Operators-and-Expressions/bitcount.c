#include <stdio.h>
#include <assert.h>

/*
page 40
*/

int bitcount(unsigned x) {
  int b;
  for (b = 0; x != 0; x >>= 1) {
    b++;
  }

  return b;
}

// generate sample data
// --------------------
// toBinary 78   => 1001110
// toBinary 89   => 1011001
// toBinary 657  => 1010010001
// toBinary 8779 => 10001001001011
int main(int argc, char const *argv[])
{
  assert(7 == bitcount(78));
  assert(7 == bitcount(89));
  assert(10 == bitcount(657));
  assert(14 == bitcount(8779));
  printf("PASS.\n");

  return 0;
}
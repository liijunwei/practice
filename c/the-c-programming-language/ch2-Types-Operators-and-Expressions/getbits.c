#include <stdio.h>

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
#include <stdio.h>

long rfun(unsigned long x) {
  if (x == 0) {
    return 0;
  }

  unsigned long nx = x >> 2;
  long rv = rfun(nx);

  return x + rv;
}

int main(int argc, char const *argv[]) {
  printf("%lu \n", rfun(0));
  printf("%lu \n", rfun(1));
  printf("%lu \n", rfun(2));
  printf("%lu \n", rfun(3));
  printf("%lu \n", rfun(4));
  printf("%lu \n", rfun(5));

  return 0;
}
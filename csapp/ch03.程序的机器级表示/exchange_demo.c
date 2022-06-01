#include <stdio.h>

long exchange(long *xp, long y);

int main(int argc, char const *argv[]) {
  long a = 4;
  long b = exchange(&a, 3);
  printf("a = %ld b = %ld\n", a, b);

  return 0;
}

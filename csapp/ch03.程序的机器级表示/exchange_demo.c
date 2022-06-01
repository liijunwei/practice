#include <stdio.h>

// cd csapp/ch03.程序的机器级表示 && gcc exchange_demo.c exchange.c && ./a.out && cd -

long exchange(long *xp, long y);

int main(int argc, char const *argv[]) {
  long a = 100;
  long b = 200;

  printf("a = %ld b = %ld\n", a, b);
  exchange(&a, b);
  printf("a = %ld b = %ld\n", a, b);

  return 0;
}

#include <stdio.h>

/**
 * p344
 */

void swap(long *xp, long *yp) {
  *xp = *xp + *yp; /* x + y     */
  *yp = *xp - *yp; /* x + y - y */
  *xp = *xp - *yp; /* x + y - x */
}


int main(int argc, char const *argv[]) {
  long a = 100;
  long b = 200;
  printf("swap diff: a: %ld b: %ld\n", a, b);
  swap(&a, &b);
  printf("swap diff: a: %ld b: %ld\n", a, b);

  printf("\n");

  long c = 100;
  printf("swap same: c: %ld c1: %ld\n", c, c);
  swap(&c, &c);
  printf("swap same: c: %ld c1: %ld\n", c, c);

  return 0;
}


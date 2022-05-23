#include <stdio.h>

// page 38
void inplace_swap(int *x, int *y) {
  *y = *x ^ *y;
  *x = *x ^ *y;
  *y = *x ^ *y;
}


int main(int argc, char const *argv[])
{
  int a = 10;
  int b = 20;
  printf("a: %d, b: %d\n", a, b);

  inplace_swap(&a, &b);
  printf("a: %d, b: %d\n", a, b);

  return 0;
}

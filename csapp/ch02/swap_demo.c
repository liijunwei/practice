#include <stdio.h>
#include <assert.h>

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

  assert(a == 10); assert(b == 20);
  inplace_swap(&a, &b);
  assert(a == 20); assert(b == 10);

  printf("PASS.\n");

  return 0;
}

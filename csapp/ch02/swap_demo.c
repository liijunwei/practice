#include <stdio.h>
#include <assert.h>

// XOR的属性: 对于任意位向量a, 有 a^a = 0 (与自身做XOR运算)

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

  assert(a == 10);
  assert(b == 20);

  inplace_swap(&a, &b);

  assert(a == 20);
  assert(b == 10);

  printf("PASS.\n");

  return 0;
}

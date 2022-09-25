#include <stdio.h>

// page 20

// 对比: ch1-A-Tutorial-Introduction/power-1.c

// 理解程序, 从理解 "幂" 这个概念开始
int power(int base, int n);

int main() {
  int i;

  printf("power | 2^power | 3^power\n");
  printf("------|---------|--------\n");

  for (i = 0; i < 10; ++i) {
    printf("%3d   |   %3d   |  %3d\n", i, power(2, i), power(-3, i));
  }

  return 0;
}

// 求底数 base 的 n次幂; 其中n >= 0
int power(int base, int n) {
  int p;

  for (p = 1; n > 0; --n) {
    p = p * base;
  }

  return p;
}
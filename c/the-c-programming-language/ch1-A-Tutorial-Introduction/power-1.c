#include <stdio.h>

// page 23

// 函数原型里的参数名是可选的
// 但是合适的参数名能起到很好的说明作用, 所以我们在函数原型里, 总是指明参数名
int power(int, int);

int main(){
  int i;

  printf("power | 2^power | 3^power\n");
  printf("------|---------|--------\n");

  for(i = 0; i < 10; ++i){
    printf("%3d   |   %3d   |  %3d\n", i, power(2, i), power(-3, i));
  }

  return 0;
}

// 求底数 base 的 n次幂; 其中n >= 0
int power(base, n){
  int p = 1;

  for(int i = 1; i <= n; ++i){
    p = p * base;
  }

  return p;
}
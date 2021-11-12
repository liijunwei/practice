#include <stdio.h>

int power(int m, int n);

int main(){
  int i;

  for(i = 0; i < 10; ++i){
    printf("%d %d %d\n", i, power(2, i), power(-3, i));
  }

  return 0;
}

// 求底数 base 的 n次幂; 其中n >= 0
int power(base, n){
  int i;
  int p = 1;

  for(i = 1; i <= n; ++i){
    p = p * base;
  }

  return p;
}
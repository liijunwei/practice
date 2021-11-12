#include <stdio.h>

void fahr2celsius();

// 使用函数重新编写 温度转换程序
int main(){
  fahr2celsius();
}

void fahr2celsius(){
  float lower = 0;
  float upper = 300;
  float step = 20;

  float fahr, celsius;
  fahr = lower;

  printf("fahr celsius\n");

  while(fahr <= upper) {
    celsius = (5.0 / 9.0) * (fahr - 32);
    printf("%3.0f %6.1f\n", fahr, celsius);
    fahr += step;
  }
}
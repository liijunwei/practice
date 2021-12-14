#include <stdio.h>

// implicit declaration of function 'celsius' is invalid in C99
float celsius(float fahr);

// page 19

// 对比: ch0-Introduction/fahr_to_celsius.c
// 使用函数重新编写 温度转换程序
int main(){
  float lower = 0;
  float upper = 300;
  float step  = 20;

  float fahr  = lower;

  printf("fahr celsius\n");
  while(fahr <= upper) {
    printf("%3.0f %6.1f\n", fahr, celsius(fahr));
    fahr += step;
  }
}

float celsius(float fahr){
  return ((5.0 / 9.0) * (fahr - 32));
}

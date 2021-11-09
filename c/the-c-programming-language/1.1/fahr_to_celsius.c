#include <stdio.h>

// 打印摄氏和华氏度转换
int main()
{

  float lower = 0;
  float upper = 300;
  float step = 20;

  float fahr, celsius;
  fahr = lower;

  printf("fahr celsius\n");

  while(fahr <= upper) {
    celsius = (5.0 / 9.0) * (fahr - 32);
    printf("%3.0f %6.1f\n", fahr, celsius);
    fahr = fahr + step;
  }
}


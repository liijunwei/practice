#include <stdio.h>

// 打印摄氏和华氏度转换
int main()
{
  float lower = -100;
  float upper = 100;
  float step = 20;

  float fahr, celsius;
  celsius = lower;

  printf("fahr \t celsius\n");
  printf("------------------\n");

  while(celsius <= upper) {
    fahr = (9.0 / 5.0) * celsius + 32;
    printf("%3.0f°C \t %6.1f°F\n", celsius, fahr);
    celsius += step;
  }
}


#include <stdio.h>

// https://www.zhihu.com/question/62003033
int main(){
  int a = 1;
  int b = 2;

  printf("Before exchange: a = %d, b = %d\n", a, b);

  a = a^b; // step.1
  b = a^b; // step.2
  a = a^b; // step.3
  printf("After exchange:  a = %d, b = %d\n", a, b);

  return 0;
}

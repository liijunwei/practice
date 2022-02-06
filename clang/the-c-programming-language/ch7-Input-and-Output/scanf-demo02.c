/*
page 139

计算器程序
*/

#include <stdio.h>

// gcc -g ch7-Input-and-Output/scanf-demo02.c && echo 1.1 2.2 3.3 | ./a.out
int main(int argc, char const *argv[])
{
  double sum = 0;
  double v;

  while (scanf("%lf", &v) == 1) {
    printf("\t%.2f\n", sum += v);
  }

  return 0;
}

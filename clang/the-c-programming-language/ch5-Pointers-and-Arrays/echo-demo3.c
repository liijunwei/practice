#include <stdio.h>

/*
page 100

V3
*/

int main(int argc, char const *argv[])
{
  while(--argc > 0){
    // printf的格式化参数也可以是表达式
    printf((argc > 1) ? "%s " : "%s", *++argv);
  }

  printf("\n");

  return 0;
}

#include <stdio.h>

// page 48

// 统计数字/空白及其他字符
int main(int argc, char const *argv[])
{
  int c;
  int i;
  int nwhite = 0;
  int nother = 0;
  int ndight[10];

  for(i = 0; i < 10; i++){
    ndight[i] = i;
  }

  while((c = getchar()) != EOF){

  }

  printf("dights =");
  for(i = 0; i < 10; i++){
    printf(" %d", ndight[i]);
  }

  printf(", white space = %d, other = %d\n", nwhite, nother);

  return 0;
}
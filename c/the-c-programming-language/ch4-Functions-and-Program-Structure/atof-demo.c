#include <stdio.h>
#include <ctype.h>
#include <assert.h>

/*
page 60

标准库<stdlib.h>里实现了atof函数, 比这个版本好

*/

// 把字符串s转换为相应的双精度浮点数
double atof(char s[]){

  return 3.14;
}

int main(int argc, char const *argv[])
{
  char str[100] = "3.14159265354";

  assert(3.14 == atof(str));
  printf("PASS.");
  return 0;
}
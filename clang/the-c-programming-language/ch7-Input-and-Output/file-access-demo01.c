/*
page 141

stdio.h中定义了结构FILE

*/

#include <stdio.h>

int main(int argc, char const *argv[]) {
  FILE *fp;

  fp = fopen("$HOME/practice/clang/the-c-programming-language/ch7-Input-and-Output/file-access-demo01.c", "r");
  fclose(fp);

  return 0;
}

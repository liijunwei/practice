/*
page 141

stdio.h中定义了结构FILE

*/

#include <stdio.h>

int main(int argc, char const *argv[]) {
  FILE *fp;
  fp = fopen("./ch7-Input-and-Output/file-access-demo01.c", "r");

  FILE *fpcopy;
  fpcopy = fopen("./ch7-Input-and-Output/file-access-demo01-copy.c", "w");

  int c;

  while ((c = getc(fp)) != EOF) {
    putc(c, fpcopy);
  }

  fclose(fp);
  fclose(fpcopy);

  return 0;
}

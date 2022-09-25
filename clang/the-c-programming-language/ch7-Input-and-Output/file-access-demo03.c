/*
page 141

ref:
  ch7-Input-and-Output/file-access-demo02.c

*/

#include <ctype.h>
#include <stdio.h>

int main(int argc, char const *argv[]) {
  FILE *fp;
  fp = fopen("./ch7-Input-and-Output/file-access-demo01.c", "r");

  int c;

  while ((c = getc(fp)) != EOF) {
    putc(toupper(c), stdout);
  }

  fclose(fp);

  return 0;
}

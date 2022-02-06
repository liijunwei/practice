/*
page 141

ref:
  ch7-Input-and-Output/file-access-demo03.c

*/

#include <stdio.h>
#include <ctype.h>

// gcc ch7-Input-and-Output/file-access-demo04.c && echo "helloworld" | ./a.out
int main(int argc, char const *argv[]) {
  int c;

  while ((c = getc(stdin)) != EOF) {
    putc(toupper(c), stdout);
  }

  return 0;
}

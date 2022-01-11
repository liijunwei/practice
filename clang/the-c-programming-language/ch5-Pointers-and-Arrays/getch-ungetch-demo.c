#include <stdio.h>

#include "../common-utils/getch-ungetch.c"
#include "../common-utils/print-array.c"

/*
理解 getch 和 ungetch函数

*/

int main(int argc, char const *argv[])
{
  ungetch('1');
  ungetch('2');
  ungetch('3');

  print_char_array(buf, BUFSIZE);
  putchar(getch());
  putchar(getch());
  putchar(getch());

  putchar(getch());
  putchar(getch());
  putchar(getch());

  return 0;
}

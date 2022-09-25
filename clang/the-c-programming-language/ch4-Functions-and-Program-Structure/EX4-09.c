#include <stdio.h>
#include <string.h>

/*
page 68

以上介绍的getch和ungetch函数不能正确地处理压回的EOF
考虑压回EOF时应该如何处理, 请实现你的设计方案

TODO 没看懂...

*/

#define BUFSIZE 100
int buf[BUFSIZE]; // char buf[BUFSIZE] -> int buf[BUFSIZE]
int bufp = 0;

int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

void ungetch(int c) {
  if (bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}

int main(int argc, char const *argv[]) {
  ungetch('h');
  ungetch('e');
  ungetch('l');
  ungetch('l');
  ungetch('o');
  ungetch(' ');
  ungetch('l');
  ungetch('j');
  ungetch('w');
  ungetch(EOF);

  int c;

  while (bufp > 0) {
    c = getch();
    putchar(c);
  }

  return 0;
}

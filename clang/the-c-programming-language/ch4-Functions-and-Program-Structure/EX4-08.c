#include <assert.h>
#include <stdio.h>

/*
page 67

假如最多只压回一个字符, 请相应地修改getch与ungetch这两个函数

*/

// 缓冲区不再是数组, 而是一个字符
char buf = 0;

int getch(void) {
  int c;
  if (buf != 0) {
    c = buf;
  } else {
    c = getchar();
  }

  buf = 0;

  return c;
}

void ungetch(int c) {
  if (buf != 0) {
    printf("error: buf has too many characters\n");
  } else {
    buf = c;
  }
}

int main(int argc, char const *argv[]) {

  // ungetch('b');
  ungetch('b');
  assert(getch() == 'b');

  ungetch('c');
  assert(getch() == 'c');

  ungetch('Z');
  assert(getch() == 'Z');

  return 0;
}

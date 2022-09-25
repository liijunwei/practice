#include <ctype.h>
#include <string.h>

#define BUFSIZE 100
#define NUMBER '0' // signal that a number was found

char buf[BUFSIZE];
int bufp = 0; // buf中下一空闲的位置

// 取一个字符(可能是压回的字符)
int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

// 把字符压回输入中
void ungetch(int c) {
  if (bufp >= BUFSIZE) {
    printf("Ungetch, too many characters\n");
  } else {
    buf[bufp++] = c;
  }
}

void ungets(char s[]) {
  int len = strlen(s);

  while (len > 0) {
    --len;
    ungetch(s[len]);
  }
}

int getop(char s[]) {
  int i;
  int c;

  while ((s[0] = c = getch()) == ' ' || c == '\t') {
    ;
  }

  s[1] = '\0';
  if (!isdigit(c) && c != '.') {
    return c; // 不是数
  }

  i = 0;
  if (isdigit(c)) {
    // 整数部分
    while (isdigit(s[++i] = c = getch())) {
      ;
    }
  }

  if (c == '.') {
    // 小数部分
    while (isdigit(s[++i] = c = getch())) {
      ;
    }
  }

  s[i] = '\0';

  if (c != EOF) {
    ungetch(c);
  }

  return NUMBER;
}

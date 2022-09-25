#include <stdio.h>
#include <string.h>

/*
page 67

编写一个函数 ungets(s), 将整个字符串s压回到输入中
ungets 函数需要使用buf和bufp吗? (需要)
它是否仅使用ungetch函数? (不能, 还需要使用string.h的方法)


*/

#define BUFSIZE 100
char buf[BUFSIZE];
int bufp = 0;

int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

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

int main(int argc, char const *argv[]) {
  char str[100] = "helloworld.\n";
  int c;

  ungets(str);

  while ((c = getch()) != EOF) {
    putchar(c);
  }

  return 0;
}

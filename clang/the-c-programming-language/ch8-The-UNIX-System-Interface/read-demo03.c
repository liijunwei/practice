/*
page 151

为了更好的掌握有关概念, 下面使用read和write构造类似于getchar/putchar等的高级函数

*/

#include <stdio.h>
#include <unistd.h>

/* getchar(v2) */
int c_getchar() {
  static char buf[BUFSIZ];
  static char *bufp = buf;
  static int n = 0;

  /* 缓冲区为空 */
  if (n == 0) {
    n = read(0, buf, sizeof(buf));
    bufp = buf;
  }

  return (--n >= 0) ? (unsigned char)*bufp++ : EOF;
}

// gcc ch8-The-UNIX-System-Interface/read-demo03.c && echo "eveninig" | ./a.out
int main(int argc, char const *argv[]) {
  char c;

  while ((c = c_getchar()) != EOF) {
    putchar(c);
  }

  return 0;
}

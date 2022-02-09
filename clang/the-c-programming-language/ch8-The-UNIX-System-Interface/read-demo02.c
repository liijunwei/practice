/*
page 151

为了更好的掌握有关概念, 下面使用read和write构造类似于getchar/putchar等的高级函数

*/

#include <stdio.h>
#include <unistd.h>

/* getchar(v1) */
int c_getchar() {
  int fd_stdin  = 0;
  char c;

  return (read(fd_stdin, &c, 1) == 1) ? (unsigned char) c : EOF;
}


int main(int argc, char const *argv[]) {
  char c;

  while ((c = c_getchar()) != EOF) {
    putchar(c);
  }

  return 0;
}

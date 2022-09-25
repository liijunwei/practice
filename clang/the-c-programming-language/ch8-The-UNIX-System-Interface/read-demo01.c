/*
page 150

*/

#include <stdio.h>
#include <unistd.h>

/* TODO 问题: BUFSIZ 为1 和 为 100 有什么区别呢? */

int main(int argc, char const *argv[]) {
  char buf[BUFSIZ]; /* BUFSIZ 在 <stdio.h> 中定义 */
  int n;

  printf("BUFSIZ %d\n", BUFSIZ);

  int fd_stdin = 0;
  int fd_stdout = 1;

  while ((n = read(fd_stdin, buf, BUFSIZ)) > 0) {
    write(fd_stdout, buf, n);
  }

  return 0;
}

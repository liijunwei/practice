/*
page 153

用read/write/open/close系统调用替代标准库中功能等价的函数, 重写第7章的cat程序,
并通过实验比较两个版本的相对执行速度 ch7-Input-and-Output/cat-demo02.c

*/

#include <fcntl.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>

#define FD_STDIN 0
#define FD_STDOUT 1

void error(char *fmt, ...);
void filecopy(int ifd, int ofd);

// gcc ch8-The-UNIX-System-Interface/EX8-01.c && ./a.out demo.c # OK
// gcc ch8-The-UNIX-System-Interface/EX8-01.c && ./a.out demo.ccc # ERROR
/* 连接多个文件(V2) */
int main(int argc, char const *argv[]) {
  clock_t start_time = clock();

  int fd;

  if (argc == 1) {
    filecopy(FD_STDIN, FD_STDOUT);
  } else {
    while (--argc > 0) {
      if ((fd = open(*++argv, 0, O_RDONLY)) == -1) {
        error("cat: can't open %s", *argv);
      } else {
        filecopy(fd, FD_STDOUT);
        close(fd);
      }
    }
  }

  double elapsed_time = (double)(clock() - start_time) / CLOCKS_PER_SEC;
  printf("Done in %f seconds\n", elapsed_time);

  return 0;
}

/* copy file ifd to file ofd */
void filecopy(int ifd, int ofd) {
  int n;
  char buf[BUFSIZ];

  while ((n = read(ifd, buf, BUFSIZ)) > 0) {
    if (write(ofd, buf, n) != n) {
      error("cat: write error");
    }
  }
}

/* print error message and skip the rest of the line */
void error(char *fmt, ...) {
  va_list args;

  va_start(args, fmt);
  fprintf(stderr, "error: ");
  vfprintf(stderr, fmt, args);
  fprintf(stderr, "\n");
  va_end(args);

  exit(1);
}

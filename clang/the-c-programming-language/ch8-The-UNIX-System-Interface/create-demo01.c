#include <fcntl.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define PERMS 0666

void error(char *fmt, ...);

// gcc ch8-The-UNIX-System-Interface/create-demo01.c && ./a.out
// ch8-The-UNIX-System-Interface/create-demo01.c
// ch8-The-UNIX-System-Interface/tmp.c
/* copy: 将文件f1复制到文件f2 */
int main(int argc, char const *argv[]) {
  int f1;
  int f2;
  int n;

  char buf[BUFSIZ];

  if (argc != 3) {
    error("Usage: cp from to");
  }

  if ((f1 = open(argv[1], O_RDONLY, 0)) == -1) {
    error("cp: can't open %s", argv[1]);
  }

  if ((f2 = creat(argv[2], PERMS)) == -1) {
    error("cp: can't create %s, mode %03o", argv[2], PERMS);
  }

  while ((n = read(f1, buf, BUFSIZ)) > 0) {
    if (write(f2, buf, n) != n) {
      error("cp: write error on file %s\n", argv[2]);
    }
  }

  return 0;
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

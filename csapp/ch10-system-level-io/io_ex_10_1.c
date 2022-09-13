#include "csapp.h"

int main(int argc, char const *argv[]) {
  int fd1;
  int fd2;

  fd1 = Open("foo.txt", O_RDONLY, 0);
  Close(fd1);

  fd2 = Open("baz.txt", O_RDONLY, 0);
  printf("fd2 = %d\n", fd2);

  exit(0);
}

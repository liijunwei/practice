#include "csapp.h"

int main(int argc, char const *argv[]) {
  int fd1;
  int fd2;

  char c;

  fd1 = Open("foobar.txt", O_RDONLY, 1);
  fd2 = Open("foobar.txt", O_RDONLY, 1);

  Read(fd1, &c, 1);
  Read(fd2, &c, 1);

  printf("c = %c\n", c);

  exit(0);
}

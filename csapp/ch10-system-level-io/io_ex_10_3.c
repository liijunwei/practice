#include "csapp.h"

int main(int argc, char const *argv[]) {
  int fd;
  char c;

  fd = Open("foobar.txt", O_RDONLY, 1);

  if(Fork() == 0) {
    Read(fd, &c, 1);
    exit(0);
  }

  Wait(NULL);
  Read(fd, &c, 1);

  printf("c = %c\n", c);

  return 0;
}

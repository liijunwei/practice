/*
page 150

*/

#include <stdio.h>
#include <unistd.h>

#define BUFSIZE 100

int main(int argc, char const *argv[]) {
  char buf[BUFSIZE];
  int n;

  int fd_stdin  = 0;
  int fd_stdout = 1;

  while ((n = read(fd_stdin, buf, BUFSIZE)) > 0) {
    write(fd_stdout, buf, n);
  }

  return 0;
}

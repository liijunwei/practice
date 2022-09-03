#include "csapp.h"

// this code is buggy

void sigint_handler1(int sig) {
  int olderrno = errno;

  if((waitpid(-1, NULL, 0)) < 0) {
    sio_error("waitpid error");
  }

  sio_puts("Handler reaped child\n");
  sleep(1);

  errno = olderrno;
}

int main(int argc, char const *argv[]) {
  int i;
  int n;
  char buf[MAXBUF]


  exit(0);
}

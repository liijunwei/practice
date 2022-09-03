#include "csapp.h"

// this code is buggy
// 这个程序是有缺陷的, 因为它假设信号是排队的

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
  char buf[MAXBUF];

  if(signal(SIGCHLD, sigint_handler1) == SIG_ERR) {
    unix_error("signal error");
  }

  for(i = 0; i < 3; i++) {
    if(Fork() == 0) {
      printf("hello from child %d\n", (int)getpid());
      exit(0);
    }
  }

  if((n = read(STDIN_FILENO, buf, sizeof(buf))) < 0) {
    unix_error("read");
  }

  printf("Parent processing input\n");

  while(1) {
    ;
  }

  exit(0);
}

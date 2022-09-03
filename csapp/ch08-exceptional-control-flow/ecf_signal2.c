#include "csapp.h"

// fix ./ecf_signal1.c
// 我们必须修改handler, 使得每次SIGCHLD处理程序被调用的时候, 回收尽可能多的僵死子进程

void sigint_handler(int sig) {
  int olderrno = errno;

  while((waitpid(-1, NULL, 0)) > 0) {
    sio_puts("Handler reaped child\n");
  }

  if(errno != ECHILD) {
    sio_error("waitpid error");
  }

  sleep(1);
  errno = olderrno;
}

int main(int argc, char const *argv[]) {
  int i;
  int n;
  char buf[MAXBUF];

  if(signal(SIGCHLD, sigint_handler) == SIG_ERR) {
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

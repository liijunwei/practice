#include "csapp.h"

volatile sig_atomic_t pid;

void sig_chld_handler(int s) {
  int olderrno = errno;
  pid = waitpid(-1, NULL, 0);
  errno = olderrno;
}

void sigint_handler(int s) {
  exit(0); // if I remove this line, ctrl+c won't interrupt the program
}

int main(int argc, char const *argv[]) {
  sigset_t mask;
  sigset_t prev;

  Signal(SIGCHLD, sig_chld_handler);
  Signal(SIGINT, sigint_handler);
  Sigemptyset(&mask);
  Sigaddset(&mask, SIGCHLD);

  while (1) {
    Sigprocmask(SIG_BLOCK, &mask, &prev); // block SIGCHLD

    if (Fork() == 0) {
      exit(0);
    }

    pid = 0;

    // wait for SIGCHLD to be received(wasteful)
    while (!pid) {
      Sigsuspend(&prev);
    }

    Sigprocmask(SIG_SETMASK, &prev, NULL); // optioonally unblock SIGCHLD

    // do some work after receiving SIGCHLD
    printf(".");
  }

  exit(0);
}

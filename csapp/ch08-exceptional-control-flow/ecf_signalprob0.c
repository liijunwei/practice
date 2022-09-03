#include "csapp.h"

// long counter = 2;
volatile long counter = 2;

void sigint_handler(int sig) {
  sigset_t mask;
  sigset_t prev_mask;

  Sigfillset(&mask);
  Sigprocmask(SIG_BLOCK, &mask, &prev_mask); // block signals
  sio_putl(--counter);
  Sigprocmask(SIG_SETMASK, &prev_mask, NULL); // restore signals

  _exit(0);
}

int main(int argc, char const *argv[]) {
  pid_t pid;
  sigset_t mask;
  sigset_t prev_mask;

  printf("%ld", counter);
  fflush(stdout);

  signal(SIGUSR1, sigint_handler);

  if((pid = Fork()) == 0) {
    while(1) {}
  }

  Kill(pid, SIGUSR1);
  Waitpid(-1, NULL, 0);

  Sigfillset(&mask);
  Sigprocmask(SIG_BLOCK, &mask, &prev_mask); // block signals
  printf("%ld", ++counter);
  Sigprocmask(SIG_SETMASK, &prev_mask, NULL); // restore signals

  exit(0);
}

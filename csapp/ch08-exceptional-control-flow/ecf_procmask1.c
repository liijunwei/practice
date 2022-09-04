#include "csapp.h"

void initjobs() {}
void addjob(int pid) {}
void deletejob(int pid) {}

void handler(int sig) {
  int olderrno = errno;
  sigset_t mask_all;
  sigset_t prev_all;
  pid_t pid;

  Sigfillset(&mask_all);
  while((pid = waitpid(-1, NULL, 0)) > 0) {
    Sigprocmask(SIG_BLOCK, &mask_all, &prev_all);
    deletejob(pid); // delete the child from the job list
    Sigprocmask(SIG_SETMASK, &prev_all, NULL);
  }

  if(errno != ECHILD) {
    sio_error("waitpid error");
  }

  _exit(0);
}

int main(int argc, char const *argv[]) {
  int pid;
  sigset_t mask_all;
  sigset_t prev_all;

  Sigfillset(&mask_all);
  Signal(SIGCHLD, handler);
  initjobs();

  while(1) {
    if((pid = Fork()) == 0) {
      Execve("/bin/date", argv, NULL);
    }

    Sigprocmask(SIG_BLOCK, &mask_all, &prev_all);
    addjob(pid);
    Sigprocmask(SIG_SETMASK, &prev_all, NULL);
  }

  exit(0);
}

#include "csapp.h"

// p543 图8-40
// fix ./ecf_procmask2.c

// unclear

void initjobs() {}
void addjob(int pid) {}

void deletejob(int pid) {}

void handler(int sig) {
  int olderrno = errno;
  sigset_t mask_all;
  sigset_t prev_all;
  pid_t pid;

  Sigfillset(&mask_all);
  while ((pid = waitpid(-1, NULL, 0)) > 0) {
    Sigprocmask(SIG_BLOCK, &mask_all, &prev_all);
    deletejob(pid); // delete the child from the job list
    Sigprocmask(SIG_SETMASK, &prev_all, NULL);
  }

  if (errno != ECHILD) {
    sio_error("waitpid error");
  }

  _exit(0);
}

// 用sigprocmask来同步进程; 这个例子中,
// 父进程保证在响应的deletejob之前执行addjob
int main(int argc, char const *argv[]) {
  int pid;
  sigset_t mask_all;
  sigset_t mask_one;
  sigset_t prev_one;

  Sigfillset(&mask_all);
  Sigemptyset(&mask_one);
  Sigaddset(&mask_one, SIGCHLD);

  Signal(SIGCHLD, handler);
  initjobs();

  while (1) {
    Sigprocmask(SIG_BLOCK, &mask_one, &prev_one); // block SIGCHLD

    if ((pid = Fork()) == 0) {
      Sigprocmask(SIG_SETMASK, &prev_one, NULL); // unblock SIGCHLD
      Execve("/bin/date", argv, NULL);
    }

    Sigprocmask(SIG_BLOCK, &mask_all, NULL); // block SIGCHLD
    addjob(pid);
    Sigprocmask(SIG_SETMASK, &prev_one, NULL); // unblock SIGCHLD
  }

  exit(0);
}

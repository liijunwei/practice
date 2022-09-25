#include "csapp.h"

// p542 图8-39
// this code is buggy

// addjob 和 deletejob之间存在竞争
// 如果deletejob赢得进展, 那么结果就是错误的
// 如果addjob赢得进展,那么结果就是正确的
// 问题: 怎么模拟这两种情况?

void initjobs() {}
void addjob(int pid) {
  // sleep(1);
}

void deletejob(int pid) {
  // sleep(1);
}

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

int main(int argc, char const *argv[]) {
  int pid;
  sigset_t mask_all;
  sigset_t prev_all;

  Sigfillset(&mask_all);
  Signal(SIGCHLD, handler);
  initjobs();

  while (1) {
    if ((pid = Fork()) == 0) {
      Execve("/bin/date", argv, NULL);
    }

    Sigprocmask(SIG_BLOCK, &mask_all, &prev_all);
    addjob(pid);
    Sigprocmask(SIG_SETMASK, &prev_all, NULL);
  }

  exit(0);
}

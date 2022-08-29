#include "csapp.h"

#define N 5

int main(int argc, char const *argv[]) {
  int status;
  int i;

  pid_t pid[N];
  pid_t retpid;

  for(i = 0; i < N; i++) {
    if((pid[i] = Fork()) == 0) {
      exit(100+i);
    }
  }

  i = 0;

  while((retpid = waitpid(pid[i++], &status, 0)) > 0) {
    if(WIFEXITED(status)) {
      printf("child %d terminated normally with exit status=%d\n", pid, WEXITSTATUS(status));
    } else {
      printf("child %d terminated abnormally\n", pid);
    }
  }

  if(errno != ECHILD) {
    unix_error("waitpid error");
  }

  exit(0);
}

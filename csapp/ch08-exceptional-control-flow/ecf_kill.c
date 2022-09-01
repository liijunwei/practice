#include "csapp.h"

int main(int argc, char const *argv[]) {
  pid_t pid;

  // child sleeps until SIGKILL signal received, then dies
  if((pid = Fork()) == 0) {
    printf("child process waiting signal to come...\n");

    Pause();
    printf("control should never reach here!\n");
    exit(0);
  }

  printf("parent sends a SIGKILL signal to a child %d\n", pid);
  sleep(10);
  Kill(pid, SIGKILL);
  exit(0);
}

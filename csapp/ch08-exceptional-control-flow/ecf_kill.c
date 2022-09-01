#include "csapp.h"

int main(int argc, char const *argv[]) {
  pid_t pid;

  // child sleeps until SIGKILL signal received, then dies
  if((pid = Fork()) == 0) {
    printf("child process waiting signal to come...\n");

    Pause(); // man pause
    printf("control should never reach here!\n");
    exit(0);
  }

  printf("parent(%d) sends a SIGKILL signal to a child(%d)\n", getpid(), pid);
  Kill(pid, SIGKILL);
  exit(0);
}

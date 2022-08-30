#include "csapp.h"

int main(int argc, char const *argv[]) {
  int status;
  pid_t pid;

  printf("hello\n");

  pid = Fork();
  printf("%d\n", !pid);

  if(pid != 0) {
    printf("a");
    fflush(stdout);
  } else {
    printf("b");
    fflush(stdout);
    waitpid(-1, NULL, 0);
  }

  printf("c");
  fflush(stdout);

  exit(0);
}

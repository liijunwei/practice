#include "csapp.h"

int main(int argc, char const *argv[]) {
  int status;
  pid_t pid;

  printf("hello\n");

  pid = Fork();
  printf("%d\n", !pid);

  if (pid != 0) {
    if (waitpid(-1, &status, 0) > 0) {
      if (WIFEXITED(status) != 0) {
        printf("%d\n", WEXITSTATUS(status));
      }
    }
  }

  printf("Bye\n");
  exit(2);
}

#include "csapp.h"

int main(int argc, char const *argv[]) {
  pid_t pid;
  int x = 1;

  pid = Fork();

  if (pid == 0) {
    printf("child: x=%d\n", ++x);
    exit(0); // early return in child process
             // if we don't exit here, child process will print `parent: x=1`
             // after child: x=2
  }

  printf("parent: x=%d\n", --x);
  exit(0);
}

// fun: fork bomb and how to preventing fork bomb on Linux
// https://www.cyberciti.biz/faq/understanding-bash-fork-bomb/

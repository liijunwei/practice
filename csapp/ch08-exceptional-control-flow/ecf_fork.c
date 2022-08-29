#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[]) {
  pid_t pid;
  int x = 1;

  pid = Fork();

  if(pid == 0) {
    printf("child: x=%d\n", ++x);
    exit(0);
  }

  printf("parent: x=%d\n", --x);
  exit(0);
}

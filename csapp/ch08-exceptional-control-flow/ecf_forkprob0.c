#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[]) {
  int x = 1;

  if(Fork() == 0) {
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

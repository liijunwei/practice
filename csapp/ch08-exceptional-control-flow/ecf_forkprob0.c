#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[]) {
  int x = 1;

  if(Fork() == 0) {
    printf("child: x=%d\n", ++x);
    exit(0);
  }

  printf("parent: x=%d\n", --x);
  exit(0);
}

#include <stdio.h>

// p613
void leak() {
  int x = (int *)Malloc(n * sizeof(int));

  return &val;
}

int main(int argc, char const *argv[]) {
  printf("stackref:       %d\n", stackref());
  printf("stackref value: %d\n", *stackref());

  return 0;
}

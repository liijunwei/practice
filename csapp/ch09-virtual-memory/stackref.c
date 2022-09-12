#include <stdio.h>

// p612
int *stackref() {
  int val;

  return &val;
}

int main(int argc, char const *argv[]) {
  printf("stackref:       %d\n", stackref());
  printf("stackref value: %d\n", *stackref());

  return 0;
}

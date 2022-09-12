#include "csapp.h"

// p613
void leak(int n) {
  int *x = (int *)Malloc(n * sizeof(int));

  // x is garbage at this point
  return;
}

int main(int argc, char const *argv[]) {
  leak(10);
  leak(20);

  return 0;
}

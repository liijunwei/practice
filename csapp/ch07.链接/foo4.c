#include <stdio.h>

void f();
int x;

int main(int argc, char const *argv[]) {
  x = 15213;
  // printf("x = %d\n", x);
  f();
  printf("x = %d\n", x);

  return 0;
}

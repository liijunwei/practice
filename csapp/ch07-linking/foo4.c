#include <stdio.h>

void f();
int x;

int main(int argc, char const *argv[]) {
  x = 15213;
  printf("before: x = %d\n", x);
  f();
  printf("after:  x = %d\n", x);

  return 0;
}

#include <stdio.h>

void f();
int x = 15213; // this value changed by f() method

int main(int argc, char const *argv[]) {
  f();
  printf("x = %d\n", x);

  return 0;
}

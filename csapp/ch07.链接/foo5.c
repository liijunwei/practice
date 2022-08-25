#include <stdio.h>

void f();

int x = 15213;
int y = 15212;

int main(int argc, char const *argv[]) {
  f();
  printf("x = 0x%x y = 0x%x \n", x, y);

  return 0;
}

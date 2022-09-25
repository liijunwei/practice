#include <stdio.h>

int main() {
  int x = 1;
  int *p = NULL;
  p = &x;

  printf("x    is %d \n", x);
  printf("*p   is %d \n", *p);
  int y = *p + 1;
  printf("y    is %d \n", y);
  printf("=============\n");

  // *p = *p + 1;
  *p += 1;
  printf("*p   is %d \n", *p);
  printf("x    is %d \n", x);

  return 0;
}
#include <stdio.h>

int main() {
  int x = 1;
  int y = 2;
  int z[10];

  int *ip = NULL;

  printf("x   is %d \n", x);
  printf("y   is %d \n", y);
  printf("z[] is %d \n", z[0]);

  ip = &x;
  printf("ip  is %d \n", ip);

  y = *ip;
  printf("y   is %d \n", y);

  *ip = 0;
  printf("x   is %d \n", x);
  printf("y   is %d \n", y);

  ip = &z[0];
  printf("ip  is %d \n", ip);

  return 0;
}
#include <stdio.h>

void swap_01(int a, int b) {
  int temp;

  temp = a;
  a = b;
  b = temp;
}

void swap_02(int *a, int *b) {
  int temp;

  temp = *a;
  *a = *b;
  *b = temp;
}

int main() {
  int x = 2;
  int y = 4;

  printf("x    is %d \n", x);
  printf("y    is %d \n", y);
  printf("======swap_01=======\n");
  swap_01(x, y);
  printf("x    is %d \n", x);
  printf("y    is %d \n", y);

  printf("======swap_02=======\n");
  swap_02(&x, &y);
  printf("x    is %d \n", x);
  printf("y    is %d \n", y);

  return 0;
}
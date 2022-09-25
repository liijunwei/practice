#include <stdio.h>

int main() {
  int a = 1;
  int *p = &a;

  printf("%lu \n", sizeof(*p));
  printf("%d \n", *p);
}

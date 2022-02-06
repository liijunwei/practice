#include <stdio.h>
#include <stdlib.h>

void foo() {
  printf("nihao\n");
  exit(999);
}

int main() {
  foo();
  printf("shijie\n");
  return 0;
}
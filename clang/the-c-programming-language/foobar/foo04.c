#include <stdio.h>
#include <stdlib.h>

void foo() {
  printf("nihao ");
  exit(999); /* 如果这里不exit, 那么会打印 nihao shijie, exit后, 只会打印 nihao
              */
}

int main() {
  foo();
  printf("shijie\n");
  return 0;
}
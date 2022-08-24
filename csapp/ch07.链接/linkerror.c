#include <stdio.h>

void foo(void);

// gcc -Wall -Og -o linkerror linkerror.c

// 编译器运行正常, 但是链接器会报错
int main(int argc, char const *argv[]) {
  foo();

  return 0;
}

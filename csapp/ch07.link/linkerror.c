#include <stdio.h>

void foo(void);

// 编译器运行正常, 但是链接器会报错
int main(int argc, char const *argv[]) {
  foo();

  return 0;
}

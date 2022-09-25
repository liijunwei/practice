#include <stdio.h>

// 标准库函数
// getchar() 一次读一个字符
// putchar() 一次写一个字符

// page 10
// 将输入复制到输出 V1
int main() {
  int c;

  c = getchar();
  while (c != EOF) {
    putchar(c);
    c = getchar();
  }
}
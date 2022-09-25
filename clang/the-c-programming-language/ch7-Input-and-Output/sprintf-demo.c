/*
page 136

sprintf执行的转换和printf一样, 但它把输出保存到一个字符串里

https://www.runoob.com/cprogramming/c-function-sprintf.html
C 库函数 int sprintf(char *str, const char *format, ...) 发送格式化输出到 str
所指向的字符串。
*/

#include <stdio.h>

int main(int argc, char const *argv[]) {
  char *s = "hello, world";

  char str_store[100];
  sprintf(str_store, "[%s]", s);

  printf("original str is: %s\n", s);
  printf("stored   str is: %s\n", str_store);

  return 0;
}

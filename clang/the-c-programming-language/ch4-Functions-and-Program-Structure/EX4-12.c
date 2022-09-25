#include <assert.h>
#include <stdio.h>
#include <string.h>

/*
page 75

用printd函数的设计思想编写一个递归版本的itoa函数
即通过递归调用把整数转换为字符串

*/

#define abs(x) ((x) < 0 ? -(x) : (x))

void itoa(int n, char s[]) {
  static int i;

  if (n / 10) {
    itoa(n / 10, s);
  } else {
    i = 0;
    if (n < 0) {
      s[i++] = '-';
    }
  }

  s[i++] = abs(n) % 10 + '0';
  s[i] = '\0';
}

int main(int argc, char const *argv[]) {
  char buffer[100];

  itoa(-2147483, buffer);
  assert(strcmp("-2147483", buffer) == 0);

  itoa(998, buffer);
  assert(strcmp("998", buffer) == 0);

  return 0;
}

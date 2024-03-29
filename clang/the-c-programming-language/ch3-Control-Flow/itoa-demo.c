#include <assert.h>
#include <stdio.h>
#include <string.h>

// page 52

// ch3-Control-Flow/reverse-demo.c
void reverse(char s[]) {
  int c;
  int i;
  int j;

  for (i = 0, j = strlen(s) - 1; i < j; i++, j--) {
    c = s[i];
    s[i] = s[j];
    s[j] = c;
  }
}

// 把数字转为字符串
// 先生成反序的字符串, 再把该字符串倒置
void itoa(int n, char s[]) {
  int i;
  int sign;

  if ((sign = n) < 0) {
    n = -n;
  }

  i = 0;
  do {
    s[i++] = n % 10 + '0';
  } while ((n /= 10) > 0);

  if (sign < 0) {
    s[i++] = '-';
  }

  s[i] = '\0';
  reverse(s);
}

int main(int argc, char const *argv[]) {
  char buffer[100];

  itoa(99, buffer);
  assert(strcmp("99", buffer) == 0);

  itoa(-100, buffer);
  assert(strcmp("-100", buffer) == 0);

  itoa(0, buffer);
  assert(strcmp("0", buffer) == 0);

  printf("PASS.\n");

  return 0;
}
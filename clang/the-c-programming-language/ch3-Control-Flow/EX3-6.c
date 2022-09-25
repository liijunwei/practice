#include <assert.h>
#include <stdio.h>
#include <string.h>

/*
page 53

修改itoa函数, 使得该函数可以接收3个参数, 其中第三个参数为最小字段宽度
为了保证转换后所得的结果至少具有第三个参数指定的最小宽度,
在必要时应在所得结果的左边填充一定的空格

*/

#define abs(x) ((x) < 0 ? -(x) : (x))

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
void itoa(int n, char s[], int w) {
  int i;
  int sign;

  sign = n;
  i = 0;

  do {
    s[i++] = abs(n % 10) + '0';
  } while ((n /= 10) != 0);

  if (sign < 0) {
    s[i++] = '-';
  }

  while (i < w) { // 必要时补充空格
    s[i++] = ' ';
  }

  s[i] = '\0';
  reverse(s);
}

int main(int argc, char const *argv[]) {
  // printf("min: %d\n", INT_MIN); // -2147483648
  char buffer[100];

  itoa(-2147483648, buffer, 11);
  assert(strcmp("-2147483648", buffer) == 0);

  itoa(998, buffer, 11);
  assert(strcmp("        998", buffer) == 0);

  itoa(2021, buffer, 11);
  assert(strcmp("       2021", buffer) == 0);

  itoa(0, buffer, 11);
  assert(strcmp("          0", buffer) == 0);

  printf("PASS.\n");
  return 0;
}

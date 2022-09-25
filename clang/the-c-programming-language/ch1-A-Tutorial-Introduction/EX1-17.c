#include <stdio.h>

// page 22

// 编写程序, 打印长度大于80个字符的所有输入行

int custom_getline(char s[], int limit);

int main() {
  int length = 0;
  int lenth_max = 1000;
  int lenth_limit = 80;
  char buffer[100];

  printf("string length > %d will be printed on screen \n\n", lenth_limit);

  while ((length = custom_getline(buffer, lenth_max)) > 0) {
    if (length > lenth_limit) {
      printf("current length is %d\n", length);
      printf("%s", buffer);
    }
  }

  return 0;
}

int custom_getline(char s[], int limit) {
  int c;
  int i;

  for (i = 0; i < limit - 1 && (c = getchar()) != EOF && c != '\n'; ++i) {
    s[i] = c;
  }

  if (c == '\n') {
    s[i] = c;
    ++i;
  }

  s[i] = '\0';

  return i;
}

#include <stdio.h>
#include <string.h>

#define ARR_SIZE 10

// page 51

// 倒置字符串s中各个字符的位置
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

int main(int argc, char const *argv[]) {
  char str[ARR_SIZE] = "abcdefg";

  printf("before: %s\n", str);
  reverse(str);
  printf("after : %s\n", str);

  return 0;
}

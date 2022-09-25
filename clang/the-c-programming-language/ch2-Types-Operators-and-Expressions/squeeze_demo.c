#include <stdio.h>

// page 37

// 从字符串s中删除字符c
void squeeze(char s[], int c) {
  int i = 0;
  int j = 0;

  for (; s[i] != '\0'; i++) {
    if (s[i] != c) {
      s[j] = s[i];
      j++;
    }
  }

  s[j] = '\0';
}

int main(int argc, char const *argv[]) {
  char str[10] = "hello";
  printf("char: %s\n", str);

  squeeze(str, 'a');
  printf("char: %s\n", str);
  squeeze(str, 'h');
  printf("char: %s\n", str);

  return 0;
}
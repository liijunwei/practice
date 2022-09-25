#include <assert.h>
#include <stdio.h>
#include <string.h>

/*
page 75

编写一个递归版本的reverse(s)函数, 以将字符串s倒置

*/

void __reverser(char s[], int i, int len) {
  int c;
  int j;

  j = len - (i + 1);

  if (i < j) {
    c = s[i];
    s[i] = s[j];
    s[j] = c;
    __reverser(s, ++i, len);
  }
}

void reverse(char s[]) { __reverser(s, 0, strlen(s)); }

#define ARR_SIZE 100

int main(int argc, char const *argv[]) {
  char str1[ARR_SIZE] = "abcdefg";
  reverse(str1);
  assert(strcmp("gfedcba", str1) == 0);

  char str2[ARR_SIZE] = "dlrowolleh";
  reverse(str2);
  assert(strcmp("helloworld", str2) == 0);

  return 0;
}

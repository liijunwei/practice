#include <assert.h>
#include <stdio.h>
#include <string.h>

/*
page 53

*/

// 删除字符串尾部的空格符/制表符和换行符
int trim(char s[]) {
  int n;

  for (n = strlen(s) - 1; n >= 0; n--) {
    if (s[n] != ' ' && s[n] != '\t' && s[n] != '\n') {
      break;
    }
  }

  s[n + 1] = '\0';

  return n;
}

int main(int argc, char const *argv[]) {
  char str1[100] = "hello";
  assert(4 == trim(str1));

  char str2[100] = "helloooo\t";
  assert(7 == trim(str2));

  char str3[100] = "";
  assert(-1 == trim(str3));

  char str4[100] = "\t";
  assert(-1 == trim(str4));

  char str5[100] = "\n\n  ";
  assert(-1 == trim(str5));

  printf("PASS.\n");

  return 0;
}
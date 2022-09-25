#include <stdio.h>

// page 48

// 统计数字/空白及其他字符
int main(int argc, char const *argv[]) {
  int c;
  int i;
  int nwhite = 0;
  int nother = 0;
  int ndight[10];

  for (i = 0; i < 10; i++) {
    ndight[i] = 0;
  }

  while ((c = getchar()) != EOF) {
    switch (c) {
    case '0':
    case '1':
    case '2':
    case '3':
    case '4':
    case '5':
    case '6':
    case '7':
    case '8':
    case '9':
      ndight[c - '0']++;
      break;
    case ' ':
    case '\n':
    case '\t':
      nwhite++;
      break;
    default:
      nother++;
      break;
    }
  }

  printf("dights =");
  for (i = 0; i < 10; i++) {
    printf(" %d", ndight[i]);
  }

  printf(", white space = %d, other = %d\n", nwhite, nother);

  return 0;
}
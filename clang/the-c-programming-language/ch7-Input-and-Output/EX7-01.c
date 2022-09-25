/*
page 135

编写一个程序, 根据他自身被调用时存放在argv[0]中的名字,
实现将大写字母转换为小写字母或者小写字母转换为大写字母的功能

按照C语言的约定, argv[0]的值是启动该程序的程序名, 因此argc的值至少为1(page 99)
*/

#include <ctype.h>
#include <stdio.h>
#include <string.h>

// bash ch7-Input-and-Output/EX7-01-test.sh
int main(int argc, char const *argv[]) {
  int c;
  const char *progname = argv[0];
  printf("progname: %s\n", progname);

  if (strcmp(progname, "./lower") == 0) {
    while ((c = getchar()) != EOF) {
      putchar(tolower(c));
    }
  } else if (strcmp(progname, "./upper") == 0) {
    while ((c = getchar()) != EOF) {
      putchar(toupper(c));
    }
  } else {
    printf("unsupported.");
    return -1;
  }

  return 0;
}

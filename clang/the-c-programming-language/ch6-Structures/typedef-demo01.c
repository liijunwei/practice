/*
page 127

typedef中声明的类型在变量名的位置出现, 而不是紧接在关键字typedef之后

*/

#include <stdio.h>
#include <stdlib.h>

typedef int Length;
typedef char *String;

int main(int argc, char const *argv[]) {
  Length a = 1;
  printf("%d\n", a);

  Length a1 = 100;
  Length a2 = 200;
  Length a3 = 300;
  Length *lenths[] = {&a1, &a2, &a3};
  printf("%d\n", *lenths[0]);
  printf("%d\n", *lenths[1]);
  printf("%d\n", *lenths[2]);

  String p;
  p = (String)malloc(100);
  p = "nihao";
  printf("%s\n", p);

  return 0;
}

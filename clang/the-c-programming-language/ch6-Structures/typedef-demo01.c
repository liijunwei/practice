/*
page 127


*/

#include <stdio.h>

typedef int Length;

int main(int argc, char const *argv[])
{
  Length a = 1;
  printf("%d\n", a);

  Length a1 = 100;
  Length a2 = 200;
  Length a3 = 300;
  Length *lenths[] = {&a1, &a2, &a3};
  printf("%d\n", *lenths[0]);
  printf("%d\n", *lenths[1]);
  printf("%d\n", *lenths[2]);

  return 0;
}

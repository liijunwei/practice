#include <stdio.h>

int foo(){
  return 'B' - 'A' + 10;
}

void *(*comp)(int *, char *, int (*fnc)());

int main(int argc, char const *argv[])
{
  printf("%d\n\n", foo());

  printf("%d\n", 'A' - 'A' + 10);
  printf("%d\n", 'B' - 'A' + 10);
  printf("%d\n", 'C' - 'A' + 10);
  printf("%d\n", 'D' - 'A' + 10);
  printf("%d\n", 'E' - 'A' + 10);
  printf("%d\n", 'F' - 'A' + 10);

  printf("%d\n", 'a' - 'a' + 10);
  printf("%d\n", 'b' - 'a' + 10);
  printf("%d\n", 'c' - 'a' + 10);
  printf("%d\n", 'd' - 'a' + 10);
  printf("%d\n", 'e' - 'a' + 10);
  printf("%d\n", 'f' - 'a' + 10);
  return 0;
}
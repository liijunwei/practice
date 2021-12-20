#include <stdio.h>

int foo(){
  int x;
  return x == (1 && x);
}

int main(int argc, char const *argv[])
{
  printf("%d\n", foo());
  return 0;
}
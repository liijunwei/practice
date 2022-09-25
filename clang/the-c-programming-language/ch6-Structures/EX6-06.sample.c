#include <stdio.h>

#define SIZE 100
#define DATE "20220203"
#define FOO 3.14
#undef FOO

int main(int argc, char const *argv[]) {
  printf("SIZE: %d\n", SIZE);
  printf("DATE: %s\n", DATE);
  // printf("FOO:  %f\n", FOO);

  return 0;
}

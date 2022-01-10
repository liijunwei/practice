#include <stdio.h>

// page 81

void swap(int *pa, int *pb){
  int temp = *pa;
  *pa = *pb;
  *pb = temp;
}

int main(int argc, char const *argv[])
{
  int a = 1;
  int b = 3;

  printf("a = %d b = %d\n", a, b);
  swap(&a, &b);

  printf("a = %d b = %d\n", a, b);
  return 0;
}

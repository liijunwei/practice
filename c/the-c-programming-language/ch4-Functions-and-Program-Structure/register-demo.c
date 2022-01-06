#include <stdio.h>

// page 71
// 寄存器

int sum(register unsigned m, register unsigned n){
  return m + n;
}


int main(int argc, char const *argv[])
{

  printf("sum is: %d\n", sum(1, 8));
  return 0;
}
#include <stdio.h>

#define CONST_LOWER 0
#define CONST_UPPER 300
#define CONST_STEP  20

int main()
{
  int fahr;

  for(fahr = CONST_LOWER; fahr <= CONST_UPPER; fahr += CONST_STEP){
    printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
  }
}


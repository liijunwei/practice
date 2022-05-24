#include <stdio.h>

int main(int argc, char const *argv[])
{
  short int v = -12345;
  unsigned short uv = (unsigned short) v;
  printf("v = %d, uv = %u\n", v, uv);

  return 0;
}


#include <stdio.h>

int main(int argc, char const *argv[])
{
  short int v = -12345;
  unsigned short uv = (unsigned short) v;
  printf("v = %d, uv = %u\n", v, uv); // v = -12345, uv = 53191

  unsigned u = 4294967295u; // UMax
  int tu = (int)u;
  printf("u = %u, tu = %d\n", u, tu); // u = 4294967295, tu = -1

  return 0;
}


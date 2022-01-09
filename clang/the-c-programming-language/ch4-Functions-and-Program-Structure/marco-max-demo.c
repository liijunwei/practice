#include <stdio.h>

// page 76

#define max(A, B) ((A) > (B) ? (A) : (B))

int main(int argc, char const *argv[])
{
  printf("max: %d\n", max(88, 1));

  return 0;
}
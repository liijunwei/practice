#include <stdio.h>

/*
page 100

V3
*/

int main(int argc, char const *argv[])
{
  while(--argc > 0){
    printf((argc > 1) ? "%s " : "%s", *++argv);
  }

  printf("\n");

  return 0;
}

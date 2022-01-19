#include <stdio.h>

/*
page 99

V2
*/

int main(int argc, char const *argv[])
{
  while(--argc > 0){
    argv++;
    char *seperator = (argc > 1) ? " " : "";
    printf("%s%s", *argv, seperator);
  }

  printf("\n");

  return 0;
}

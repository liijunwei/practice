/*
page 146

*/

#include <stdio.h>
#include <stdlib.h>

/* man 3 system */
int main(int argc, char const *argv[]) {
  printf("exit code: %d\n", system("pwd"));
  printf("exit code: %d\n", system("date"));

  return 0;
}

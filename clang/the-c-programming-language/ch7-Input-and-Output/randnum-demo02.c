/*
page 148

example from:
https://www.tutorialspoint.com/c_standard_library/c_function_srand.htm

*/

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// gcc ch7-Input-and-Output/randnum-demo02.c && watch -d -n1 "./a.out"
int main() {
  int i;
  int n = 5;
  time_t t;

  /* Intializes random number generator */
  srand((unsigned int)time(&t));

  /* Print 5 random numbers from 0 to 50 */
  for (i = 0; i < n; i++) {
    printf("%d\n", rand() % 50);
  }

  return (0);
}

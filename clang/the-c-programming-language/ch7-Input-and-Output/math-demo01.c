/*
page 148

*/

#include <stdio.h>
#include <math.h>

/* man 3 sin */
int main(int argc, char const *argv[]) {
  printf("sin(3.14/2): %f\n", sin(3.14/2));
  printf("cos(3.14/2): %f\n", cos(3.14/2));
  printf("tan(3.14/2): %f\n", tan(3.14/2));
  printf("log(3.14/2): %f\n", log(3.14/2));
  printf("sqrt(2):     %f\n", sqrt(2));
  printf("fabs(-99.1): %f\n", fabs(-99.1));

  return 0;
}

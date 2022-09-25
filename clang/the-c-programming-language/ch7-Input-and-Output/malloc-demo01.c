/*
page 147

*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* man 3 malloc */
int main(int argc, char const *argv[]) {
  int *p;
  p = (int *)malloc(sizeof(int));
  free(p);

  char *pa;
  int arr_size = 10;
  pa = (char *)calloc(arr_size, sizeof(char));
  free(pa);

  return 0;
}

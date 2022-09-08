#include "csapp.h"

int main(int argc, char const *argv[]) {
  int *array;
  int i;
  int n;

  scanf("%d", &n);
  array = (int *)Malloc(n * sizeof(int));

  for(i = 0; i < n; i++) {
    scanf("%d", &array[i]);
  }

  free(array);

  exit(0);
}

#include <stdio.h>

int sum(int *a, int n);
int sum1(int *a, int n);

int array[2] = {1, 2};

int main(int argc, char const *argv[]) {
  int val = sum(array, 2);

  printf("val = %d\n", val);
  printf("val1 = %d\n", sum1(array, 2));

  return 0;
}

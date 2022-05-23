#include <stdio.h>
#include <assert.h>

// page 38
void inplace_swap(int *x, int *y) {
  *y = *x ^ *y;
  *x = *x ^ *y;
  *y = *x ^ *y;
}

void reverse_array(int a[], int count) {
  int first;
  int last;

  for (first = 0, last = count - 1; first <= last; first++, last--) {
    inplace_swap(&a[first], &a[last]);
  }
}

void print_array(int a[], int count) {
  for (int i = 0; i < count; i ++) {
    printf("%d ", a[i]);
  }

  printf("\n");
}

int main(int argc, char const *argv[])
{
  int count = 3;
  // int a[4] = {1, 2, 3, 4};
  int a[count] = {1, 2, 3};
  print_array(a, count);
  reverse_array(a, count);
  print_array(a, count);

  return 0;
}

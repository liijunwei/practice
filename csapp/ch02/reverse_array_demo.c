#include <stdio.h>
#include <assert.h>

// page 38
void inplace_swap(int *x, int *y) {
  *y = *x ^ *y;
  *x = *x ^ *y;
  *y = *x ^ *y;
}

// swap head and tail element of an array
void reverse_array(int a[], int count) {
  int first;
  int last;

  for (first = 0, last = count - 1; first <= last; first++, last--) {
    inplace_swap(&a[first], &a[last]);
  }
}

int main(int argc, char const *argv[])
{
  int a[4] = {1, 2, 3, 4};


  return 0;
}

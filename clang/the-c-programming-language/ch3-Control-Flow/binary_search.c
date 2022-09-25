#include <assert.h>
#include <stdio.h>

// page 47

#define ARR_SIZE 6

int binsearch(int x, int v[], int n) {
  int low = 0;
  int mid;
  int high = n - 1;

  while (low <= high) {
    mid = (low + high) >> 1;

    if (x < v[mid]) {
      high = mid - 1;
    } else if (x > v[mid]) {
      low = mid + 1;
    } else {
      return mid;
    }
  }

  return -1;
}

int main(int argc, char const *argv[]) {
  int a[ARR_SIZE] = {1, 2, 4, 665, 876, 988};

  assert(0 == binsearch(1, a, ARR_SIZE));
  assert(1 == binsearch(2, a, ARR_SIZE));
  assert(2 == binsearch(4, a, ARR_SIZE));
  assert(3 == binsearch(665, a, ARR_SIZE));
  assert(4 == binsearch(876, a, ARR_SIZE));
  assert(5 == binsearch(988, a, ARR_SIZE));
  assert(-1 == binsearch(1000, a, ARR_SIZE));
  assert(-1 == binsearch(-199, a, ARR_SIZE));

  printf("PASS.\n");
  return 0;
}
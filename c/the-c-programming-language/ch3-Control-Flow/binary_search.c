#include <stdio.h>
#include <assert.h>

#define ARR_SIZE 6
int binsearch(int x, int v[], int n){

  return 100;
}

int main(int argc, char const *argv[])
{
  int a[ARR_SIZE] = {1, 2, 4, 665, 876, 988};

  assert(100 == binsearch(1, a, ARR_SIZE));
  assert(100 == binsearch(2, a, ARR_SIZE));
  assert(100 == binsearch(4, a, ARR_SIZE));
  assert(100 == binsearch(665, a, ARR_SIZE));
  assert(100 == binsearch(876, a, ARR_SIZE));
  assert(100 == binsearch(988, a, ARR_SIZE));

  printf("PASS.\n");
  return 0;
}
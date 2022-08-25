#include <stdio.h>

int x[2] = {1, 2};
int y[2] = {3, 4};
int m[2];
int n[2];

int main(int argc, char const *argv[]) {
  addvec(x, y, m, 2);
  printf("m = [%d %d]\n", m[0], m[1]);

  multvec(x, y, n, 2);
  printf("n = [%d %d]\n", n[0], n[1]);

  return 0;
}

/*
page 112


*/

#include <stdio.h>
#include <math.h>

struct point {
  int x;
  int y;
};

int main(int argc, char const *argv[])
{
  struct point pt = {300, 400};
  double dist;

  dist = sqrt((double)(pt.x) * pt.x + (double)(pt.y) * pt.y);
  printf("%f\n", dist);

  return 0;
}

/*
page 115

*/

#include <stdio.h>

struct point {
  int x;
  int y;
};

struct rect {
  struct point pt1;
  struct point pt2;
};

int main(int argc, char const *argv[]) {
  struct point pt1 = {1, 1};
  struct point pt2 = {2, 2};

  struct rect r = {pt1, pt2};
  struct rect *rp = &r;

  /* 运算符 '.' 和 '->' 都是从左至右结合的, 所以以下声明都是等价的 */
  printf("r.pt1.x:    %d\n", r.pt1.x);
  printf("r->pt1.x:   %d\n", r.pt1.x);
  printf("(r.pt1).x:  %d\n", r.pt1.x);
  printf("(r->pt1).x: %d\n", r.pt1.x);

  return 0;
}

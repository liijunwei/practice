#include <stdio.h>

struct point1 {
  int x;
  int y;
};

void print_point1(struct point1 *ps) {
  printf("x: %d y: %d sum: %d\n", ps->x, ps->y, ps->x + ps->y);
  printf("===========\n");
  printf("x: %d y: %d sum: %d\n", (*ps).x, (*ps).y, (*ps).x + (*ps).y);
}

int main(int argc, char const *argv[]) {
  struct point1 d = {100, 200};
  print_point1(&d);

  return 0;
}

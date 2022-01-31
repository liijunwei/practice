/*
page 111


*/

#include <stdio.h>

struct point1 {
  int x;
  int y;
};

struct point2 {
  int x;
  int y;
} a, b, c;

typedef struct point2 Point ;

void print_point(Point point) {
  printf("x: %d y: %d\n", point.x, point.y);
}

int main(int argc, char const *argv[])
{
  a.x = 1;
  a.y = 2;
  print_point(a);

  b.x = 6;
  b.y = 8;
  print_point(b);

  c.x = 88;
  c.y = 99;
  print_point(c);

  return 0;
}

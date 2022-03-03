/*
page 111

结构是一个或多个变量的集合, 这些变量可能为不同的类型, 为了处理的方便而将这些变量组织在一个名字下

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

void print_point1(struct point1 point) {
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

  struct point1 d = {100, 200};
  print_point1(d);

  return 0;
}

/*
page 113


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

int main(int argc, char const *argv[])
{
  struct point pt1 = {1, 1};
  struct point pt2 = {2, 2};

  struct rect screen = {pt1, pt2};

  printf("pt1.x: %d pt1.y: %d\n", screen.pt1.x, screen.pt1.y);
  printf("pt2.x: %d pt2.y: %d\n", screen.pt2.x, screen.pt2.y);

  return 0;
}

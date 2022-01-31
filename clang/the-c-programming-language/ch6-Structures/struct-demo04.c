/*
page 113


*/

#include <stdio.h>

#define XMAX 10
#define YMAX 10

struct point {
  int x;
  int y;
};

struct rect {
  struct point pt1;
  struct point pt2;
};

struct point makepoint(int x, int y) {
  struct point temp;

  temp.x = x;
  temp.y = y;

  return temp;
}

// 将两个点相加
struct point addpoint(struct point p1, struct point p2) {
  p1.x += p2.x;
  p1.y += p2.y;

  return p1;
}

// 如果点p在矩形r内, 返回1, 否则返回0
int ptinrect(struct point p, struct rect r) {
  return p.x >= r.pt1.x && p.x < r.pt2.x
      && p.y >= r.pt1.y && p.y < r.pt2.y;
}

#define min(a, b) ((a) < (b) ? (a): (b))
#define max(a, b) ((a) > (b) ? (a): (b))

// 将矩形坐标规范化
struct rect canonrect(struct rect r) {
  struct rect temp;

  temp.pt1.x = min(r.pt1.x, r.pt2.x);
  temp.pt1.y = min(r.pt1.y, r.pt2.y);
  temp.pt2.x = max(r.pt1.x, r.pt2.x);
  temp.pt2.y = max(r.pt1.y, r.pt2.y);

  return temp;
}

int main(int argc, char const *argv[])
{
  struct point pt1;
  struct point pt2;

  struct rect screen = {pt1, pt2};
  screen.pt1 = makepoint(10, 10);
  screen.pt2 = makepoint(6, 6);

  struct rect new_screen = canonrect(screen);
  printf("pt1.x: %d pt1.y: %d\n", new_screen.pt1.x, new_screen.pt1.y);
  printf("pt2.x: %d pt2.y: %d\n", new_screen.pt2.x, new_screen.pt2.y);

  struct point middle;
  middle = makepoint((new_screen.pt1.x + new_screen.pt2.x) / 2, (new_screen.pt1.y + new_screen.pt2.y) / 2);
  printf("middle.x: %d middle.y: %d\n", middle.x, middle.y);

  return 0;
}

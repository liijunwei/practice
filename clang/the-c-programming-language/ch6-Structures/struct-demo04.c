/*
page 113

结构的合法操作只有几种
    + 作为一个整体复制和赋值
    + 通过&取地址符取地址
    + 访问其成员

为了进一步地理解结构, 下面编写几个对点和矩形进行操作的函数;
至少可以通过3种可能的方法传递结构(各有利弊):
    + 分别传递各个结构成员
    + 传递整个结构
    + 传递指向结构的指针

如果传递给函数的结构很大, 使用指针方式的效率通常比复制整个结构的效率要高
结构指针类似于普通变量指针: `struct point *pp`

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

#define min(a, b) ((a) < (b) ? (a) : (b))
#define max(a, b) ((a) > (b) ? (a) : (b))

// 将矩形坐标规范化
struct rect canonrect(struct rect r) {
  struct rect temp;

  temp.pt1.x = min(r.pt1.x, r.pt2.x);
  temp.pt1.y = min(r.pt1.y, r.pt2.y);
  temp.pt2.x = max(r.pt1.x, r.pt2.x);
  temp.pt2.y = max(r.pt1.y, r.pt2.y);

  return temp;
}

int main(int argc, char const *argv[]) {
  struct point pt1;
  struct point pt2;

  struct rect screen = {pt1, pt2};
  screen.pt1 = makepoint(10, 10);
  screen.pt2 = makepoint(6, 6);

  struct rect new_screen = canonrect(screen);
  printf("pt1.x: %d pt1.y: %d\n", new_screen.pt1.x, new_screen.pt1.y);
  printf("pt2.x: %d pt2.y: %d\n", new_screen.pt2.x, new_screen.pt2.y);

  struct point middle;
  middle = makepoint((new_screen.pt1.x + new_screen.pt2.x) / 2,
                     (new_screen.pt1.y + new_screen.pt2.y) / 2);
  printf("middle.x: %d middle.y: %d\n", middle.x, middle.y);

  return 0;
}

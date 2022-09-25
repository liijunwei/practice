/*
page 115

*/

#include <stdio.h>

struct point {
  int x;
  int y;
};

struct point makepoint(int x, int y) {
  struct point temp;

  temp.x = x;
  temp.y = y;

  return temp;
}

int main(int argc, char const *argv[]) {
  struct point origin = makepoint(99, 81);
  struct point *pp;

  pp = &origin;
  /* (*pp).x 里的括号是必须的, 因为运算符'.'的优先级高于'*' */
  printf("origin is (%d, %d)(V1)\n", (*pp).x, (*pp).y);

  /* 结构指针的使用频度非常高, 为了使用方便, C语言提供了另一种简写方式 */
  /* 假定p是一个执行结构的指针, 则 p->结构成员 <=> (*p).结构成员       */
  printf("origin is (%d, %d)(V2)\n", pp->x, pp->y);

  return 0;
}

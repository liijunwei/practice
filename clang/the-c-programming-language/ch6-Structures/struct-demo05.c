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

int main(int argc, char const *argv[])
{
  struct point origin = makepoint(99, 81);
  struct point *pp;

  pp = &origin;
  printf("origin is (%d, %d)\n", (*pp).x, (*pp).y); /* (*pp).x 里的括号是必须的, 因为运算符'.'的优先级高于'*' */

  return 0;
}

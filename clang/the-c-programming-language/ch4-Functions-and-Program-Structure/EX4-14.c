#include <stdio.h>

/*
page 77

定义宏swap(t, x, y)以交换t类型的两个参数(使用程序块结构会对你有所帮助)
*/

#define swap(t, x, y) { t _tmp; _tmp = x; x = y; y = _tmp;}

int main(int argc, char const *argv[])
{
  int a[2] = {1, 9};
  printf("before: a[0] = %d a[1] = %d\n", a[0], a[1]);

  swap(int, a[0], a[1]);

  printf("before: a[0] = %d a[1] = %d\n", a[0], a[1]);

  return 0;
}

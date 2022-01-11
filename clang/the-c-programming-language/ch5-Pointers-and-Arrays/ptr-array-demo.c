#include <stdio.h>
#include <assert.h>

/*
page 84

理解指针变量的加减运算

一个通过数组和下标实现的表达式, 可等价地通过指针和偏移量实现
*/

#define ARR_SIZE 10

int main(int argc, char const *argv[])
{
  int *pa;
  int a[ARR_SIZE] = {1, 2, 3, 5};

  pa = &a[0]; // <=> pa = a;

  // a[i]  <=> *(pa + i)
  // &a[i] <=> pa + i
  // pa[i] <=> *(pa + i)
  assert(a[0] == *(pa + 0));
  assert(a[1] == *(pa + 1));
  assert(a[2] == *(pa + 2));
  assert(a[3] == *(pa + 3));

  // assert(2 == *(pa + 3));

  return 0;
}


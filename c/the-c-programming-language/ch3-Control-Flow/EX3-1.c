#include <stdio.h>
#include <assert.h>

/*
page 47

ch3-Control-Flow/binary_search.c 的例子里, while循环语句内共执行了两次测试, 其实只要一次就够了(代价是将更多的测试在循环外执行)
1. 重写该函数, 使得子啊循环内部只执行一次
2. 比较两个函数的运行时间(how?)

"how to do benchmark for c program?"
https://stackoverflow.com/questions/2349776/how-can-i-benchmark-c-code-easily
*/

#define ARR_SIZE 6

int binsearch_v1(int x, int v[], int n){
  int low = 0;
  int mid;
  int high = n -1;

  while(low <= high){
    mid = (low + high) >> 1;

    if(x < v[mid]){
      high = mid - 1;
    } else if(x > v[mid]){
      low = mid + 1;
    } else {
      return mid;
    }
  }

  return -1;
}

int binsearch_v2(int x, int v[], int n){
  int low = 0;
  int high = n -1;
  int mid = (low + high) >> 1;

  while(low <= high && x != v[mid]){
    if(x < v[mid]){
      high = mid - 1;
    } else {
      low = mid + 1;
    }
    mid = (low + high) >> 1;
  }

  return (x == v[mid]) ? mid : -1;
}

int main(int argc, char const *argv[])
{
  int a[ARR_SIZE] = {1, 2, 4, 665, 876, 988};

  assert(0 == binsearch_v2(1, a, ARR_SIZE));
  assert(1 == binsearch_v2(2, a, ARR_SIZE));
  assert(2 == binsearch_v2(4, a, ARR_SIZE));
  assert(3 == binsearch_v2(665, a, ARR_SIZE));
  assert(4 == binsearch_v2(876, a, ARR_SIZE));
  assert(5 == binsearch_v2(988, a, ARR_SIZE));
  assert(-1 == binsearch_v2(1000, a, ARR_SIZE));
  assert(-1 == binsearch_v2(-199, a, ARR_SIZE));
  printf("PASS.\n");

  return 0;
}
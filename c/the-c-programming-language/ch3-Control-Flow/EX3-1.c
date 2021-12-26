#include <stdio.h>
#include <assert.h>
#include <time.h>

/*
page 47

ch3-Control-Flow/binary_search.c 的例子里, while循环语句内共执行了两次测试, 其实只要一次就够了(代价是将更多的测试在循环外执行)
1. 重写该函数, 使得子啊循环内部只执行一次
2. 比较两个函数的运行时间(how?)

"how to do benchmark for c program?"
https://stackoverflow.com/questions/2349776/how-can-i-benchmark-c-code-easily

benchmark result

lijunwei@bxzy:the-c-programming-language(main)$ uname -a
Darwin bxzy 20.6.0 Darwin Kernel Version 20.6.0: Tue Oct 12 18:33:42 PDT 2021; root:xnu-7195.141.8~1/RELEASE_X86_64 x86_64
lijunwei@bxzy:the-c-programming-language(main)$ gcc ch3-Control-Flow/EX3-1.c && ./a.out
binsearch_v1 costs: 0.547102
binsearch_v2 costs: 0.657599

root@sanchez:~$ uname -a
Linux sanchez 3.10.0-1160.25.1.el7.x86_64 #1 SMP Wed Apr 28 21:49:45 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux
root@sanchez:~$ gcc ex3-1.c && ./a.out
binsearch_v1 costs: 0.781838
binsearch_v2 costs: 0.844247

说明减少一次循环没啥意义嘛?

又试了一下 修改 get_mid 的实现, 改为位计算 会快很多
*/

#define ARR_SIZE 6
#define TEST_ROUNDS 10000000

int get_mid(int l, int h){
  return (l + h) >> 1;
  // return (l + h) / 2;
}

int binsearch_v1(int x, int v[], int n){
  int low = 0;
  int mid;
  int high = n -1;

  while(low <= high){
    mid = get_mid(low, high);

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
  int high = n - 1;
  int mid = get_mid(low, high);

  while(low <= high && x != v[mid]){
    if(x < v[mid]){
      high = mid - 1;
    } else {
      low = mid + 1;
    }
    mid = get_mid(low, high);
  }

  return (x == v[mid]) ? mid : -1;
}

int main(int argc, char const *argv[])
{
  int a[ARR_SIZE] = {1, 2, 4, 665, 876, 988};
  float start_time;
  float end_time;
  float time_elapsed;

  start_time = (float)clock()/CLOCKS_PER_SEC;
  for(int i = 0; i < TEST_ROUNDS; i++){
    assert(0 == binsearch_v1(1, a, ARR_SIZE));
    assert(1 == binsearch_v1(2, a, ARR_SIZE));
    assert(2 == binsearch_v1(4, a, ARR_SIZE));
    assert(3 == binsearch_v1(665, a, ARR_SIZE));
    assert(4 == binsearch_v1(876, a, ARR_SIZE));
    assert(5 == binsearch_v1(988, a, ARR_SIZE));
    assert(-1 == binsearch_v1(1000, a, ARR_SIZE));
    assert(-1 == binsearch_v1(-199, a, ARR_SIZE));
    // printf("PASS.\n");
  }
  end_time = (float)clock()/CLOCKS_PER_SEC;
  time_elapsed = end_time - start_time;
  printf("binsearch_v1 costs: %f\n", time_elapsed);

  start_time = (float)clock()/CLOCKS_PER_SEC;
  for(int i = 0; i < TEST_ROUNDS; i++){
    assert(0 == binsearch_v2(1, a, ARR_SIZE));
    assert(1 == binsearch_v2(2, a, ARR_SIZE));
    assert(2 == binsearch_v2(4, a, ARR_SIZE));
    assert(3 == binsearch_v2(665, a, ARR_SIZE));
    assert(4 == binsearch_v2(876, a, ARR_SIZE));
    assert(5 == binsearch_v2(988, a, ARR_SIZE));
    assert(-1 == binsearch_v2(1000, a, ARR_SIZE));
    assert(-1 == binsearch_v2(-199, a, ARR_SIZE));
    // printf("PASS.\n");
  }
  end_time = (float)clock()/CLOCKS_PER_SEC;
  time_elapsed = end_time - start_time;
  printf("binsearch_v2 costs: %f\n", time_elapsed);

  return 0;
}
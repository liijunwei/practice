#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=93

// 1. 指针变量做函数的参数
// 2. 指针变量做函数的返回值

void my_rank_impl(int *q1, int *q2){
  int temp;
  if(*q1 < *q2){
    temp = *q1;
    *q1 = *q2;
    *q2 = temp;
  }
}

int main(){
  int a;
  int b;
  int *p1;
  int *p2;

  cin >> a >> b;

  p1 = &a;
  p2 = &b;

  my_rank_impl(p1, p2);

  cout << a << " " << b << endl;


  cout << endl;
  return 0;
}

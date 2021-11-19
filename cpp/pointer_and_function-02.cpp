#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=93

// 数组名作为函数参数

// 数组名相当于指向数组第一个元素的指针
// 数组名相当于指向数组第一个元素的指针

void sum_v1(int *p, int n){
  int total = 0;
  for(int i = 0; i < n; i++){
    total += *p++;
  }

  cout << __LINE__ << " Total: " << total << endl;
}

void sum_v2(int *p, int n){
  int total = 0;
  for(int i = 0; i < n; i++){
    total += *p;
    p++;
  }

  cout << __LINE__ << " Total: " << total << endl;
}

void sum_v3(int *p, int n){
  int total = 0;
  for(int i = 0; i < n; i++, p++){
    total += *p;
  }

  cout << __LINE__ << " Total: " << total << endl;
}

int main(){
  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

  sum_v1(a, 10);
  sum_v2(a, 10);
  sum_v3(a, 10);

  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=95

// 两种定义方式等价
int *get_v1(int array[][4], int n, int m);
int *get_v2(int (*array)[4], int n, int m);

// TODO 又懵了, 传入的 (*array)[4] 是什么含义, 怎么使用

// 指针 作为函数的返回值
int main(){
  int a[4][4] = {
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 10, 11, 12},
    {13, 14, 15, 16}
  };

  int *pv1;
  int *pv2;

  pv1 = get_v1(a, 2, 3);
  cout << "pv1 " << pv1 << endl;
  cout << endl;
  pv2 = get_v1(a, 2, 3);
  cout << "pv2 " << pv2 << endl;
  cout << endl;

  cout << endl;
  return 0;
}

int *get_v1(int array[][4], int n, int m){
  int *pt;
  pt = *(array + n - 1) + m - 1;

  return pt;
}

int *get_v2(int (*array)[4], int n, int m){
  int *pt;
  pt = *(array + n - 1) + m - 1;

  return pt;
}
#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=95

// 指针 作为函数的返回值

int *getInt();

int main(){
  int *p;
  p = getInt();

  cout << *p << endl; // p 指向了一片已经被释放了的地址!!!

  cout << endl;
  return 0;
}

int *getInt(){
  int value1 = 20;
  return &value1; // warning: address of stack memory associated with local variable 'value1' returned [-Wreturn-stack-address]
}
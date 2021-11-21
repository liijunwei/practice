#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=87
int main(){
  int a[5] = {10, 11, 12, 13, 14};
  int *p = NULL;
  int *q = NULL;

  // 数组名相当于指向数组第一个元素的指针, 即 a 指向 a[0], 变量a里存的是 a[0] 的地址
  cout << "a    " << a << endl;
  // 注意: a是常量, p是变量; 所以 a++ 无意义, 但是p++有意义
  p = a;
  q = a;

  cout << "     " << q[3] << endl;
  cout << "     " << q[4] << endl;

  cout << "p    " << p    << endl;
  cout << "*p   " << *p   << endl;
  cout << "     " << *p++ << endl;
  cout << "     " << *p++ << endl;
  cout << "     " << *++q << endl;
  cout << "     " << *++q << endl;

  cout << endl;
  return 0;
}

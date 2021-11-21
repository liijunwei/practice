#include <iostream>

using namespace std;

/*

*/

// https://www.bilibili.com/video/BV1bs41197KN?p=87
int main(){
  int a[5] = {10, 11, 12, 13, 14};
  int *p = NULL;

  // 数组名相当于指向数组第一个元素的指针, 即 a 指向 a[0], 变量a里存的是 a[0] 的地址
  cout << "a    " << a << endl;

  p = a;

  cout << "p    " << p    << endl;
  cout << "*p   " << *p   << endl;
  cout << "     " << *p++ << endl;
  cout << "     " << *p++ << endl;

  cout << endl;
  return 0;
}

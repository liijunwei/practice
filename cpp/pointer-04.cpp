#include <iostream>

using namespace std;

// 用指针变量访问数组元素
int main(){

  int a[5] = {10, 11, 12, 13, 14};

  int *p = NULL;
  p = a;

  cout << "a:    " << a << endl;
  cout << "p:    " << p << endl;
  cout << "*p:   " << *p << endl;
  cout << "*p++: " << *p++ << endl; // 先试用p, 再 ++
  cout << "*p++: " << *p++ << endl;
  cout << "*p:   " << *p << endl;

  cout << endl;
  return 0;
}

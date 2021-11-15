#include <iostream>

using namespace std;

// 用指针变量访问数组元素
int main(){

  int a[5] = {10, 11, 12, 13, 14};

  int *p = NULL;
  cout << "a:    " << a << endl;

  p = a;

  cout << "p:    " << p << endl;
  cout << "*p:   " << *p << endl;
  cout << "*p++: " << *p++ << endl;
  cout << "*p++: " << *p++ << endl;

  cout << endl;
  return 0;
}

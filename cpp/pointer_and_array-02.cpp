#include <iostream>

using namespace std;

/*
1. 数组名代表数组首元素的地址(数组名相当于指向数组第一个元素的指针)
2. 数组名a不是变量, 不能给a赋值

*/

// https://www.bilibili.com/video/BV1bs41197KN?p=86&spm_id_from=pageDriver
int main(){
  int a[5] = {10, 11, 12, 13, 14};
  cout << "a     " << a << endl;
  cout << "*a    " << *a << endl;
  cout << "&a[0] " << &a[0] << endl;
  cout << "a[0]  " << a[0] << endl;

  cout << endl;
  return 0;
}

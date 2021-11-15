#include <iostream>

using namespace std;

// 用指针变量访问数组元素
int main(){

  int a[5] = {10, 11, 12, 13, 14};

  for(int i = 0; i < sizeof(a) / sizeof(a[0]); i++){
    cout << a[i] << " ";
  }

  cout << endl;
  cout << endl;

  int *pointer = a;

  cout << "pointer:        " << pointer << endl;
  cout << "&a[0]:          " << &a[0] << endl;

  cout << "(pointer + 1)   " << (pointer + 1) << endl;
  cout << "(pointer + 2)   " << (pointer + 2) << endl;
  cout << "(pointer + 3)   " << (pointer + 3) << endl;

  cout << "*(pointer + 1)  " << *(pointer + 1) << endl;
  cout << "*(pointer + 2)  " << *(pointer + 2) << endl;
  cout << "*(pointer + 3)  " << *(pointer + 3) << endl;

  cout << "pointer[1]      " << *(pointer + 1) << endl;
  cout << "pointer[2]      " << *(pointer + 2) << endl;
  cout << "pointer[3]      " << *(pointer + 3) << endl;

  cout << endl;
  return 0;
}

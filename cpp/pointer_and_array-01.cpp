#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=86&spm_id_from=pageDriver
int main(){
  int a[5] = {1, 2, 3, 4, 5};

  int *p = &a[3];

  cout << "a " << a << endl;
  cout << "p " << p << endl;
  cout << "*p " << *p << endl;
  cout << "*p++ " << *p++ << endl;
  cout << "*p " << *p << endl;

  *p = 100;
  cout << a[3] << endl;

  cout << endl;
  return 0;
}

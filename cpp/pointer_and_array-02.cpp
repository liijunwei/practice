#include <iostream>

using namespace std;

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

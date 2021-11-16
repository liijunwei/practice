#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=87
int main(){
  int a[5] = {10, 11, 12, 13, 14};

  cout << "a[0] " << a[0] << endl;
  cout << "a[1] " << a[1] << endl;
  cout << "a[2] " << a[2] << endl;
  cout << "a[3] " << a[3] << endl;
  cout << "a[4] " << a[4] << endl;
  cout << endl;

  int *p = NULL;
  p = a;

  cout << "p[0] " << p[0] << endl;
  cout << "p[1] " << p[1] << endl;
  cout << "p[2] " << p[2] << endl;
  cout << "p[3] " << p[3] << endl;
  cout << "p[4] " << p[4] << endl;

  cout << endl;
  return 0;
}

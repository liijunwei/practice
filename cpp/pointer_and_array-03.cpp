#include <iostream>

using namespace std;

/*

*/

// https://www.bilibili.com/video/BV1bs41197KN?p=87
int main(){
  int a[5] = {10, 11, 12, 13, 14};
  int *p = NULL;

  cout << "a    " << a << endl;

  p = a;

  cout << "p    " << p    << endl;
  cout << "*p   " << *p   << endl;
  cout << "*p++ " << *p++ << endl;
  cout << "*p++ " << *p++ << endl;

  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;

int main(){
  int a[5] = {10, 11, 12, 13, 14};
  int *p = &a[3];
  *p = 100;

  cout << *p++ << endl; // 100
  cout << *p-- << endl; // 14
  cout << *--p << endl; // 12

  cout << endl;
  return 0;
}

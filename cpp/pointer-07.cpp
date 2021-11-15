#include <iostream>
#include <iomanip>

using namespace std;

int main(){
  int a[5] = {1, 2, 3, 4, 5};
  int *p = &a[3];

  *p = 100;

  cout << "*p++ " << *p++ << endl; // 100
  cout << "*p   " << *p << endl;   // 5
  cout << "*p-- " << *p-- << endl; // 5
  cout << "*--p " << *--p << endl; // 3

  cout << endl;
  return 0;
}


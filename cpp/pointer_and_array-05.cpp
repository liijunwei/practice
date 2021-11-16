#include <iostream>

using namespace std;

int main(){
  int a[5] = {10, 11, 12, 13, 14};

  for(int i = 0; i < 5; i++){
    cout << i << " " << a[i] << endl;
  }
  cout << endl;

  int *p = &a[3];

  cout << *p++ << endl; // 13
  cout << *p-- << endl; // 14
  cout << *--p << endl; // 12

  cout << endl;
  return 0;
}

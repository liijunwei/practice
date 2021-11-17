#include <iostream>

using namespace std;

int main(){
  int a[5] = {10, 11, 12, 13, 14};
  // int a[3] = {10, 11, 12};

  cout << __LINE__ << " " << a << endl;
  cout << __LINE__ << " " << (a + 1) << endl;
  cout << __LINE__ << " " << &a << endl;       // 指向      整个数组空间
  cout << __LINE__ << " " << (&a + 1) << endl; // 指向 下个 整个数组空间
  cout << __LINE__ << " " << *(&a) << endl;
  cout << __LINE__ << " " << (*(&a) + 1) << endl;

  cout << endl;
  return 0;
}

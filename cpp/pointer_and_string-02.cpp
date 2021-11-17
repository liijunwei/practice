#include <iostream>

using namespace std;

int main(){
  int a = 5;
  int *pa = &a;

  int b[6] = {1, 2, 3, 4, 5, 6};
  int *pb = b;

  char c[6] = {'h', 'e', 'l', 'l', 'o', '\0'};
  char *pc = c;

  cout << a << endl;
  cout << pa << endl;
  cout << endl;

  cout << b << endl;
  cout << pb << endl;
  cout << endl;

  cout << c << endl;
  cout << pc << endl;
  cout << endl;

  cout << static_cast<void*>(c) << endl;
  cout << static_cast<void*>(pc) << endl;
  cout << static_cast<void*>(++pc) << endl;
  cout << static_cast<void*>(++pc) << endl;
  cout << static_cast<void*>(++pc) << endl;
  cout << endl;

  cout << endl;
  return 0;
}

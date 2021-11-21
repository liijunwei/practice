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

  char d[100] = "nihao xiaoli";
  char *pd = d;
  cout << d << endl;
  cout << pd << endl;

  cout << static_cast<void*>(d) << endl;
  cout << static_cast<void*>(pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << static_cast<void*>(++pd) << endl;
  cout << endl;

  char e[100] = "nihao xiaoli";
  char *pe = e;
  cout << e[0] << endl;
  cout << e[1] << endl;
  cout << e[2] << endl;
  cout << pe[0] << endl;
  cout << pe[1] << endl;
  cout << pe[2] << endl;

  cout << endl;
  return 0;
}

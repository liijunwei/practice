#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=85
int main(){
  int a[5] = {1, 2, 3, 4, 5};
  double c[5] = {1.111, 2.111, 3.111, 4.111, 5.111};

  int *pa = a;
  double *pc = c;

  cout << "pa " << pa++ << endl;
  cout << "pa " << pa++ << endl;
  cout << "pa " << pa++ << endl;
  cout << endl;

  cout << "pc " << pc++ << endl;
  cout << "pc " << pc++ << endl;
  cout << "pc " << pc++ << endl;

  cout << endl;
  return 0;
}


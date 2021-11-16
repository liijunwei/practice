#include <iostream>

using namespace std;

int main(){
  int a[5] = {1, 2, 3, 4, 5};
  char c[5] = {'a', 'b', 'c', 'd', 'e'};

  int *pa = a;
  char *pc = c;

  cout << "pa " << pa++ << endl;
  cout << "pa " << pa++ << endl;
  cout << "pa " << pa++ << endl;

  cout << "pc " << pc++ << endl;
  cout << "pc " << pc++ << endl;
  cout << "pc " << pc++ << endl;

  cout << endl;
  return 0;
}


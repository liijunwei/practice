#include <iostream>

using namespace std;

void exchange(int *m, int *n);

int main(){
  int m = 10;
  int n = 100;

  cout << "Before exchange value: " << m << " " << n << endl;
  exchange(&m, &n);
  cout << "After exchange value: " << m << " " << n << endl;
  cout << endl;

  return 0;
}

void exchange(int *m, int *n){
  int p = *m;
  *m = *n;
  *n = p;
}

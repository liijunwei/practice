#include <iostream>

using namespace std;

void exchange(int a[]);

int main(){
  int a[2] = {3, 5};
  cout << a[0] << " " << a[1] << endl;

  exchange(a);

  cout << a[0] << " " << a[1] << endl;
  cout << endl;

  return 0;
}

void exchange(int a[]){
  a[0] = 30;
  a[1] = 50;
}

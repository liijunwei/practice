#include <iostream>

using namespace std;

int main(){
  int a = 0;
  int b = 0;
  int temp;

  int *p1 = NULL;
  int *p2 = NULL;

  cin >> a >> b;
  p1 = &a;
  p2 = &b;

  if(*p1 < *p2){
    temp = *p1;
    *p1 = *p2;
    *p2 = temp;
  }

  cout << "Max = " << *p1 << ", min = " << *p2 << endl;

  cout << endl;
  return 0;
}

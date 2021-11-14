#include <iostream>

using namespace std;

int a = 0;
int b = 0;

void exchange();

int main(){
  cin >> a >> b;

  exchange();

  cout << a << " " << b << endl;
  cout << endl;

  return 0;
}

void exchange(){
  int p;

  if(a > b){
    p = a;
    a = b;
    b = p;
  }
}

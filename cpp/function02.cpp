#include <iostream>

using namespace std;

int max(int m, int n);

int main(){
  int m = 10;
  int n = 100;

  cout << "The max value of " << m << " and " << n << " is " << max(m, n) << endl;

  return 0;
}

int max(int m, int n){
  return m > n ? m : n;
}

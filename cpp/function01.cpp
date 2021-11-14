#include <iostream>

using namespace std;

// g++ ./function01.cpp && ./a.out
int absolute(int n){
  if(n < 0){
    return (-n);
  }
  else{
    return n;
  }
}

int main(){
  int m = -123;
  int result = 0;

  result = absolute(m);
  cout << "The absolute value of " << m << " is " << result << endl;

  return 0;
}

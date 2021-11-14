#include <iostream>

using namespace std;

// g++ ./function01.cpp && ./a.out
// cat "$HOME/Library/Application Support/Sublime Text 3/Packages/User/CPP-Custom.sublime-build"
// cmd+shift+b

int absolute(int n);

int main(){
  int m = -123;
  int result = 0;

  result = absolute(m);
  cout << "The absolute value of " << m << " is " << result << endl;

  return 0;
}

int absolute(int n){
  if(n < 0){
    return (-n);
  }
  else{
    return n;
  }
}

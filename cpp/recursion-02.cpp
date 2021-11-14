#include <iostream>

using namespace std;

int recur();

int main(){
  recur();
  return 0;
}

int recur(){
  char c;

  c = cin.get();

  if(c != '\n'){
    recur();
  }

  cout << c;

  return 0;
}
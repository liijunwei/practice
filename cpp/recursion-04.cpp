#include <iostream>

using namespace std;

int fab_seq(int n);

int main(){
  cout << "fab_seq(40): " << fab_seq(40) << endl;

  cout << endl;
  return 0;
}

int fab_seq(int n){
  if(n == 1){
    return 1;
  }
  else if(n == 2){
    return 1;
  }
  else{
    return (fab_seq(n - 1) + fab_seq(n - 2));
  }
}
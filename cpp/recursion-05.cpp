#include <iostream>

using namespace std;

// 二进制转换

void convert_to_binary(int n);

int main(){
  convert_to_binary(40);
  cout << endl;
  convert_to_binary(61);
  cout << endl;
  return 0;
}

void convert_to_binary(int n){
  if((n / 2) != 0){
    convert_to_binary(n / 2);
    cout << (n % 2);
  }
  else{
    cout << n;
  }
}

#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=76
// "函数可以嵌套调用, 但不能嵌套地定义"
int fact(int n);

int main(){
  cout << "fact(4) " << fact(4) << endl;

  cout <<  endl;
  return 0;
}

int fact(int n){
  if(n == 1){
    return 1;
  }
  else {
    return fact(n - 1) * n;
  };
}
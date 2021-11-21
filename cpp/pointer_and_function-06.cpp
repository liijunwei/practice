#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=94

void demo_withtou_const(){
  int a = 256;
  int *p = &a;
  *p = 257;
  cout << *p << endl;
}

void demo_const(){
  int a = 256;
  const int *p = &a;
  // *p = 257;
  cout << *p << endl;
}

int main(){
  demo_withtou_const();
  demo_const();

  cout << endl;
  return 0;
}

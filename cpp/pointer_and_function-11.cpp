#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=95

int *getIntV1();
int *getIntV2();

int main(){
  int *p;
  int *q;

  p = getIntV1();
  q = getIntV2();

  cout << *q << endl;

  cout << endl;
  return 0;
}

int *getIntV1(){
  int value1 = 20;
  return &value1;
}

int *getIntV2(){
  int value1 = 30;
  return &value1;
}


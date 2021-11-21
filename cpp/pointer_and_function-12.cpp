#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=95

// Q: 只能返回全局变量的地址吗?
// A: 当然不是, 接下来学习 static 关键字
int value1 = 20;
int value2 = 30;

// 返回一个 处于生命周期中的变量的地址
int *getIntV1();
int *getIntV2();

int main(){
  int *p;
  int *q;

  p = getIntV1();
  q = getIntV2();

  cout << "*p " << *p << endl;
  cout << "*q " << *q << endl;

  cout << endl;
  return 0;
}

int *getIntV1(){
  return &value1;
}

int *getIntV2(){
  return &value2;
}


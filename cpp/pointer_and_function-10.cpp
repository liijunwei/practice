#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=95

int *getInt();

// TODO 问题: 为什么运行了 1000000000次, 得到的值都是20?
int main(){
  int *p;
  int demo = 0;

  for(int i = 0; i < 1000000000; i++){
    p = getInt();
    demo = *p;
    if(demo != 20){
      cout << demo << endl;
    }
  }

  cout << endl;
  return 0;
}

int *getInt(){
  int value1 = 20;
  return &value1; // warning: address of stack memory associated with local variable 'value1' returned [-Wreturn-stack-address]
}
#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=96
/*
使用 static 关键字, 定义 静态局部变量

定义: 函数中的局部变量的值, 在函数调用结束后不消失而保留原值
      即其占用的存储单元不释放, 在下一次该函数调用时, 仍可以继续使用该变量

*/

void incr();

int main(){
  for(int i = 0; i < 5; i++){
    incr();
    cout << "Call Again!" << endl;
  }

  cout << endl;
  return 0;
}

void incr(){
  int a = 0; // 动态变量
  static int b = 0; // 静态变量, 只初始化一次

  a += 1;
  b += 1;

  cout << "a = " << a << endl;
  cout << "b = " << b << endl;
}
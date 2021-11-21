#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=94

int main(){
  const int a = 78;
  // int a = 78;
  const int b = 28;
  int c = 18;

  const int *pi = &a;
  // *pi = 58; // 会报错, 不能对*pi进行重新赋值
            // "read-only variable is not assignable"

  pi = &b; // 可以让pi指向新的变量

  // *pi = 68; // 会报错: "error: read-only variable is not assignable"

  pi = &c;
  // *pi = 88; // 会报错: "error: read-only variable is not assignable"

  cout << endl;
  return 0;
}

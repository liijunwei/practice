#include <iostream>

using namespace std;

/*
1. 什么是指针(是地址)
2. 什么是指针变量(是变量, 里面存放的是地址)
3. *pointer 的含义(指针变量指向的某个内存空间里的内容)
4. pointer++ 的含义(体现了指针变量的 基类型 的作用)
*/

void demo_short();
void demo_int();
void demo_float();
void demo_double();
void demo_long();

// https://www.bilibili.com/video/BV1bs41197KN?p=85
int main(){
  demo_short();
  demo_int();
  demo_float();
  demo_double();
  demo_long();

  cout << endl;
  return 0;
}

void demo_short(){
  cout << "demo_short: " << endl;
  short n = 0;
  short *p = &n;

  cout << p << endl;
  p++;
  cout << p << endl;

  cout << endl;
}

void demo_int(){
  cout << "demo_int: " << endl;
  int n = 0;
  int *p = &n;

  cout << p << endl;
  p++;
  cout << p << endl;

  cout << endl;
}

void demo_float(){
  cout << "demo_float: " << endl;
  float n = 0;
  float *p = &n;

  cout << p << endl;
  p++;
  cout << p << endl;

  cout << endl;
}

void demo_double(){
  cout << "demo_double: " << endl;
  double n = 0;
  double *p = &n;

  cout << p << endl;
  p++;
  cout << p << endl;

  cout << endl;
}

void demo_long(){
  cout << "demo_long: " << endl;
  long n = 0;
  long *p = &n;

  cout << p << endl;
  p++;
  cout << p << endl;

  cout << endl;
}

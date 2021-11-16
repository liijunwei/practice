#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=85
int main(){
  short n = 0;
  short *p = &n;

  cout << p << endl;
  p++;
  cout << p << endl;

  cout << endl;
  return 0;
}


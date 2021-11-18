#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=92

// 遍历数组元素
int main(){
  int a[3][4] = {
    {1, 3,    5,  7},
    {9, 11,  13, 15},
    {17, 19, 21, 23},
  };

  int *p = NULL;

  for(p = &a[0][0]; p < &a[0][0] + 12; p++){
    cout << p << " " << *p << endl;
  }

  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=92

/*
  Q: 如何定义一个 指向 "包含4个int型元素的一位数组"的指针变量?

  A: int (*p)[4]

  (*p)[0], (*p)[1], (*p)[2], (*p)[3]
*/

int main(){
  int a[3][4] = {
    {1, 3,    5,  7},
    {9, 11,  13, 15},
    {17, 19, 21, 23},
  };

  int i;
  int j;
  int (*p)[4] = NULL;

  p = a;

  cin >> i >> j; // i < 3 代表行号, j < 4 代表列号

  cout << *(*(p + i) + j) << endl;
  cout << p[i][j] << endl;

  cout << endl;
  return 0;
}

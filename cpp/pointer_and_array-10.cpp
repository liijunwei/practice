#include <iostream>

using namespace std;


/*
  二维数组名的含义(需要反复看几遍, 多花点时间理解)
  https://www.bilibili.com/video/BV1bs41197KN?p=91

  三条规律

  1. 数组名 相当于 指向数组第一个元素的指针
  2. &E 相当于 把 E的管辖范围 上升了一个级别
  3. *E 相当于 把 E的管辖范围 下降了一个级别
*/

int main(){
  int a[3][4] = {
    {1, 2,   3,  4},
    {5, 6,   7,  8},
    {9, 10, 11, 12},
  };

  // a 是 2维数组的名字
  cout << __LINE__ << "  a = " << a << endl;
  cout << __LINE__ << "  &a[0] = " << &a[0] << endl;
  cout << endl;

  // a[0] 是 1维数组的名字
  cout << __LINE__ << "  a + 1 = " << a + 1 << endl;
  cout << __LINE__ << "  &a[0] + 1 = " << &a[0] + 1 << endl;
  cout << endl;

  cout << __LINE__ << "  a[1] = " << a[1] << endl;
  cout << __LINE__ << "  &a[1] = " << &a[1] << endl;
  cout << __LINE__ << "  *(a + 1) = " << *(a + 1) << endl;
  cout << endl;

  cout << __LINE__ << "  *a + 1 = " << *a + 1 << endl;
  cout << endl;

  cout << __LINE__ << "  &a = " << &a << endl;
  cout << __LINE__ << "  &a + 1 = " << &a + 1 << endl;
  cout << endl;

  cout << endl;
  return 0;
}

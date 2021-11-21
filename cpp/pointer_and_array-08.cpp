#include <iostream>

using namespace std;


/*
  二维数组名的含义
  https://www.bilibili.com/video/BV1bs41197KN?p=90

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

  // a 的含义是 指向 a[0] 的指针
  // a[0] 的含义是 指向 a[0] 这个数组的第一个元素的指针

  /*
    a       指向 a[0] 这个数组
    a[0]    指向 a[0] 这个数组的第0个元素
    a[0][0] 指向 a[0][0] 这个地址里的值
    &a      指向整个数组
  */

  /*
    等价关系

    --------|--------------
    a       | <=> &a[0]
    a[0]    | <=> &a[0][0]
    a[0]    | <=> *a
    a[0][0] | <=> **a
  */

  cout << __LINE__ << " " << a[0] << endl;
  cout << __LINE__ << " " << a[1] << endl;
  cout << __LINE__ << " " << a[2] << endl;
  cout << endl;

  cout << __LINE__ << " " << *a[0] << endl;
  cout << __LINE__ << " " << *a[1] << endl;
  cout << __LINE__ << " " << *a[2] << endl;
  cout << endl;

  cout << endl;
  return 0;
}

#include <iostream>

using namespace std;

// https://www.bilibili.com/video/BV1bs41197KN?p=89

// 数组名相当于指向数组第一个元素的指针

// int a[4]
// 若 a 是指向数组第一个元素的指针, 即a相当于 &a[0]
// &a 是 "指向数组" 的指针, (&a + 1) 将跨越16个字节

// &a 相当于 "管辖范围" "上升"了一级
// *a 相当于 "管辖范围" "下降"了一级

int main(){
  int a[5] = {10, 11, 12, 13, 14};
  // int a[3] = {10, 11, 12};

  cout << __LINE__ << " " << a << endl;
  cout << __LINE__ << " " << (a + 1) << endl;
  cout << __LINE__ << " " << &a << endl;       // 指向      整个数组空间
  cout << __LINE__ << " " << (&a + 1) << endl; // 指向 下个 整个数组空间
  cout << __LINE__ << " " << *(&a) << endl;
  cout << __LINE__ << " " << (*(&a) + 1) << endl;

  cout << endl;
  return 0;
}

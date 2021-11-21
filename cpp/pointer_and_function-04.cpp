#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=93

// Q: 数组名可以做形参吗?
// A: C++ 编译器会 将形参数组名, 作为 指针变量 来处理, 所以可以

// 需要注意, 这个sum函数, 除了做求和这件事外, 还修改了数组内元素的值!!
int sum(int array[], int n){
  for(int i = 0; i < n - 1; i++){
    *(array + 1) = *array + *(array + 1);
    array++;
  }

  return *array;
}

int main(){
  int a[12] = {1, 3, 5, 7, 9, 11, 13, 15, 2, 4, 6, 8};

  cout << "Before using sum, a: ";
  for(int i = 0; i < 12; i++){
    cout << a[i] << " ";
  }

  cout << endl << "The sum is: " << sum(a, 12) << endl;

  cout << "After using sum, a:  ";
  for(int i = 0; i < 12; i++){
    cout << a[i] << " ";
  }

  cout << endl;
  return 0;
}

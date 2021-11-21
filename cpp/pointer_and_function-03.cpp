#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=93

// 多维数组名作为函数参数

// 数组名相当于指向数组第一个元素的指针
// 数组名相当于指向数组第一个元素的指针

int maxvalue(int (*p)[4]){
  int max = p[0][0];

  for(int i = 0; i < 3; i++){
    for(int j = 0; j < 4; j++){
      if(p[i][j] > max){
        max = p[i][j];
      }
    }
  }

  return max;
}

int main(){
  int a[3][4] = {
    {1, 3,   5,  7},
    {9, 11, 13, 15},
    {2,  4,  6,  8},
  };

  cout << "The max value is: " << maxvalue(a);

  cout << endl;
  return 0;
}

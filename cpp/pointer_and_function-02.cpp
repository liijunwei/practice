#include <iostream>

using namespace std;
// https://www.bilibili.com/video/BV1bs41197KN?p=93

// 数组名作为函数参数

void sum(int *p, int n){
  int total = 0;
  for(int i = 0; i < n; i++, p++){
    total += *p;
  }

  cout << "total " << total << endl;
}

int main(){
  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

  sum(a, 10);

  cout << endl;
  return 0;
}

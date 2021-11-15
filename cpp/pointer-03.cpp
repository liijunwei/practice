#include <iostream>

using namespace std;

int main(){
  int a = 0;
  int b = 0;
  int *temp = NULL; // 为指针变量赋初值

  int *p1 = NULL; // 为指针变量赋初值
  int *p2 = NULL; // 为指针变量赋初值

  cin >> a >> b;
  p1 = &a;
  p2 = &b;

  // 交换指针变量里的值
  if(*p1 < *p2){
    temp = p1;
    p1 = p2;
    p2 = temp;
  }

  cout << "Max = " << *p1 << ", min = " << *p2 << endl;

  cout << endl;
  return 0;
}

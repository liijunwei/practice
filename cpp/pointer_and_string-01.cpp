#include <iostream>

using namespace std;
// 指向 数组 的指针 int a[10]; int *p; p = a
// 指向 字符串 的指针 char a[10]; char *p; p = a

// https://www.bilibili.com/video/BV1bs41197KN?p=88
int main(){
  char a[] = "How are you?";
  char b[20];
  char *p1;
  char *p2;

  for(p1 = a, p2 = b; *p1 != '\0'; p1++, p2++){
    *p2 = *p1;
  }

  *p2 = '\0';

  cout << "String a is: " << a << endl;
  cout << "String b is: " << b << endl;

  cout << endl;
  return 0;
}

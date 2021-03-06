#include <iostream>
#include <iomanip>

using namespace std;

void iterate_with_index();
void iterate_with_pointer();

// g++ pointer-06.cpp && ./a.out
int main(){
  iterate_with_index();
  iterate_with_pointer();

  cout << endl;
  return 0;
}


void iterate_with_index(){
  cout << "iterate_with_index:   ";

  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  for(int i = 9; i >= 0; i--){
    cout << setw(3) << a[i] << " ";
  }

  cout << endl;
}

// https://www.bilibili.com/video/BV1bs41197KN?p=87&t=507.9
void iterate_with_pointer(){
  cout << "iterate_with_pointer: ";

  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  int *p = a;

  for(int i = 9; i >= 0; i--){
    cout << setw(3) << p[i] << " ";
  }

  cout << endl;
}



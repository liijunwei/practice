#include <iostream>
#include <iomanip>

using namespace std;

void iterate_with_index();
void iterate_with_pointer();

int main(){
  iterate_with_index();
  iterate_with_pointer();

  cout << endl;
  return 0;
}


void iterate_with_index(){
  cout << "iterate_with_index:   ";
  int a[3];
  int i;
  for(i = 0; i < 3; i++){
    cin >> a[i];
  }

  for(i = 2; i >= 0; i--){
    cout << setw(3) << a[i];
  }

  cout << endl;
}

// https://www.bilibili.com/video/BV1bs41197KN?p=87&t=507.9
void iterate_with_pointer(){
  cout << "iterate_with_pointer: ";
  int a[3];
  int i;
  int *p = a;

  for(i = 0; i < 3; i++){
    cin >> *p++;
  }

  for(p--; p >= a; ){
    cout << setw(3) << *p--;
  }

  cout << endl;
}



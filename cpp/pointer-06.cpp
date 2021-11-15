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
  cout << "iterate_with_index: ";

  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  for(int i = 9; i >= 0; i--){
    cout << setw(3) << a[i] << " ";
  }

  cout << endl;
}

// TODO 没看懂, 没打印对
void iterate_with_pointer(){
  cout << "iterate_with_pointer: ";

  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  int *p = a;

  for(p--; p >= a;){
    cout << setw(3) << *p++ << " ";
  }

  cout << endl;
}



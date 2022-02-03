#include <iostream>

using namespace std;

// 倒置数组元素
// reverse elements in an array

int main(){
  int a[3];
  int *p = NULL;
  int *q = NULL;
  int temp;

  for(p = a; p < a + 3; p++){
    cin >> *p;
  }

  for(p = a, q = a + 2; p < q; p++, q--){
    temp = *p;
    *p = *q;
    *q = temp;
  }

  for(p = a; p < a + 3; p++){
    cout << " " << *p;
  }

  cout << endl;
  return 0;
}


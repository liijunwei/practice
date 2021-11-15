#include <iostream>
#include <iomanip>

using namespace std;

// 倒置数组元素
// reverse elements in an array

int main(){
  int a[5] = {1, 2, 3, 4, 5};

  int *p = NULL;
  int *q = NULL;
  int temp;

  for(p = a, q = a + 4; p < q; p++, q--){
    temp = *p;
    *p = *q;
    *q = temp;
  }

  for(p = a; p < a + 5; p++){
    cout << setw(3) << *p;
  }

  cout << endl;
  return 0;
}


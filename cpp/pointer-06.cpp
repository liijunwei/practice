#include <iostream>
#include <iomanip>

using namespace std;

int main(){
  int a[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
  for(int i = 9; i >= 0; i--){
    cout << setw(3) << a[i] << " ";
  }

  cout << endl;
  return 0;
}
